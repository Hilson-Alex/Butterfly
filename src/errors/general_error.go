package errors

import (
	"fmt"
)

type NestedGenericError struct {
	message string
	nested  []error
}

func (err *NestedGenericError) Error() string {
	if len(err.nested) == 0 {
		return err.message
	}
	var joinedErr = fmt.Sprintf("%s. Nested error is:", err.message)
	for _, nested := range err.nested {
		joinedErr = fmt.Sprintf("%s\n\t%s", joinedErr, nested.Error())
	}
	return joinedErr
}

func CreateNestedErr(message string, nested ...error) *NestedGenericError {
	return &NestedGenericError{
		message, nested,
	}
}
