package manipulate

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type checkRuneFunc = func(ch rune) bool

type parserToken int

const (
	parserTokenILLEGAL parserToken = iota
	parserTokenEOF
	parserTokenWHITESPACE
	parserTokenWORD
	parserTokenLEFTPARENTHESE
	parserTokenRIGHTPARENTHESE
	parserTokenAND
	parserTokenOR
	parserTokenQUOTE
	parserTokenEQUAL
	parserTokenNOTEQUAL
	parserTokenLT
	parserTokenLTE
	parserTokenGT
	parserTokenGTE
	parserTokenCONTAINS
	parserTokenMATCHES
	parserTokenTRUE
	parserTokenFALSE
	parserTokenLEFTSQUAREPARENTHESE
	parserTokenRIGHTSQUAREPARENTHESE
	parserTokenCOMMA
	parserTokenNOTCONTAINS
	// parserTokenNOTMATCHES not implemented yet in filters.
	parserTokenIN
	parserTokenNOTIN
	parserTokenNOT
)

const (
	wordAND         = "AND"
	wordCONTAINS    = "CONTAINS"
	wordEQUAL       = "=="
	wordFALSE       = "FALSE"
	wordGT          = ">"
	wordGTE         = ">="
	wordIN          = "IN"
	wordLT          = "<"
	wordLTE         = "<="
	wordMATCHES     = "MATCHES"
	wordNOTCONTAINS = "NOT CONTAINS"
	wordNOTEQUAL    = "!="
	wordNOTIN       = "NOT IN"
	wordOR          = "OR"
	wordTRUE        = "TRUE"
	wordNOT         = "NOT"
)

const (
	runeEOF                   = rune(0)
	runeLEFTPARENTHESE        = '('
	runeRIGHTPARENTHESE       = ')'
	runeQUOTE                 = '"'
	runeLEFTSQUAREPARENTHESE  = '['
	runeRIGHTSQUAREPARENTHESE = ']'
	runeCOMMA                 = ','
)

var specialLetters = map[rune]interface{}{
	'-':  nil,
	'_':  nil,
	'@':  nil,
	':':  nil,
	'$':  nil,
	'#':  nil,
	'.':  nil,
	'/':  nil,
	'\\': nil,
	'<':  nil,
	'>':  nil,
	'=':  nil,
	'!':  nil,
	'*':  nil,
}

var stringToToken = map[string]parserToken{
	wordAND:         parserTokenAND,
	wordOR:          parserTokenOR,
	wordEQUAL:       parserTokenEQUAL,
	wordNOTEQUAL:    parserTokenNOTEQUAL,
	wordLT:          parserTokenLT,
	wordLTE:         parserTokenLTE,
	wordGT:          parserTokenGT,
	wordGTE:         parserTokenGTE,
	wordCONTAINS:    parserTokenCONTAINS,
	wordNOTCONTAINS: parserTokenNOTCONTAINS,
	wordMATCHES:     parserTokenMATCHES,
	wordTRUE:        parserTokenTRUE,
	wordFALSE:       parserTokenFALSE,
	wordIN:          parserTokenIN,
	wordNOTIN:       parserTokenNOTIN,
	wordNOT:         parserTokenNOT,
}

var runeToToken = map[rune]parserToken{
	runeEOF:                   parserTokenEOF,
	runeLEFTPARENTHESE:        parserTokenLEFTPARENTHESE,
	runeRIGHTPARENTHESE:       parserTokenRIGHTPARENTHESE,
	runeQUOTE:                 parserTokenQUOTE,
	runeLEFTSQUAREPARENTHESE:  parserTokenLEFTSQUAREPARENTHESE,
	runeRIGHTSQUAREPARENTHESE: parserTokenRIGHTSQUAREPARENTHESE,
	runeCOMMA:                 parserTokenCOMMA,
}

var datePattern = regexp.MustCompile("^date\\((.*)\\)$")
var nowPattern = regexp.MustCompile("^now\\((.*)\\)$")
var dateLayout = "2006-01-02"
var dateTimeLayout = "2006-01-02 15:04"
var errorNotADate = "not a date"

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

// Parse parses the input string and returns a new Filter.
func (p *FilterParser) Parse() (*Filter, error) {

	token, literal := p.peekIgnoreWhitespace()

	// The input needs to start with a word, a quote or a left parenthese.
	if token != parserTokenWORD &&
		token != parserTokenQUOTE &&
		token != parserTokenLEFTPARENTHESE {
		return nil, fmt.Errorf("invalid start of expression. found %s", literal)
	}

	// Stack all discovered the filters
	var stack []*Filter
	// Default conjunction is an and.
	var conjunction parserToken = -1

	finalFilter := NewFilterComposer()
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

			filter, err := p.makeFilter(literal, operator, value)
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

			filter, err := p.makeFilter(literal, operator, value)
			if err != nil {
				return nil, err
			}

			stack = append(stack, filter)
			continue
		}

		if token == parserTokenAND {
			// Switch the conjunction to an AND
			if conjunction != -1 && conjunction == parserTokenOR && len(stack) > 1 {
				filter := NewFilterComposer().Or(stack...).Done()
				stack = []*Filter{filter}
			}
			conjunction = token
			continue
		}

		if token == parserTokenOR {
			// Switch the conjunction to an OR
			if conjunction != -1 && conjunction == parserTokenAND && len(stack) > 1 {
				filter := NewFilterComposer().And(stack...).Done()
				stack = []*Filter{filter}
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

// scan returns the next token scanned or the last bufferred one
func (p *FilterParser) scan() (parserToken, string) {
	// If a token has been buffered, use it
	if p.buffer.size != 0 {
		p.buffer.size = 0
		return p.buffer.token, p.buffer.literal
	}

	// Otherwise scan the next token
	token, literal := p.scanner.scan()

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

func (p *FilterParser) makeFilter(key string, operator parserToken, value interface{}) (*Filter, error) {

	filter := NewFilterComposer()

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

	token, literal := p.peekIgnoreWhitespace()
	if token != parserTokenAND &&
		token != parserTokenOR &&
		token != parserTokenRIGHTPARENTHESE &&
		token != parserTokenEOF {
		return nil, fmt.Errorf("invalid keyword after %s. Found %s", filter.Done().String(), literal)
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

	t, err := p.parseDateValue()
	if err == nil {
		return t, nil
	}

	if err.Error() != errorNotADate {
		return nil, err
	}

	v, err := p.parseStringValue()
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (p *FilterParser) parseDateValue() (time.Time, error) {

	p.unscan()
	token, literal := p.scanIgnoreWhitespace()

	if !strings.HasPrefix(literal, "date") && !strings.HasPrefix(literal, "now") {
		return time.Time{}, fmt.Errorf(errorNotADate)
	}

	expression := literal
	token, literal = p.scanIgnoreWhitespace()
	if token != parserTokenLEFTPARENTHESE {
		p.unscan()
		return time.Time{}, fmt.Errorf(errorNotADate)
	}

	expression += literal

	for { // Read the expression until the next )
		token, literal = p.scan()
		expression += literal

		if token == parserTokenLEFTPARENTHESE ||
			token == parserTokenEOF {
			return time.Time{}, fmt.Errorf("invalid date format. Found %s", literal)
		}

		if token == parserTokenRIGHTPARENTHESE {
			break
		}
	}

	// Scan for a format date("...")
	matches := datePattern.FindStringSubmatch(expression)
	if len(matches) == 2 {
		inputValue := strings.Trim(strings.TrimSpace(matches[1]), "\"")

		// Attempt multiple format...
		t, err := time.Parse(time.RFC3339, inputValue)
		if err == nil {
			return t, nil
		}
		t, err = time.Parse(dateLayout, inputValue)
		if err == nil {
			return t, nil
		}
		t, err = time.Parse(dateTimeLayout, inputValue)
		if err == nil {
			return t, nil
		}

		return t, fmt.Errorf("unable to parse date format %s", matches[1])
	}

	// Scan for a format now() or now(-1h)
	matches = nowPattern.FindStringSubmatch(expression)
	if len(matches) == 2 {
		inputValue := strings.Trim(strings.TrimSpace(matches[1]), "\"")

		if inputValue == "" {
			return time.Now(), nil
		}

		d, err := time.ParseDuration(inputValue)
		if err != nil {
			return time.Time{}, fmt.Errorf("unable to parse duration %s: %s", matches[1], err.Error())
		}
		return time.Now().Add(d), nil
	}

	return time.Time{}, fmt.Errorf("invalid date format %s", expression)
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

func isWhitespace(ch rune) bool { return ch == ' ' || ch == '\t' || ch == '\n' }
func isDigit(ch rune) bool      { return (ch >= '0' && ch <= '9') }

func isLetter(ch rune) bool {

	if _, ok := specialLetters[ch]; ok {
		return true
	}

	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}
