package tests

import "reflect"

func (ctx Context[CaptureType]) Flatten() Context[CaptureType] {
	ctx.tb.Helper()

	capture, _ := ctx.capture()
	switch capture.Value.Kind() {
	case reflect.Array, reflect.Slice:
		// do nothing
	default:
		ctx.tb.Fatalf("cannot flatten a %s", capture.Value.Kind())
	}

	values := ctx.args[:ctx.ix]
	for ix := 0; ix < capture.Value.Len(); ix++ {
		values = append(values, capture.Value.Index(ix))
	}
	values = append(values, ctx.args[ctx.ix+1:]...)

	return Context[CaptureType]{
		tb:   ctx.tb,
		fail: ctx.fail,
		args: values,
		ix:   ctx.ix + (len(values) - len(ctx.args)),
	}
}

func (ctx Context[CaptureType]) Len(expected int) Context[CaptureType] {
	return ctx.Lenf(expected, "%d", ActualFormatKey)
}

func (ctx Context[CaptureType]) Lenf(expected int, format string, args ...any) Context[CaptureType] {
	ctx.tb.Helper()

	capture, rest := ctx.capture()
	switch capture.Value.Kind() {
	case reflect.Slice, reflect.Array, reflect.Chan, reflect.Map, reflect.String, reflect.Ptr:
		// do nothing
	default:
		ctx.tb.Fatalf("cannot call len on %s", capture.Value.Kind())
	}

	if capture.Value.Len() != expected {
		fail(ctx, capture.Value.Len(), format, args...)
	}

	return rest
}
