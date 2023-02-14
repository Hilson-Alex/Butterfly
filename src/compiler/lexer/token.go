package lexer

import (
	"fmt"

	"github.com/Hilson-Alex/Butterfly/src/compiler/parser"
)

type Token struct {
	tokenType int
	value     string
	line      int
	column    int
}

func (token *Token) TokenType() int {
	return token.tokenType
}

func (token *Token) Value() string {
	return token.value
}

func (token *Token) Line() int {
	return token.line
}

func (token *Token) Column() int {
	return token.column
}

func (token *Token) String() string {
	return fmt.Sprintf("{code: %d, value: %s, line: %d, column: %d, tokenClass: %s}",
		token.tokenType, token.value, token.line, token.column, parser.TokenName(token.tokenType),
	)
}