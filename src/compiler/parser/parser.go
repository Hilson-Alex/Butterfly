package parser

const (
	EOF = iota
	ERROR
	UNKNOWN
)

type token interface {
	TokenType() int
}

type Parser[T token] struct {
	TokBuffer []T
}

func (lex *Parser[T]) Lex(lval *T) int {
	if len(lex.TokBuffer) == 0 {
		return EOF
	}
	*lval = lex.TokBuffer[0]
	lex.TokBuffer = lex.TokBuffer[1:]
	return (*lval).TokenType()
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

// func (lex *Lexer) consume() {
// 	lex.TokBuffer = lex.TokBuffer.Next()
// }
