package tests

import "reflect"

func (ctx Context[CaptureType]) Empty() Context[CaptureType] {
	ctx.tb.Helper()
	return ctx.Emptyf("%v", ActualFormatKey)
}

func (ctx Context[CaptureType]) NotEmpty() Context[CaptureType] {
	ctx.tb.Helper()
	return ctx.NotEmptyf("%s", ActualFormatKey)
}

func (ctx Context[CaptureType]) Emptyf(format string, args ...any) Context[CaptureType] {
	ctx.tb.Helper()

	value, next := ctx.capture()
	if !empty(value.Value) {
		fail(ctx, value.Value.Interface(), format, args...)
	}

	return next
}

func (ctx Context[CaptureType]) NotEmptyf(format string, args ...any) Context[CaptureType] {
	ctx.tb.Helper()

	value, next := ctx.capture()
	if empty(value.Value) {
		fail(ctx, "(empty)", format, args...)
	}

	return next
}

func empty(value reflect.Value) bool {
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}
