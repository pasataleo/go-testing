package tests

import "testing"

func expectError(t testing.TB, ctx genericContext[error], opts ...Opt) {
	t.Helper()
	if ctx.value() == nil {
		newValidator(t, opts...).Fail("no error")
	}
}

func expectNoError(t testing.TB, ctx genericContext[error], opts ...Opt) {
	t.Helper()
	if ctx.value() != nil {
		newValidator(t, opts...).Fail("%s", ctx.value())
	}
}

func matchesError(t testing.TB, ctx genericContext[error], expected string, opts ...Opt) {
	t.Helper()
	if ctx.value() == nil {
		newValidator(t, opts...).Fail("no error")
		return
	}

	if ctx.value().Error() != expected {
		newValidator(t, opts...).Fail("error does not match")
	}
}
