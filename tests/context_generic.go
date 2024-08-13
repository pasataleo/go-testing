package tests

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type genericContext[A any] interface {
	value() A
}

func genericDiff[A any](t testing.TB, ctx genericContext[A], want A, opts ...Opt) {
	t.Helper()

	if value, ok := any(want).(DiffInterface[A]); ok {
		if diff := value.Diff(ctx.value()); len(diff) > 0 {
			newValidator(t, opts...).Fail(diff)
		}
		return
	}

	if diff := cmp.Diff(ctx.value(), want); len(diff) > 0 {
		newValidator(t, opts...).Fail(diff)
	}
}

func genericEquality[A any](t testing.TB, ctx genericContext[A], want A, opts ...Opt) {
	t.Helper()

	if value, ok := any(want).(EqualityInterface[A]); ok {
		if !value.Equals(ctx.value()) {
			newValidator(t, opts...).Fail("not equal; expected %v, got %v", want, ctx.value())
		}
		return
	}

	// fall back to default equality check
	if !cmp.Equal(ctx.value(), want) {
		newValidator(t, opts...).Fail("not equal; expected %v, got %v", want, ctx.value())
	}
}

func genericNotEqual[A any](t testing.TB, ctx genericContext[A], want A, opts ...Opt) {
	t.Helper()

	if value, ok := any(want).(EqualityInterface[A]); ok {
		if value.Equals(ctx.value()) {
			newValidator(t, opts...).Fail("equal; got %v", ctx.value())
		}
		return
	}

	// fall back to default equality check
	if cmp.Equal(ctx.value(), want) {
		newValidator(t, opts...).Fail("equal; got %v", ctx.value())
	}
}
