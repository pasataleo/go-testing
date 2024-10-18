package tests

import (
	"testing"

	"github.com/pasataleo/go-errors/errors"
)

// ExecuteE creates a new context with the given error. Ideally, you'd just call the function you are testing
// directly: tests.ExecuteE(struct.function()). This allows you to use the context methods to test the result:
// tests.ExecuteE(struct.function()).NoError(t).
func ExecuteE(err error) *Context1E {
	return &Context1E{one: err}
}

// Context1E is a test context that allows for easy testing of errors. This context contains only a single error
// and provides methods to test for errors.
type Context1E struct {
	one error
}

// error implements the errorContext interface.
func (ctx *Context1E) value() error {
	return ctx.one
}

// Error checks that the value is an error.
func (ctx *Context1E) Error(t testing.TB, opts ...Opt) {
	t.Helper()
	expectError(t, ctx, opts...)
}

// MatchesError checks that the value is an error and matches the expected error.
func (ctx *Context1E) MatchesError(t testing.TB, expected string, opts ...Opt) {
	t.Helper()
	matchesError(t, ctx, expected, opts...)
}

// NoError checks that the value is not an error.
func (ctx *Context1E) NoError(t testing.TB, opts ...Opt) {
	t.Helper()
	expectNoError(t, ctx, opts...)
}

// ErrorCode checks that the value is an error and has the expected error code.
func (ctx *Context1E) ErrorCode(t testing.TB, expected errors.ErrorCode, opts ...Opt) {
	t.Helper()
	matchesErrorCode(t, ctx, expected, opts...)
}

// Validate runs the given callback with the value.
func (ctx *Context1E) Validate(t testing.TB, cb func(error)) {
	t.Helper()
	cb(ctx.one)
}

// Capture returns the value.
func (ctx *Context1E) Capture() error {
	return ctx.one
}
