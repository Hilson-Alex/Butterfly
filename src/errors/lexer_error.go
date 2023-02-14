package errors

import (
	"fmt"
)

// TokenError is the definition for an error
// while parsing the tokens on the lexer.
type TokenError struct {
	Line, Column, Offset int
	Value, FileName      string
}

// Message for Lexical error.
// The Token must appear exactly as parsed, so I
// used a quoted %s instead of %q. To avoid escapes.
func (unknown *TokenError) Error() string {
	var lastColumn = unknown.Column + unknown.Offset
	return fmt.Sprintf("unknown token for value \"%s\" on line %d and columns from %d to %d on file: %s",
		unknown.Value, unknown.Line, unknown.Column, lastColumn, unknown.FileName,
	)
}
