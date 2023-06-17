package lexer

import (
	"github.com/Hilson-Alex/Butterfly/src/compiler/parser"
)

func printToken(category string) string {
	switch category {
	case parser.TokenName(parser.ANY):
		return "\"any\""
	case parser.TokenName(parser.CONST):
		return "\"const\""
	case parser.TokenName(parser.LET):
		return "\"let\""
	case parser.TokenName(parser.DO):
		return "\"do\""
	case parser.TokenName(parser.WHILE):
		return "\"while\""
	case parser.TokenName(parser.FOR):
		return "\"for\""
	case parser.TokenName(parser.IF):
		return "\"if\""
	case parser.TokenName(parser.ELSE):
		return "\"else\""
	case parser.TokenName(parser.SWITCH):
		return "\"switch\""
	case parser.TokenName(parser.CASE):
		return "\"case\""
	case parser.TokenName(parser.DEFAULT):
		return "\"default\""

	case parser.TokenName(parser.DELIMITER):
		return "\";\""
	case parser.TokenName(parser.INCREMENT):
		return "incremental operator (\"++\" or \"--\")"
	case parser.TokenName(parser.EXP_ASSIGN):
		return "expression assign (\"+=\", \"-=\", \"*=\", \"/=\", \"\\=\" or \"%=\")"
	case parser.TokenName(parser.ASSIGN):
		return "\"=\""
	case parser.TokenName(parser.NOT):
		return "\"!\""
	case parser.TokenName(parser.LOGIC):
		return "Logical operator (\"&&\" or \"||\")"
	case parser.TokenName(parser.ARITHMETIC):
		return "arithmetic operator (\"+\", \"-\", \"*\", \"/\", \"\\\" or \"%\")"
	case parser.TokenName(parser.COMPARATOR):
		return "relational operator (\">\", \"<\", \">=\", \"<=\", \"!=\" or \"==\")"
	case parser.TokenName(parser.COMMA):
		return "\",\""
	case parser.TokenName(parser.COLON):
		return "\":\""
	case parser.TokenName(parser.DOT):
		return "\".\""

	case parser.TokenName(parser.ON):
		return "\"on\""
	case parser.TokenName(parser.SHARE):
		return "\"share\""
	case parser.TokenName(parser.FINISH):
		return "\"finish\""
	case parser.TokenName(parser.MODULE):
		return "\"module\""
	case parser.TokenName(parser.MESSAGE):
		return "\"message\""

	case parser.TokenName(parser.OP_CURLY):
		return "\"{\""
	case parser.TokenName(parser.CL_CURLY):
		return "\"}\""
	case parser.TokenName(parser.OP_PARENT):
		return "\"(\""
	case parser.TokenName(parser.CL_PARENT):
		return "\")\""
	case parser.TokenName(parser.OP_SQUARE):
		return "\"[\""
	case parser.TokenName(parser.CL_SQUARE):
		return "\"]\""
	}
	return category
}

func init() {
	parser.PrintToken = printToken
}
