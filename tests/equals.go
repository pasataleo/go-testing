package tests

import (
	"github.com/google/go-cmp/cmp"
)

func (ctx Context[CaptureType]) Equals(expected any) Context[CaptureType] {
	ctx.tb.Helper()

	value, next := ctx.capture()
	if diff := cmp.Diff(expected, value.Value.Interface()); len(diff) > 0 {
		fail(ctx, value.Value.Interface(), diff)
	}

	return next
}

func (ctx Context[CaptureType]) Equalsf(expected any, format string, args ...any) Context[CaptureType] {
	ctx.tb.Helper()

	value, next := ctx.capture()
	if diff := cmp.Diff(expected, value.Value.Interface()); len(diff) > 0 {
		fail(ctx, value.Value.Interface(), format, args...)
	}

	return next
}
