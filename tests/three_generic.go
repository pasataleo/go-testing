package tests

import "testing"

// Execute3 creates a new context with the given values. Ideally, you'd just call the function you are testing
// directly: tests.Execute3(struct.function()). This allows you to use the context methods to test the result:
// tests.Execute3(struct.function()).Equal(t, expected).Equal(t, otherExpected).Equal(t, anotherExpected).
//
// This context contains three values and provides methods to test equality and differences.
//
// Every method returns a new context with the same first two values, allowing validates to be chained together.
func Execute3[A, B, C any](one A, two B, three C) *Context3[A, B, C] {
	return &Context3[A, B, C]{
		one:   one,
		two:   two,
		three: three,
	}
}

// Context3 is a test context that allows for easy testing of three values. This context contains three values
// and provides methods to test equality and differences.
type Context3[A, B, C any] struct {
	one   A
	two   B
	three C
}

// value implements the context interface.
func (ctx *Context3[A, B, C]) value() C {
	return ctx.three
}

// Diff checks that the value is different from the given value.
func (ctx *Context3[A, B, C]) Diff(t testing.TB, other C, opts ...Opt) *Context2[A, B] {
	t.Helper()
	genericDiff(t, ctx, other, opts...)
	return &Context2[A, B]{
		one: ctx.one,
		two: ctx.two,
	}
}

// Equal checks that the value is equal to the given value.
func (ctx *Context3[A, B, C]) Equal(t testing.TB, other C, opts ...Opt) *Context2[A, B] {
	t.Helper()
	genericEquality(t, ctx, other, opts...)
	return &Context2[A, B]{
		one: ctx.one,
		two: ctx.two,
	}
}

// NotEqual checks that the value is not equal to the given value.
func (ctx *Context3[A, B, C]) NotEqual(t testing.TB, other C, opts ...Opt) *Context2[A, B] {
	t.Helper()
	genericNotEqual(t, ctx, other, opts...)
	return &Context2[A, B]{
		one: ctx.one,
		two: ctx.two,
	}
}

// Capture returns the value.
func (ctx *Context3[A, B, C]) Capture() (C, *Context2[A, B]) {
	return ctx.three, &Context2[A, B]{
		one: ctx.one,
		two: ctx.two,
	}
}

// Validate runs the given callback with the value.
func (ctx *Context3[A, B, C]) Validate(t testing.TB, cb func(C)) *Context2[A, B] {
	t.Helper()
	cb(ctx.three)
	return &Context2[A, B]{
		one: ctx.one,
		two: ctx.two,
	}
}
