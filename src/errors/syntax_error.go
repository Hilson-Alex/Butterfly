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

func CreateSyntaxError(yaccErr, tokenName string, token failedToken, translator func(string) string) *SyntaxError {
	var syntaxErr = &SyntaxError{
		tokenName: tokenName,
		token:     token,
	}
	if strings.Contains(yaccErr, "expecting") {
		syntaxErr.expectedTokens = "\n\t Expecting " + parseExpected(strings.Split(yaccErr, "expecting ")[1], translator)
	}
	return syntaxErr
}

func parseExpected(expected string, translate func(string) string) string {
	const separator = " or "
	var expectedTokens = strings.Split(expected, separator)
	var translated = make([]string, 0, len(expectedTokens))
	for _, token := range expectedTokens {
		translated = append(translated, translate(token))
	}
	return strings.Join(translated, separator)
}

func (syntaxErr *SyntaxError) Error() string {
	var tokenName, token, expectedTokens = syntaxErr.tokenName, syntaxErr.token, syntaxErr.expectedTokens
	return fmt.Sprintf(
		"unexpected token %s %q at line %d and column %d of %s.%s",
		tokenName, token.Value(), token.Line(), token.Column(), token.FileName(), expectedTokens,
	)
}
