package tests

import "testing"

// Execute2E is a helper function that returns a new context with a value and an error. Ideally, you'd just call the
// function you are testing directly: tests.Execute2E(struct.function()). This allows you to use the context methods to
// test the result: tests.Execute2E(struct.function()).NoError(t).Equal(t, expected).
func Execute2E[A any](a A, err error) *Context2E[A] {
	return &Context2E[A]{
		one: a,
		two: err,
	}
}

// Context2E is a test context that allows for easy testing of a value and an error. This context contains a value and
// an error and provides methods to test equality and differences.
type Context2E[A any] struct {
	one A
	two error
}

// value implements the context interface.
func (ctx *Context2E[A]) value() error {
	return ctx.two
}

// Error checks that the value is an error.
func (ctx *Context2E[A]) Error(t testing.TB, opts ...Opt) *Context1[A] {
	t.Helper()
	expectError(t, ctx, opts...)
	return &Context1[A]{
		one: ctx.one,
	}
}

// MatchesError checks that the value is an error and matches the expected error.
func (ctx *Context2E[A]) MatchesError(t testing.TB, expected string, opts ...Opt) *Context1[A] {
	t.Helper()
	matchesError(t, ctx, expected, opts...)
	return &Context1[A]{
		one: ctx.one,
	}
}

// NoError checks that the value is not an error.
func (ctx *Context2E[A]) NoError(t testing.TB, opts ...Opt) *Context1[A] {
	t.Helper()
	expectNoError(t, ctx, opts...)
	return &Context1[A]{
		one: ctx.one,
	}
}

// Validate runs the given callback with the value.
func (ctx *Context2E[A]) Validate(t testing.TB, cb func(error)) *Context1[A] {
	t.Helper()
	cb(ctx.two)
	return &Context1[A]{
		one: ctx.one,
	}
}
