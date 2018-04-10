package parser

import (
	"bytes"
	"strings"
)

// scanner scans a given input
type scanner struct {
	buf          bytes.Buffer
	isWhitespace checkRuneFunc
	isLetter     checkRuneFunc
	isDigit      checkRuneFunc
}

// newScanner returns an instance of a Scanner.
func newScanner(
	input string,
) *scanner {
	var buf bytes.Buffer
	buf.WriteString(input)

	return &scanner{
		buf:          buf,
		isWhitespace: isWhitespace,
		isLetter:     isLetter,
		isDigit:      isDigit,
	}
}

// read returns the next rune or eof
func (s *scanner) read() rune {

	ch, _, err := s.buf.ReadRune()
	if err != nil {
		return runeEOF
	}
	return ch
}

// unread a previously read rune
func (s *scanner) unread() {
	_ = s.buf.UnreadRune()
}

// Scan returns the next token and literal value.
func (s *scanner) Scan() (parserToken, string) {

	ch := s.read()

	if s.isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()

	} else if s.isLetter(ch) || s.isDigit(ch) {
		s.unread()
		return s.scanWord()
	}

	token, ok := runeToToken[ch]
	if !ok {
		return parserTokenILLEGAL, string(ch)
	}

	return token, string(ch)
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *scanner) scanWhitespace() (parserToken, string) {

	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == runeEOF {
			break
		} else if !s.isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return parserTokenWHITESPACE, buf.String()
}

// scanWord consumes the current rune and all contiguous letters / digits.
func (s *scanner) scanWord() (parserToken, string) {

	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == runeEOF {
			break
		} else if !s.isLetter(ch) && !s.isDigit(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	output := buf.String()

	token, ok := stringToToken[strings.ToUpper(output)]
	if !ok {
		return parserTokenWORD, output
	}

	return token, output
}
