package parser

import (
	"github.com/Hilson-Alex/Butterfly/src/compiler/checker"
	bfErrors "github.com/Hilson-Alex/Butterfly/src/errors"
)

// EOF represents the syntactic value of the End Of File.
const EOF = 0

// Config to get more info on the goyacc generated butterfly.go file errors.
func init() {
	yyErrorVerbose = true
}

// Logger for errors when parsing the syntax.
var logger = bfErrors.NewBfErrLogger("SYNTAX ERROR:")

var PrintToken = func(code int) string {
	return TokenName(code)
}

// token interface breaks the lexer dependency, so there's no circular
// imports. The lexer needs to import the parser to get the Token Codes,
// but the parser don't need to import the lexer to access its info.
type token interface {
	TokenType() int
	Value() string
	Line() int
	Column() int
	FileName() string
}

// parser is a facade to communicate to the goyacc generated file
type parser[T token] struct {
	TokBuffer    []T
	lastToken    T
	success      bool
	currentScope *checker.BFScope
	result       string
}

type ParserResult struct {
	ModuleName, Result string
	Success            bool
}

// Lex passes the next token to the goyacc generated parser
func (lex *parser[T]) Lex(lval *yySymType) int {
	if len(lex.TokBuffer) == EOF {
		return EOF
	}
	var currentToken = lex.TokBuffer[0]
	lex.lastToken = currentToken
	lval.yys = currentToken.TokenType()
	lval.content = currentToken.Value()
	lval.currentToken = currentToken
	lval.scope = lex.currentScope
	lval.result = &lex.result
	lex.TokBuffer = lex.TokBuffer[1:]
	return lval.yys
}

// Error logs any syntactic error
func (lex *parser[T]) Error(error string) {
	lex.success = false
	var lastToken = lex.lastToken
	logger.Log(bfErrors.CreateSyntaxError(error, PrintToken(lastToken.TokenType()), lastToken))
}

func Parse[T token](lexer []T) ParserResult {
	var parserInfo = &parser[T]{TokBuffer: lexer, success: true, currentScope: new(checker.BFScope)}
	yyParse(parserInfo)
	return ParserResult{
		ModuleName: parserInfo.currentScope.Module().Name(),
		Result:     parserInfo.result,
		Success:    parserInfo.success,
	}
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
