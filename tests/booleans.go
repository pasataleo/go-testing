package tests

import "reflect"

func (ctx Context[CaptureType]) True() Context[CaptureType] {
	ctx.tb.Helper()
	return ctx.Truef("%v", ActualFormatKey)
}

func (ctx Context[CaptureType]) False() Context[CaptureType] {
	ctx.tb.Helper()
	return ctx.Falsef("%v", ActualFormatKey)
}

func (ctx Context[CaptureType]) Truef(format string, args ...any) Context[CaptureType] {
	ctx.tb.Helper()
	value, next := ctx.capture()
	if value.Value.Kind() != reflect.Bool {
		fail(ctx, value.Value.Interface(), format, args...)
		return next
	}

	if !value.Value.Bool() {
		fail(ctx, false, format, args...)
	}

	return next
}

func (ctx Context[CaptureType]) Falsef(format string, args ...any) Context[CaptureType] {
	ctx.tb.Helper()
	value, next := ctx.capture()
	if value.Value.Kind() != reflect.Bool {
		fail(ctx, value.Value.Interface(), format, args...)
		return next
	}

	if value.Value.Bool() {
		fail(ctx, true, format, args...)
	}

	return next
}
