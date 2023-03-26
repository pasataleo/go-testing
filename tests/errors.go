package tests

import (
	"reflect"
	"strings"

	"github.com/google/go-cmp/cmp"
)

var (
	err = reflect.TypeOf((*error)(nil)).Elem()
)

func (ctx Context[CaptureType]) CaptureError() (Capture[error], Context[CaptureType]) {
	ctx.tb.Helper()

	capture, rest := ctx.capture()
	if !capture.Value.Type().ConvertibleTo(err) {
		ctx.tb.Fatalf("expected error type but found %s", capture.Value.Kind())
	}

	if capture.Value.Interface() == nil {
		return Capture[error]{
			Context: switchCaptureType[CaptureType, error](ctx),
			Value:   nil,
		}, rest
	}

	return Capture[error]{
		Context: switchCaptureType[CaptureType, error](ctx),
		Value:   capture.Value.Interface().(error),
	}, rest
}

func (ctx Context[CaptureType]) NoError() Context[CaptureType] {
	ctx.tb.Helper()
	return ctx.NoErrorf("%s", ActualFormatKey)
}

func (ctx Context[CaptureType]) Error() Context[CaptureType] {
	ctx.tb.Helper()
	return ctx.Errorf("%s", ActualFormatKey)
}

func (ctx Context[CaptureType]) ErrorMatches(expected string) Context[CaptureType] {
	ctx.tb.Helper()
	return ctx.ErrorMatchesf(expected, "%s", ActualFormatKey)
}

func (ctx Context[CaptureType]) NoErrorf(format string, args ...any) Context[CaptureType] {
	ctx.tb.Helper()

	err, next := ctx.CaptureError()
	if err.Value != nil {
		fail(ctx, err.Value.Error(), format, args...)
	}
	return next
}

func (ctx Context[CaptureType]) Errorf(format string, args ...any) Context[CaptureType] {
	ctx.tb.Helper()

	err, next := ctx.CaptureError()
	if err.Value == nil {
		fail(ctx, "(nil)", format, args...)
	}
	return next
}

func (ctx Context[CaptureType]) ErrorMatchesf(expected string, format string, args ...any) Context[CaptureType] {
	ctx.tb.Helper()

	err, next := ctx.CaptureError()
	if err.Value == nil {
		fail(ctx, "(nil)", format, args...)
		fmt := strings.ReplaceAll(format, ActualFormatKey, "(nil)")
		ctx.fail(fmt, args...)
	}

	if diff := cmp.Diff(expected, err.Value.Error()); len(diff) > 0 {
		fail(ctx, err.Value.Error(), format, args...)
	}

	return next
}
