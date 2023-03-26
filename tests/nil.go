package tests

func (ctx Context[CaptureType]) IsNil() Context[CaptureType] {
	return ctx.IsNilf("%v", ActualFormatKey)
}

func (ctx Context[CaptureType]) IsNotNil() Context[CaptureType] {
	return ctx.IsNotNilf("%s", ActualFormatKey)
}

func (ctx Context[CaptureType]) IsNilf(format string, args ...any) Context[CaptureType] {
	ctx.tb.Helper()

	value, next := ctx.capture()
	if !value.Value.IsNil() {
		fail(ctx, value.Value.Interface(), format, args...)
	}

	return next
}

func (ctx Context[CaptureType]) IsNotNilf(format string, args ...any) Context[CaptureType] {
	ctx.tb.Helper()

	value, next := ctx.capture()
	if value.Value.IsNil() {
		fail(ctx, "(nil)", format, args...)
	}

	return next
}
