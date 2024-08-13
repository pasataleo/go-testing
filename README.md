# pasataleo/go-testing

This is my testing library. It is a work in progress.

## Installation

```bash
go get github.com/pasataleo/go-testing
```

## Usage

```go
package main

import (
    "testing"

    "github.com/pasataleo/go-testing/tests"
)

func FunctionUnderTest(value string) string {
	return value
}

func FunctionWithErrorUnderTest(value string) (string, error) {
    return value, nil
}

func TestFunctionUnderTest(t *testing.T) {
	// First, we can test the function returns the expected value.
	tests.Execute(FunctionUnderTest("test")).Equal(t, "test")
	
	// Second, we can also automatically test the function returns no error, before validating the expected value.
	tests.Execute2E(FunctionWithErrorUnderTest("test")).NoError(t).Equal(t, "test")
}
```
