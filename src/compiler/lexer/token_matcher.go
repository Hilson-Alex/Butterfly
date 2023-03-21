package lexer

import (
	"regexp"

	"github.com/Hilson-Alex/Butterfly/src/compiler/parser"
)

// tokPattern creates a regexp that matches the pattern on
// the start of a string
func tokPattern(pattern string) *regexp.Regexp {
	return regexp.MustCompile("^(" + pattern + ")")
}

// util patterns for string matching
const (
	bSlash       = "\\\\"
	validChars   = "[^\"\\n" + bSlash + "]"
	doubleBSlash = bSlash + bSlash
	escaped      = bSlash + "(\"|\\w)"
)

// regex patterns for tokens
const (
	// Values
	unsignedNumber = "[0-9]+"
	signedNumber   = "-" + unsignedNumber
	floatNumber    = "(" + signedNumber + "|" + unsignedNumber + ")[.]" + unsignedNumber
	text           = "\"(" + validChars + "|" + escaped + "|" + doubleBSlash + ")*\""

	// Reserved keywords
	boolean          = "(true|false)\\b"
	typePattern      = "(bool|byte|int|uint|float|string)\\b"
	constDeclaration = "const\\b"
	letDeclaration   = "let\\b"
	doLoop           = "do\\b"
	whileLoop        = "while\\b"
	forLoop          = "for\\b"
	ifStatement      = "if\\b"
	elseStatement    = "else\\b"
	selector         = "switch\\b"
	caseStatement    = "case\\b"

	// Event keywords
	on      = "on\\b"
	share   = "share\\b"
	finish  = "finish\\b"
	module  = "module\\b"
	message = "message\\b"

	// Symbols
	delimiter  = ";"
	increment  = "[+]{2}|--"
	not        = "!"
	logic      = "&&|[|]{2}"
	arithmetic = "[+\\-*/%]"
	comparator = "[=!><]=|>|<"
	assign     = arithmetic + "?="
	comma      = ","
	colon      = ":"

	// Groupers
	openCurly        = "\\{"
	closeCurly       = "\\}"
	openParenthesis  = "\\("
	closeParenthesis = "\\)"
	openBrackets     = "\\["
	closeBrackets    = "\\]"

	// Identifier
	identifier = "[a-zA-Z_][a-zA-Z0-9_]*"
)

// tokMatcher is a type for token matching. it links a token
// code to its pattern.
type tokMatcher struct {
	matcher *regexp.Regexp
	code    int
}

// special patterns
var (
	lineComment  = tokPattern("//[^\\n]*(\\n|$)")
	unknownToken = tokPattern("(.+?)(;| |[\\-+*/%]|[=!><]|,|:|&&|[|]{2}|\\r?\\n|$)")
)

// getUnknownToken parse everything until the next symbol token
// this must be used only to recover from an error
func getUnknownToken(match string) string {
	return unknownToken.FindStringSubmatch(match)[2]
}

// array that stores all the token matchers
var matchers = [...]tokMatcher{
	{tokPattern(text), parser.TEXT},
	{tokPattern(floatNumber), parser.FLOAT_NUMBER}, // float must be parsed before other numerics
	{tokPattern(unsignedNumber), parser.UNSIGNED_NUMBER},
	{tokPattern(signedNumber), parser.SIGNED_NUMBER},

	{tokPattern(boolean), parser.BOOLEAN},
	{tokPattern(typePattern), parser.TYPE},
	{tokPattern(constDeclaration), parser.CONST},
	{tokPattern(letDeclaration), parser.LET},
	{tokPattern(doLoop), parser.DO},
	{tokPattern(whileLoop), parser.WHILE},
	{tokPattern(forLoop), parser.FOR},
	{tokPattern(ifStatement), parser.IF},
	{tokPattern(elseStatement), parser.ELSE},
	{tokPattern(selector), parser.SWITCH},
	{tokPattern(caseStatement), parser.CASE},

	{tokPattern(delimiter), parser.DELIMITER},
	{tokPattern(increment), parser.INCREMENT},
	{tokPattern(not), parser.NOT},
	{tokPattern(logic), parser.LOGIC},
	{tokPattern(arithmetic), parser.ARITHMETIC},
	{tokPattern(comparator), parser.COMPARATOR},
	{tokPattern(assign), parser.ASSIGN},
	{tokPattern(comma), parser.COMMA},
	{tokPattern(colon), parser.COLON},

	{tokPattern(on), parser.ON},
	{tokPattern(share), parser.SHARE},
	{tokPattern(finish), parser.FINISH},
	{tokPattern(module), parser.MODULE},
	{tokPattern(message), parser.MESSAGE},

	{tokPattern(openCurly), parser.OP_CURLY},
	{tokPattern(closeCurly), parser.CL_CURLY},
	{tokPattern(openParenthesis), parser.OP_PARENT},
	{tokPattern(closeParenthesis), parser.CL_PARENT},
	{tokPattern(openBrackets), parser.OP_SQUARE},
	{tokPattern(closeBrackets), parser.CL_SQUARE},

	{tokPattern(identifier), parser.IDENTIFIER},
}
