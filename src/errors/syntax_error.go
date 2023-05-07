package errors

import (
	"fmt"
	"strings"
)

type failedToken interface {
	Value() string
	Line() int
	Column() int
	FileName() string
}

type SyntaxError struct {
	expectedTokens string
	tokenName      string
	token          failedToken
}

func CreateSyntaxError(yaccErr, tokenName string, token failedToken) *SyntaxError {
	var syntaxErr = &SyntaxError{
		tokenName: tokenName,
		token:     token,
	}
	if strings.Contains(yaccErr, "expecting") {
		syntaxErr.expectedTokens = "\n\t Expecting " + strings.Split(yaccErr, "expecting ")[1]
	}
	return syntaxErr
}

func (syntaxErr *SyntaxError) Error() string {
	var tokenName, token, expectedTokens = syntaxErr.tokenName, syntaxErr.token, syntaxErr.expectedTokens
	return fmt.Sprintf(
		"unexpected token %s %q at line %d and column %d of %s.%s",
		tokenName, token.Value(), token.Line(), token.Column(), token.FileName(), expectedTokens,
	)
}
