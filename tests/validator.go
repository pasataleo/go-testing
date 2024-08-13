package tests

import (
	"testing"
)

func newValidator(t testing.TB, opts ...Opt) *validator {
	ctx := &validator{
		t:     t,
		fatal: false,
	}

	for _, opt := range opts {
		ctx = opt.transform(ctx)
	}
	return ctx
}

type validator struct {
	t     testing.TB
	fatal bool
}

func (ctx *validator) Fatal() {
	ctx.t.Helper()
	ctx.fatal = true
}

func (ctx *validator) Error() {
	ctx.t.Helper()
	ctx.fatal = false
}

func (ctx *validator) Fail(msg string, args ...any) {
	ctx.t.Helper()
	if ctx.fatal {
		ctx.t.Fatalf(msg, args...)
		return
	}

	ctx.t.Errorf(msg, args...)
	return
}
