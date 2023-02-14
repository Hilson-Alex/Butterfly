package errors

import (
	"fmt"
)

// NestedGenericError adds a custom user-friendly message
// while stacks go errors to give context or detailed info.
type NestedGenericError struct {
	message string
	nested  []error
}

// CreateNestedErr create an error with stacked nested errors
func CreateNestedErr(message string, nested ...error) *NestedGenericError {
	return &NestedGenericError{
		message, nested,
	}
}

// implements error interface
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
