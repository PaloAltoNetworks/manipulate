package parser

type checkRuneFunc = func(ch rune) bool

type parserToken int

const (
	parserTokenILLEGAL               parserToken = iota // 0
	parserTokenEOF                                      // 1
	parserTokenWHITESPACE                               // 2
	parserTokenWORD                                     // 3
	parserTokenLEFTPARENTHESE                           // 4
	parserTokenRIGHTPARENTHESE                          // 5
	parserTokenAND                                      // 6
	parserTokenOR                                       // 7
	parserTokenQUOTE                                    // 8
	parserTokenEQUAL                                    // 9
	parserTokenNOTEQUAL                                 // 10
	parserTokenLT                                       // 11
	parserTokenLTE                                      // 12
	parserTokenGT                                       // 13
	parserTokenGTE                                      // 14
	parserTokenCONTAINS                                 // 15
	parserTokenMATCHES                                  // 16
	parserTokenTRUE                                     // 17
	parserTokenFALSE                                    // 18
	parserTokenLEFTSQUAREPARENTHESE                     // 19
	parserTokenRIGHTSQUAREPARENTHESE                    // 20
	parserTokenCOMMA                                    // 21
)

const (
	wordAND      = "AND"
	wordOR       = "OR"
	wordEQUAL    = "=="
	wordNOTEQUAL = "!="
	wordLT       = "<"
	wordLTE      = "<="
	wordGT       = ">"
	wordGTE      = ">="
	wordCONTAINS = "IN"
	wordMATCHES  = "MATCHES"
	wordTRUE     = "TRUE"
	wordFALSE    = "FALSE"
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

var stringToToken = map[string]parserToken{
	wordAND:      parserTokenAND,
	wordOR:       parserTokenOR,
	wordEQUAL:    parserTokenEQUAL,
	wordNOTEQUAL: parserTokenNOTEQUAL,
	wordLT:       parserTokenLT,
	wordLTE:      parserTokenLTE,
	wordGT:       parserTokenGT,
	wordGTE:      parserTokenGTE,
	wordCONTAINS: parserTokenCONTAINS,
	wordMATCHES:  parserTokenMATCHES,
	wordTRUE:     parserTokenTRUE,
	wordFALSE:    parserTokenFALSE,
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

// isWhitespace returns true if the given rune is a whitespace
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

var specialLetters = map[rune]interface{}{
	'-': nil,
	'_': nil,
	'@': nil,
	':': nil,
	'$': nil,
	'#': nil,
	'.': nil,
	'/': nil,
	'<': nil,
	'>': nil,
	'=': nil,
	'!': nil,
	'*': nil,
}

// isLetter returns true if the given rune is a letter.
func isLetter(ch rune) bool {

	if _, ok := specialLetters[ch]; ok {
		return true
	}

	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z')
}

// isDigit returns true if the given rune is a digit.
func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}
