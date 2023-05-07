package errors

import (
	"fmt"
)

type SemanticError struct {
	Line int
	Info string
}

func (semanticErr *SemanticError) Error() string {
	return fmt.Sprintf("invalid state at line %d. %s", semanticErr.Line, semanticErr.Info)
}
