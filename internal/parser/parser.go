package parser

import (
	"fmt"
	"strconv"

	"github.com/aporeto-inc/manipulate"
)

// FilterParser represents a Parser
type FilterParser struct {
	scanner *scanner
	buffer  struct {
		token   parserToken // last read token
		literal string      // last read literal
		size    int         // buffer size (max=1)
	}
}

// NewFilterParser returns an instance of FilterParser for the given input
func NewFilterParser(input string) *FilterParser {
	return &FilterParser{
		scanner: newScanner(input),
	}
}

// scan returns the next token scanned or the last bufferred one
func (p *FilterParser) scan() (parserToken, string) {
	// If a token has been buffered, use it
	if p.buffer.size != 0 {
		p.buffer.size = 0
		return p.buffer.token, p.buffer.literal
	}

	// Otherwise scan the next token
	token, literal := p.scanner.Scan()

	// Save it to the buffer in case we unscan later.
	p.buffer.token, p.buffer.literal = token, literal

	return token, literal
}

// unscan put back the token into the buffer.
func (p *FilterParser) unscan() {
	p.buffer.size = 1
}

// peekIgnoreWhitespace scans to the next whitespace and unscan it
func (p *FilterParser) peekIgnoreWhitespace() (tok parserToken, lit string) {

	tok, lit = p.scanIgnoreWhitespace()
	p.unscan()
	return
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *FilterParser) scanIgnoreWhitespace() (tok parserToken, lit string) {

	tok, lit = p.scan()

	if tok == parserTokenWHITESPACE {
		tok, lit = p.scan()
	}

	return
}

// Parse parses the input string and returns a new manipulate.Context
func (p *FilterParser) Parse() (*manipulate.Filter, error) {

	token, literal := p.peekIgnoreWhitespace()

	// The input needs to start with a word, a quote or a left parenthese.
	if token != parserTokenWORD &&
		token != parserTokenQUOTE &&
		token != parserTokenLEFTPARENTHESE {
		return nil, fmt.Errorf("invalid start of expression. found %s", literal)
	}

	// Stack all discovered the filters
	var stack []*manipulate.Filter
	// Default conjunction is an and.
	var conjunction parserToken = -1

	finalFilter := manipulate.NewFilterComposer()
	for {
		token, literal := p.scanIgnoreWhitespace()

		if token == parserTokenEOF || token == parserTokenRIGHTPARENTHESE {
			// In case of EOF or ")", we need to finalize the filter.
			if len(stack) == 1 {
				finalFilter = stack[0]

			} else if len(stack) > 1 {
				if conjunction == parserTokenOR {
					finalFilter.Or(stack...)
				} else {
					finalFilter.And(stack...)
				}
				stack = []*manipulate.Filter{}
			}

			break
		}

		if token == parserTokenQUOTE {
			// Handle expression starting with QUOTE like "a" operator b
			token, literal = p.scanIgnoreWhitespace()
			fmt.Println(token, literal)
			if token != parserTokenWORD {
				return nil, fmt.Errorf("invalid word after the quote. found %s", literal)
			}
			quote, _ := p.scanIgnoreWhitespace()
			if quote != parserTokenQUOTE {
				return nil, fmt.Errorf("missing quote after the word %s", literal)
			}

			operator, err := p.parseOperator()
			if err != nil {
				return nil, err
			}

			value, err := p.parseValue()
			if err != nil {
				return nil, err
			}

			filter, err := makeFilter(literal, operator, value)
			if err != nil {
				return nil, err
			}

			stack = append(stack, filter)
			continue
		}

		if token == parserTokenWORD {
			// Handle expression without QUOTE like a operator b
			// In that case, the literal is the word scanned.
			operator, err := p.parseOperator()
			if err != nil {
				return nil, err
			}

			value, err := p.parseValue()
			if err != nil {
				return nil, err
			}

			filter, err := makeFilter(literal, operator, value)
			if err != nil {
				return nil, err
			}

			stack = append(stack, filter)
			continue
		}

		if token == parserTokenAND {
			// Switch the conjunction to an AND
			if conjunction != -1 && conjunction != parserTokenAND {
				return nil, fmt.Errorf("misleading \"and\" condition. please add parentheses")
			}
			conjunction = token
			continue
		}

		if token == parserTokenOR {
			// Switch the conjunction to an OR
			if conjunction != -1 && conjunction != parserTokenOR {
				return nil, fmt.Errorf("misleading \"or\" condition. please add parentheses")
			}
			conjunction = token
			continue
		}

		if token == parserTokenLEFTPARENTHESE {
			// In case of "(", a subfilter needs to be computed
			// and stacked to the previously found filters.

			if len(stack) == 0 {
				// If there are no previously found filters,
				// start with the subfilter.
				subFilter, err := p.Parse()
				if err != nil {
					return nil, err
				}
				// finalFilter = subFilter
				stack = append(stack, subFilter)

			} else {
				subFilter, err := p.Parse()
				if err != nil {
					return nil, err
				}

				stack = append(stack, subFilter)
				// if conjunction == parserTokenOR {
				// 	finalFilter.Or(stack...)
				// } else {
				// 	finalFilter.And(stack...)
				// }
				// stack = []*manipulate.Filter{}
			}
		}
	}

	return finalFilter.Done(), nil
}

func makeFilter(key string, operator parserToken, value interface{}) (*manipulate.Filter, error) {

	filter := manipulate.NewFilterComposer()

	// Create filter
	switch operator {
	case parserTokenEQUAL:
		filter.WithKey(key).Equals(value)
	case parserTokenNOTEQUAL:
		filter.WithKey(key).NotEquals(value)
	case parserTokenLT:
		filter.WithKey(key).LesserThan(value)
	case parserTokenLTE:
		filter.WithKey(key).LesserThan(value)
	case parserTokenGT:
		filter.WithKey(key).GreaterThan(value)
	case parserTokenGTE:
		filter.WithKey(key).GreaterThan(value)
	case parserTokenCONTAINS:
		filter.WithKey(key).Contains(value)
	case parserTokenMATCHES:
		filter.WithKey(key).Matches(value)
	default:
		return nil, fmt.Errorf("unsupported operator")
	}

	return filter.Done(), nil
}

func (p *FilterParser) parseStringValue() (string, error) {

	p.unscan()
	token, literal := p.scanIgnoreWhitespace()
	isQuoted := false

	if token == parserTokenQUOTE {
		isQuoted = true
		token, literal = p.scanIgnoreWhitespace()
	}

	if token != parserTokenWORD {
		return "", fmt.Errorf("invalid value. found %s", literal)
	}

	if isQuoted {
		if token, _ = p.scanIgnoreWhitespace(); token != parserTokenQUOTE {
			return "", fmt.Errorf("missing quote after the value %s", literal)
		}
	} else {
		if token, _ = p.peekIgnoreWhitespace(); token == parserTokenQUOTE {
			return "", fmt.Errorf("missing quote before the value %s", literal)
		}
	}

	return literal, nil
}

func (p *FilterParser) parseOperator() (parserToken, error) {

	token, literal := p.scanIgnoreWhitespace()

	if token != parserTokenEQUAL &&
		token != parserTokenNOTEQUAL &&
		token != parserTokenLT &&
		token != parserTokenLTE &&
		token != parserTokenGT &&
		token != parserTokenGTE &&
		token != parserTokenCONTAINS &&
		token != parserTokenMATCHES {
		return parserTokenILLEGAL, fmt.Errorf("invalid operator. found %s", literal)
	}

	return token, nil
}

func (p *FilterParser) parseValue() (interface{}, error) {

	token, literal := p.scanIgnoreWhitespace()

	if token == parserTokenQUOTE {
		return p.parseStringValue()
	}

	if token == parserTokenTRUE {
		return true, nil
	}

	if token == parserTokenFALSE {
		return false, nil
	}

	if i, err := strconv.ParseInt(literal, 10, 0); err == nil {
		return i, nil
	}

	if f, err := strconv.ParseFloat(literal, 0); err == nil {
		return f, nil
	}

	v, err := p.parseStringValue()
	if err != nil {
		return nil, err
	}

	return v, nil
}
