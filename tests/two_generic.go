package tests

import (
	"testing"
)

// Execute2 creates a new context with two values. Ideally, you'd just call the function you are testing
// directly: tests.Execute2(struct.function()). This allows you to use the context methods to test the result:
// tests.Execute2(struct.function()).Equal(t, expected).Equal(t, otherExpected).
//
// This context contains two values and provides methods to test equality and differences.
//
// Every method returns a new context with the same first value, allowing validates to be chained together.
func Execute2[A, B any](one A, two B) *Context2[A, B] {
	return &Context2[A, B]{
		one: one,
		two: two,
	}
}

// Context2 is a test context that allows for easy testing of two values. This context contains two values
// and provides methods to test equality and differences.
type Context2[A, B any] struct {
	one A
	two B
}

// value implements the context interface.
func (ctx *Context2[A, B]) value() B {
	return ctx.two
}

// Diff checks that the value is different from the given value.
func (ctx *Context2[A, B]) Diff(t testing.TB, other B, opts ...Opt) *Context1[A] {
	t.Helper()
	genericDiff(t, ctx, other, opts...)
	return &Context1[A]{
		one: ctx.one,
	}
}

// Equal checks that the value is equal to the given value.
func (ctx *Context2[A, B]) Equal(t testing.TB, other B, opts ...Opt) *Context1[A] {
	t.Helper()
	genericEquality(t, ctx, other, opts...)
	return &Context1[A]{
		one: ctx.one,
	}
}

// NotEqual checks that the value is not equal to the given value.
func (ctx *Context2[A, B]) NotEqual(t testing.TB, other B, opts ...Opt) *Context1[A] {
	t.Helper()
	genericNotEqual(t, ctx, other, opts...)
	return &Context1[A]{
		one: ctx.one,
	}
}

// Capture returns the value.
func (ctx *Context2[A, B]) Capture() (B, *Context1[A]) {
	return ctx.two, &Context1[A]{
		one: ctx.one,
	}
}

// Validate runs the given callback with the value.
func (ctx *Context2[A, B]) Validate(t testing.TB, cb func(B)) *Context1[A] {
	t.Helper()
	cb(ctx.two)
	return &Context1[A]{
		one: ctx.one,
	}
}
