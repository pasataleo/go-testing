package tests

import (
	"testing"
)

// Execute creates a new context with the given value. Ideally, you'd just call the function you are testing
// directly: tests.Execute(struct.function()). This allows you to use the context methods to test the result:
// tests.Execute(struct.function()).Equal(t, expected).
func Execute[A any](one A) *Context1[A] {
	return &Context1[A]{
		one: one,
	}
}

// Context1 is a test context that allows for easy testing of values. This context contains only a single value
// and provides methods to test equality and differences.
type Context1[A any] struct {
	one A
}

// value implements the context interface.
func (ctx *Context1[A]) value() A {
	return ctx.one
}

// Diff checks that the value is different from the given value.
func (ctx *Context1[A]) Diff(t testing.TB, other A, opts ...Opt) {
	t.Helper()
	genericDiff(t, ctx, other, opts...)
}

// Equal checks that the value is equal to the given value.
func (ctx *Context1[A]) Equal(t testing.TB, other A, opts ...Opt) {
	t.Helper()
	genericEquality(t, ctx, other, opts...)
}

// NotEqual checks that the value is not equal to the given value.
func (ctx *Context1[A]) NotEqual(t testing.TB, other A, opts ...Opt) {
	t.Helper()
	genericNotEqual(t, ctx, other, opts...)
}

// Validate runs the given callback with the value.
func (ctx *Context1[A]) Validate(t testing.TB, cb func(testing.TB, A)) {
	t.Helper()
	cb(t, ctx.one)
}

// Capture returns the value.
func (ctx *Context1[A]) Capture() A {
	return ctx.one
}
