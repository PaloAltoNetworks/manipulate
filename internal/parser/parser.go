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

			switch len(stack) {
			case 0:
				// Nothing to do here.
			case 1:
				// No specific combination
				finalFilter = stack[0]
			default:
				// Combination depending on the conjunction
				if conjunction == parserTokenOR {
					finalFilter.Or(stack...)
				} else {
					finalFilter.And(stack...)
				}
			}
			break
		}

		if token == parserTokenQUOTE {
			// Handle expression starting with QUOTE like "a" operator b
			token, literal = p.scanIgnoreWhitespace()

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
			if conjunction != -1 && conjunction == parserTokenOR && len(stack) > 1 {
				filter := manipulate.NewFilterComposer().Or(stack...).Done()
				stack = []*manipulate.Filter{filter}
			}
			conjunction = token
			continue
		}

		if token == parserTokenOR {
			// Switch the conjunction to an OR
			if conjunction != -1 && conjunction == parserTokenAND && len(stack) > 1 {
				filter := manipulate.NewFilterComposer().And(stack...).Done()
				stack = []*manipulate.Filter{filter}
			}
			conjunction = token
			continue
		}

		if token == parserTokenLEFTPARENTHESE {
			// In case of "(", a subfilter needs to be computed
			// and stacked to the previously found filters.

			subFilter, err := p.Parse()
			if err != nil {
				return nil, err
			}
			stack = append(stack, subFilter)
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
		if values, ok := value.([]interface{}); ok {
			filter.WithKey(key).Contains(values...)
		} else {
			filter.WithKey(key).Contains(value)
		}
	case parserTokenNOTCONTAINS:
		if values, ok := value.([]interface{}); ok {
			filter.WithKey(key).NotContains(values...)
		} else {
			filter.WithKey(key).NotContains(value)
		}
	case parserTokenIN:
		if values, ok := value.([]interface{}); ok {
			filter.WithKey(key).In(values...)
		} else {
			filter.WithKey(key).In(value)
		}
	case parserTokenNOTIN:
		if values, ok := value.([]interface{}); ok {
			filter.WithKey(key).NotIn(values...)
		} else {
			filter.WithKey(key).NotIn(value)
		}
	case parserTokenMATCHES:
		if values, ok := value.([]interface{}); ok {
			filter.WithKey(key).Matches(values...)
		} else {
			filter.WithKey(key).Matches(value)
		}
	default:
		return nil, fmt.Errorf("unsupported operator")
	}

	return filter.Done(), nil
}

func (p *FilterParser) parseOperator() (parserToken, error) {

	token, literal := p.scanIgnoreWhitespace()

	operatorNot := false
	if token == parserTokenNOT {
		operatorNot = true
		token, literal = p.scanIgnoreWhitespace()
	}

	if token != parserTokenEQUAL &&
		token != parserTokenNOTEQUAL &&
		token != parserTokenLT &&
		token != parserTokenLTE &&
		token != parserTokenGT &&
		token != parserTokenGTE &&
		token != parserTokenCONTAINS &&
		token != parserTokenIN &&
		token != parserTokenMATCHES {
		return parserTokenILLEGAL, fmt.Errorf("invalid operator. found %s", literal)
	}

	if operatorNot {
		switch token {
		case parserTokenCONTAINS:
			return parserTokenNOTCONTAINS, nil
		case parserTokenIN:
			return parserTokenNOTIN, nil
		default:
			return parserTokenILLEGAL, fmt.Errorf("invalid usage of operator NOT before %s", literal)
		}
	}

	return token, nil
}

func (p *FilterParser) parseValue() (interface{}, error) {

	token, literal := p.scanIgnoreWhitespace()

	if token == parserTokenQUOTE {
		return p.parseStringValue()
	}

	if token == parserTokenLEFTSQUAREPARENTHESE {
		return p.parseArrayValue()
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

func (p *FilterParser) parseStringValue() (string, error) {

	p.unscan()
	token, literal := p.scanIgnoreWhitespace()
	var value string

	// Quoted string
	if token == parserTokenQUOTE {
		// Scan everything until the next quote or the end of the input
		for {
			token, literal = p.scan()
			if token == parserTokenEOF {
				return "", fmt.Errorf("unable to find quote after value: %s", value)
			}

			if token == parserTokenQUOTE {
				return value, nil
			}

			// Add anything to the value
			value += literal
		}
	}

	// Unquoted string can have only one word
	if token != parserTokenWORD {
		return "", fmt.Errorf("invalid value. found %s", literal)
	}

	token, next := p.peekIgnoreWhitespace()
	switch token {
	case parserTokenQUOTE:
		return "", fmt.Errorf("missing quote before the value: %s", literal)
	case parserTokenWORD:
		return "", fmt.Errorf("missing parenthese to protect value: %s %s", literal, next)
	}

	return literal, nil
}

func (p *FilterParser) parseArrayValue() ([]interface{}, error) {

	p.unscan()
	token, literal := p.scanIgnoreWhitespace()
	if token != parserTokenLEFTSQUAREPARENTHESE {
		return nil, fmt.Errorf("invalid start of list. found %s", literal)
	}

	values := []interface{}{}

	for {
		token, literal := p.scanIgnoreWhitespace()

		if token == parserTokenEOF ||
			token == parserTokenLEFTPARENTHESE ||
			token == parserTokenRIGHTPARENTHESE {
			return nil, fmt.Errorf("invalid end of array. found %s", literal)
		}

		if token == parserTokenRIGHTSQUAREPARENTHESE {
			break
		}

		if token == parserTokenCOMMA {
			continue
		}

		p.unscan()
		value, err := p.parseValue()
		if err != nil {
			return nil, err
		}

		values = append(values, value)
	}

	return values, nil
}
