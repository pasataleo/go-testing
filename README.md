# pasataleo/go-testing

This is my testing library.

## Test Contexts

This library functions by building and returning `tests.Context` objects. 
Users validate the chain of objects within a context in reverse order.

### Generating Contexts

Contexts are created using the `tests.ExecFn` function: 
`tests.ExecFn(t, MyFunction, myArgOne, myArgTwo)`.

The first argument to `tests.ExecFn` should be a `testing.T` or `testing.B` 
pointer, from the default Go `testing` package.

The returned context will contain all the results of calling `MyFunction` with
`myArgOne` and `myArgTwo`. The types of the arguments are validated dynamically
and will fail the surrounding test using the provided `testing` obejct if the
types are not as expected.

### Validating Contexts

Once you have a `Context`, you can validate the returned values it contains in
reverse order.

Imagine a function returns `(string, error)`, you could validate that there is 
no returned error and the string equals an expected value with: 
`tests.ExecFn(t, MyFunction, myArgOne, myArgTwo).NoError().Equals("expected value")`

The library provides an equality function (`.Equals(...)`) that uses 
[go-cmp](https://github.com/google/go-cmp) under the hood.

The library provides `NoError()` and `Error()` functions for validating an 
expected or unexpected error, alongside an `ErrorMatches(err string)` function 
that uses [go-cmp](https://github.com/google/go-cmp) to validate that an error 
was returned and calling Error() on that error matches the provided string.

Finally, the library provides `Nil()` and `NotNil()` functions for validating 
whether a returned pointer value is or isn't nil.

Further custom validations are explored under the [Custom Validations](#custom-validation) header.

### Custom error messages

All functions have an equivalent format function (eg. `Equalsf` for `Equals`) 
that allows callers to provide custom error messages. To print out the string
value of the value under validation use `tests.ActualFormatKey`, for example 
`.Equalsf("expected %s, but found %s", expectedValue, tests.ActualFormatKey)` 
would print "expected $expectedValue but found $actualValue".

### E2E Examples

```go
package my_library

import (
	"testing"
	
	"github.com/pasataleo/go-testing/tests"
)

func TestSimple(t *testing.T) {
	// A simple function that just returns the provided string value.
	f := func(value string) string {
		return value
    }
	
	// Execute the function, and validate the returned value matches the
	// provided value.
	tests.ExecFn(t, f, "value").Equals("value")
}

func TestComplex(t *testing.T) {
	// Slightly more complex function that returns two provided values, and does
	// not error.
	f := func(one string, two int) (string, int, error) {
		return one, two, nil
    }
	
	// Validate the test function returns the expected values, note the reverse
	// order that the returned values are validated in.
	tests.ExecFn(t, f, "value", 0).
		NoError().
		Equals(0).
		Equals("value")
}

```

## Test Captures

Sometimes you don't only want to validate the return values, but also want to
use them in later functions. We provide the `Capture()` and `ExecFnT()` 
functions to enable this.

While validating a chain of results in a `Context` you can break the chain and
return a `Capture` containing the value that would have been validating and a
new `Context` that will continue validating the result of the chain.

The `Capture` itself implements `Context` so the same set of validation 
functions are available to validate the value held in the `Capture`. Crucially,
the `Capture` also contains a `.Value` field that users can retrieve to actually
use the held value.

By default, as in when using the usual `ExecFn` function, the values within any
returned function will be `interface{}` types that need to be cast into the 
values required/expected by the user. The `ExecFnT` function allows the user
to specify the type of values that will be held within any returned captures:

```go
capture, _ := tests.ExecFn(t, MyFunction, myArgOne).Capture() // capture.Value == interface{}
capture, _ := tests.ExecFnT[string](t, MyFunction, myArgOne).Capture() // capture.Value == string
```

If the type specified by `ExecFnT` does not match the type actually held by the
`Context` then the test will perform an immediate fail.

### Capturing different types

A single `Context` with a concrete capture type can only return captures that 
match the concrete type, so a function that returns multiple value types is
unable to simply return captures for all the different types. For example:

```go
myFunction := func() (string, int) { return "value", 0 }
capture, rest := tests.ExecFnT[int](t, myFunction).Capture() // capture.Value == int
_, _ := rest.Capture() // error, expected int type but context held string type
```

In order to counter this we provide the `Switch` and `SwitchT` functions, that
change the capture type for a given context. `Switch` can be called against 
contexts that were created with the simple `Run(...)` function while `SwitchT`
can change the capture type for contexts that were created with `RunT(...)`:

```go
myFunction := func() (string, int) { return "value", 0 }
capture, rest := tests.ExecFnT[int](t, myFunction).Capture() // capture.Value == int
capture2, _ := tests.SwitchT[int, string](rest).Capture() // capture2.Value == string
```

### E2E Examples

```go
package my_library

import (
	"testing"
	
	"github.com/pasataleo/go-testing/tests"
)

func TestSimple(t *testing.T) {
	// A simple function that just returns the provided string values.
	f := func(one, two string) (string, string) {
		return one, two
	}
	
	// We'll capture the second value ("two") and simply validate the first one.
	capture, rest := tests.ExecFn(t, f, "one", "two").Capture()
	capture.Equals("two")
	rest.Equals("one")
	
	// capture.Value is an interface{} type, as we didn't use ExecFnT to create
	// a concrete capture type.
	var value interface{}
	value = capture.Value
}

func TestComplex(t *testing.T) {
	// Slightly more complex function that returns two provided of different 
	// types values, and does not error.
	f := func(one string, two int) (string, int, error) {
		return one, two, nil
	}
	
	// Validate the test function returns the expected values, note the reverse
	// order that the returned values are validated in.
	capture, rest := tests.ExecFnT[int](t, f, "value", 0).
		NoError().
		Capture()
	
	// The captured value is 0, and the actual type is concretized to int.
	capture.Equals(0)
	var two int
	two = capture.Value
	
	// rest contains the remaining values, in this case it is just the single
	// remaining string "value". The equality check passes as it doesn't care
	// about the types.
	rest.Equals("value")
	
	// Trying to immediately capture will fail as the capture is still set to
	// capture int values but we are holding a string.
	captureFailed, _ := rest.Capture() // Immediate fail.
	
	// Instead, switch the type and then capture the string value.
	captureSuccess, _ := tests.SwitchT[int, string](rest).Capture()
	var one string
	one = captureSuccess.Value
}
```

## Custom Validations

In addition to the provided validation functions (listed under the 
[Validating Contexts](#validating-contexts) heading), you can use the `Validate`
function for complex or custom validation operations. For example, validating if
a string has given prefix or validating only a single entry in a returned 
structure.

The `Validate` function accepts a `ValidateFn` argument which accepts the value
currently held by the context and returns a boolean value. True if the 
validation was successful and the test should pass, and false if the validation 
failed and the test should fail. Also provided is a ValidateFnf function, this
behaviour differs from the usual equivalent format functions. The returned value 
of the provided validate function is a string which should be 
`tests.ValidateSuccess` if the value matches your expectations, otherwise the 
returned string will be used as the error message when reporting the failure.

As with the capture functionality, the validate function will operate on 
`interface{}` types by default. Callers can also use the `ExecFnT` function to 
tell the library to handle the casting the generic types to expected concrete
types. This comes with the same caveats as capturing values, in that you will
need to use `Switch` and `SwitchT` for multiple different validations.

### E2E Examples

```go
package my_library

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pasataleo/go-testing/tests"
)

func TestSimple(t *testing.T) {
	// A simple function that just returns the provided string values.
	f := func(one, two string) (string, string) {
		return one, two
	}

	// We'll validate the returned strings using custom logic, not the validate
	// functions are expected to provide `interface{}` types.
	tests.ExecFn(t, f, "one", "two").
		Validate(func(value interface{}) bool {
			return value.(string) == "two"
		}).
		Validate(func(value interface{}) bool {
			return value.(string) == "one"
		})
}

func TestComplex(t *testing.T) {
	// Slightly more complex function that returns two provided of different 
	// types values, and does not error.
	f := func(one string, two int) (string, int, error) {
		return one, two, nil
	}

	// As with the simple example, we call Validate but this time the first 
	// function is already expecting an int type.
	ctx := tests.ExecFnT[int](t, f, "value", 0).
		NoError().
		Validate(func(value int) bool {
			return value == 0 // validate success if the value matches.
		})

	// Trying to call Validate again, as with the simple example, would fail as
	// the types don't match. Instead, we have to switch the types.
	//
	// For this function we are showing the Validatef behavior, instead of a
	// boolean we are returning a string. If the validation passes we use 
	// tests.ValidateSuccess, otherwise we return the string we want to display
	// in the failure message.
	tests.SwitchT[int, string](ctx).Validatef(func(value string) string {
		if diff := cmp.Diff(value, "value"); len(diff) > 0 {
			return diff
        }
		return tests.ValidateSuccess
	})
}
```
