package lexer

import (
	"github.com/Hilson-Alex/Butterfly/src/compiler/parser"
)

func printToken(code int) string {
	switch code {
	case parser.ANY:
		return "any"
	case parser.CONST:
		return "const"
	case parser.LET:
		return "let"
	case parser.DO:
		return "do"
	case parser.WHILE:
		return "while"
	case parser.FOR:
		return "for"
	case parser.IF:
		return "if"
	case parser.ELSE:
		return "else"
	case parser.SWITCH:
		return "switch"
	case parser.CASE:
		return "case"
	case parser.DEFAULT:
		return "default"

	case parser.DELIMITER:
		return ";"
	case parser.INCREMENT:
		return "incremental operator (++ or --)"
	case parser.EXP_ASSIGN:
		return "expression assign (+=, -=, *=, /=, \\= or %=)"
	case parser.ASSIGN:
		return "="
	case parser.NOT:
		return "!"
	case parser.LOGIC:
		return "Logical operator (&& or ||)"
	case parser.ARITHMETIC:
		return "arithmetic operator (+, -, *, /, \\ or %)"
	case parser.COMPARATOR:
		return "relational operator (>, <, >=, <=, != or ==)"
	case parser.COMMA:
		return ","
	case parser.COLON:
		return ":"
	case parser.DOT:
		return "."

	case parser.ON:
		return "on"
	case parser.SHARE:
		return "share"
	case parser.FINISH:
		return "finish"
	case parser.MODULE:
		return "module"
	case parser.MESSAGE:
		return "message"

	case parser.OP_CURLY:
		return "{"
	case parser.CL_CURLY:
		return "}"
	case parser.OP_PARENT:
		return "("
	case parser.CL_PARENT:
		return ")"
	case parser.OP_SQUARE:
		return "["
	case parser.CL_SQUARE:
		return "]"
	}
	return parser.TokenName(code)
}

func init() {
	parser.PrintToken = printToken
}
