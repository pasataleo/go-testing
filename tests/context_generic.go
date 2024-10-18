package tests

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pasataleo/go-objects/objects"
)

type genericContext[A any] interface {
	value() A
}

func genericDiff[A any](t testing.TB, ctx genericContext[A], want A, opts ...Opt) {
	t.Helper()
	opts, cmpOpts := filterCmpOpts(opts...)

	if value, ok := any(want).(DiffInterface[A]); ok {
		if diff := value.Diff(ctx.value()); len(diff) > 0 {
			newValidator(t, opts...).Fail(diff)
		}
		return
	}

	if diff := cmp.Diff(ctx.value(), want, cmpOpts...); len(diff) > 0 {
		newValidator(t, opts...).Fail(diff)
	}
}

func genericEquality[A any](t testing.TB, ctx genericContext[A], want A, opts ...Opt) {
	t.Helper()
	opts, cmpOpts := filterCmpOpts(opts...)

	// First, check if the value implements the EqualityInterface.
	if value, ok := any(want).(EqualityInterface[A]); ok {
		if !value.Equals(ctx.value()) {
			newValidator(t, opts...).Fail("not equal; expected %v, got %v", want, ctx.value())
		}
		return
	}

	// Next, check if the value is an object.Object which means it has its own equals function.
	if object, ok := any(want).(objects.Object); ok {
		if !object.Equals(ctx.value()) {
			newValidator(t, opts...).Fail("not equal; expected %v, got %v", want, ctx.value())
		}
		return
	}

	// Finally, fall back to the generic equality check.
	if !cmp.Equal(ctx.value(), want, cmpOpts...) {
		newValidator(t, opts...).Fail("not equal; expected %v, got %v", want, ctx.value())
	}
}

func genericNotEqual[A any](t testing.TB, ctx genericContext[A], want A, opts ...Opt) {
	t.Helper()
	opts, cmpOpts := filterCmpOpts(opts...)

	// First, check if the value implements the EqualityInterface.
	if value, ok := any(want).(EqualityInterface[A]); ok {
		if value.Equals(ctx.value()) {
			newValidator(t, opts...).Fail("equal; got %v", ctx.value())
		}
		return
	}

	// Next, check if the value is an object.Object which means it has its own equals function.
	if object, ok := any(want).(objects.Object); ok {
		if object.Equals(ctx.value()) {
			newValidator(t, opts...).Fail("equal; got %v", ctx.value())
		}
		return
	}

	// fall back to default equality check
	if cmp.Equal(ctx.value(), want, cmpOpts...) {
		newValidator(t, opts...).Fail("equal; got %v", ctx.value())
	}
}
