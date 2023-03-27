package tests

import "reflect"

func (ctx Context[CaptureType]) IsNil() Context[CaptureType] {
	ctx.tb.Helper()
	return ctx.IsNilf("%v", ActualFormatKey)
}

func (ctx Context[CaptureType]) IsNotNil() Context[CaptureType] {
	ctx.tb.Helper()
	return ctx.IsNotNilf("%s", ActualFormatKey)
}

func (ctx Context[CaptureType]) IsNilf(format string, args ...any) Context[CaptureType] {
	ctx.tb.Helper()

	value, next := ctx.capture()
	if !nillable(value.Value) || !value.Value.IsNil() {
		fail(ctx, value.Value.Interface(), format, args...)
	}

	return next
}

func (ctx Context[CaptureType]) IsNotNilf(format string, args ...any) Context[CaptureType] {
	ctx.tb.Helper()

	value, next := ctx.capture()
	if nillable(value.Value) && value.Value.IsNil() {
		fail(ctx, "(nil)", format, args...)
	}

	return next
}

func nillable(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer,
		reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return true
	default:
		return false
	}
}
