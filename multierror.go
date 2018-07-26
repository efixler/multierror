// MultiError lets you wrap multiple errors into one error. Use this for iterators
// that don't necessarily bail when they hit an error, but want to return all of the
// encountered errors to the caller.
//
// MultiError is a typed []error that implements golang's error interface. As such, you can use slice
// operations to modify or inspect the list of contained errors. Error() attempts to provide
// coherent stringification based on the number of errors in the MultiError.
//
package multierror

import (
	"fmt"
)

var (
	// The string output by Error() when the list has no errors.
	NoErrorsMessage = "(no errors)"
)

type MultiError []error

func New(errors ...error) MultiError {
	me := make(MultiError, len(errors))
	for i := 0; i < len(errors); i++ {
		me[i] = errors[i]
	}
	return me
}

// The error message returned by Error() always contains the error message of
// the first error in the list (if there is one). If len(MultiError) > 1
// the number of errors in the list is also noted.
// If zero, the NoErrorsMessage variable is used.
func (m MultiError) Error() string {
	lm := len(m)
	switch lm {
	case 0:
		return NoErrorsMessage
	case 1:
		return m[0].Error()
	case 2:
		return fmt.Sprintf("%s (and 1 other error)", m[0].Error())
	default:
		return fmt.Sprintf("%s (and %d other errors)", m[0].Error(), lm-1)
	}
}

// Returns the original MultiError if there's at least one error inside it,
// or nil if there isn't. Saves you from having to check len(MultiError) when
// returning from your function.
//
//	func Worker() error {
//	   merr := make(MultiError, 0)
//	   for _, item := range work {
//		 err := doSomething()
//		 if err != nil {
//		   merr = append(merr, err)
//		 }
//	    }
//	    return merr.NilWhenEmpty()
//	 }
//
//
func (m MultiError) NilWhenEmpty() error {
	switch len(m) {
	case 0:
		return nil
	default:
		return m
	}
}
