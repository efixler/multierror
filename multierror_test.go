package multierror

import (
	"errors"
	"fmt"
	"testing"
)

func Example() {
	err := New()
	err = append(err, errors.New("Add errors to MultiError by append-ing"))
	fmt.Println(err)
	err = make(MultiError, 0)
	err = append(err, errors.New("It's fine to make MultiErrors with make"))
	err = append(err, errors.New("Error() will let you know how many errors are contained."))
	fmt.Printf("%s", err)

	// Output:
	// Add errors to MultiError by append-ing
	// It's fine to make MultiErrors with make (and 1 other error)
}

// in case a fork-er gets confused about typed nils in Go
func TestNilWhenEmpty(t *testing.T) {
	var err error
	me := New()
	err = me.NilWhenEmpty()
	if err != nil {
		t.Errorf("NilWhenEmpty() should have returned a nil error, but returned a %T instead (len: %d)",
			err, len(me))
	}
}
