package tests

import (
	"github.com/pasataleo/go-objects/objects"
)

func (ctx Context[CaptureType]) CaptureObject() (Capture[objects.Object], Context[CaptureType]) {
	ctx.tb.Helper()

	raw, next := ctx.capture()

	capture, ok := setCaptureType[objects.Object](raw)
	if !ok {
		ctx.tb.Fatalf("invalid type, requested %T but held %T", *new(CaptureType), raw.Value.Interface())
	}

	return capture, next
}

func (ctx Context[CaptureType]) ObjectEquals(expected any) Context[CaptureType] {
	ctx.tb.Helper()
	return ctx.ObjectEqualsf(expected, "%s", ActualFormatKey)
}

func (ctx Context[CaptureType]) ObjectEqualsf(expected any, format string, args ...any) Context[CaptureType] {
	ctx.tb.Helper()

	obj, next := ctx.CaptureObject()
	if !obj.Value.Equals(expected) {
		fail(ctx, obj.Value.String(), format, args...)
	}
	return next
}
