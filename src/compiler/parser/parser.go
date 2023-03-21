package parser

import (
	"fmt"
	"strings"

	bfErrors "github.com/Hilson-Alex/Butterfly/src/errors"
)

const (
	EOF = iota
)

func init() {
	yyErrorVerbose = true
}

var logger = bfErrors.NewBfErrLogger("SYNTAX ERROR:")

type token interface {
	TokenType() int
	Value() string
	Line() int
	Column() int
	FileName() string
}

type Parser[T token] struct {
	TokBuffer []T
	lastToken T
}

func (lex *Parser[T]) Lex(lval *yySymType) int {
	if len(lex.TokBuffer) == EOF {
		return EOF
	}
	lex.lastToken = lex.TokBuffer[0]
	lval.yys = lex.lastToken.TokenType()
	lex.TokBuffer = lex.TokBuffer[1:]
	return lval.yys
}

func (lex *Parser[T]) Error(error string) {
	var lastToken = lex.lastToken
	var expectedTokens = ""
	if strings.Contains(error, "expecting") {
		expectedTokens = "\n\t Expecting " + strings.Split(error, "expecting ")[1]
	}
	var errorMessage = fmt.Sprintf(
		"unexpected token %s %q at line %d and column %d of %s.%s",
		TokenName(lastToken.TokenType()), lastToken.Value(), lastToken.Line(),
		lastToken.Column(), lastToken.FileName(), expectedTokens,
	)
	logger.Log(errorMessage)
}

func Parse[T token](lexer []T) {
	yyParse(&Parser[T]{TokBuffer: lexer})
}

func TokenName(code int) string {
	var index = 2
	if code >= yyPrivate {
		if code < yyPrivate+len(yyTok2) {
			index = int(yyTok2[code-yyPrivate])
		}
	}
	return yyTokname(index)
}
