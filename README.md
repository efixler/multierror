# multierror
Go package implementing an error that can contain multiple errors

[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]

[godocs]: https://godoc.org/github.com/efixler/multierror

`multierror` wraps multiple errors into a typed slice that implements Go's `error` interface. It's useful
for iterators that don't bail when they encounter an error, but still need to return error information
to the caller.

`multierror.MultiError` aims toward leanness and is mainly utilized with a combination of `error` and 
`slice` semantics.


## Installation

`go get github.com/efixler/multierror`

## Usage

````
import (
	github.com/efixler/multierror
)

func Worker() error {
   merr := multierror.New()
   for _, item := range work {
	 err := doSomething()
	 if err != nil {
	   merr = append(merr, err)
	 }
    }
    return merr.NilWhenEmpty()
 }
 ````

See the [Godoc]() for more details 

