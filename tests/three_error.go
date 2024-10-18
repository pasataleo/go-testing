package tests

import (
	"testing"

	"github.com/pasataleo/go-errors/errors"
)

// Execute3E is a helper function that returns a new context with two values and an error. Ideally, you'd just call the
// function you are testing directly: tests.Execute3A(struct.function()). This allows you to use the context methods to
// test the result: tests.Execute3A(struct.function()).NoError(t).Equal(t, expected).Equal(t, expected).
func Execute3E[A any, B any](a A, b B, err error) *Context3E[A, B] {
	return &Context3E[A, B]{
		one:   a,
		two:   b,
		three: err,
	}
}

// Context3E is a test context that allows for easy testing of two values and an error. This context contains two values
// and an error and provides methods to test equality and differences.
type Context3E[A any, B any] struct {
	one   A
	two   B
	three error
}

// value implements the context interface.
func (ctx *Context3E[A, B]) value() error {
	return ctx.three
}

// Error checks that the value is an error.
func (ctx *Context3E[A, B]) Error(t testing.TB, opts ...Opt) *Context2[A, B] {
	t.Helper()
	expectError(t, ctx, opts...)
	return &Context2[A, B]{
		one: ctx.one,
		two: ctx.two,
	}
}

// MatchesError checks that the value is an error and matches the expected error.
func (ctx *Context3E[A, B]) MatchesError(t testing.TB, expected string, opts ...Opt) *Context2[A, B] {
	t.Helper()
	matchesError(t, ctx, expected, opts...)
	return &Context2[A, B]{
		one: ctx.one,
		two: ctx.two,
	}
}

// NoError checks that the value is not an error.
func (ctx *Context3E[A, B]) NoError(t testing.TB, opts ...Opt) *Context2[A, B] {
	t.Helper()
	expectNoError(t, ctx, opts...)
	return &Context2[A, B]{
		one: ctx.one,
		two: ctx.two,
	}
}

// ErrorCode checks that the value is an error and has the expected error code.
func (ctx *Context3E[A, B]) ErrorCode(t testing.TB, expected errors.ErrorCode, opts ...Opt) *Context2[A, B] {
	t.Helper()
	matchesErrorCode(t, ctx, expected, opts...)
	return &Context2[A, B]{
		one: ctx.one,
		two: ctx.two,
	}
}

// Validate runs the given callback with the value.
func (ctx *Context3E[A, B]) Validate(t testing.TB, cb func(error)) *Context2[A, B] {
	t.Helper()
	cb(ctx.three)
	return &Context2[A, B]{
		one: ctx.one,
		two: ctx.two,
	}
}

// Capture returns the value.
func (ctx *Context3E[A, B]) Capture() (error, *Context2[A, B]) {
	return ctx.three, &Context2[A, B]{
		one: ctx.one,
		two: ctx.two,
	}
}
