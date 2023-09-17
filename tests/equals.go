package tests

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Equals(t *testing.T, expected, actual any) {
	t.Helper()
	if diff := cmp.Diff(expected, actual); len(diff) > 0 {
		t.Errorf("expected: %v, actual: %v, diff: %s", expected, actual, diff)
	}
}

func Equalsf(t *testing.T, expected, actual any, format string, args ...any) {
	t.Helper()
	if diff := cmp.Diff(expected, actual); len(diff) > 0 {
		t.Errorf(format, replaceActualFormatKey(actual, args...)...)
	}
}

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
