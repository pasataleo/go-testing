package tests

import "reflect"

type Capture[CaptureType any] struct {
	Context[CaptureType]

	Value CaptureType
}

func setCaptureType[CaptureType any](capture Capture[reflect.Value]) (Capture[CaptureType], bool) {
	ctx := Context[CaptureType]{
		tb:   capture.tb,
		fail: capture.fail,
		args: capture.args,
		ix:   capture.ix,
	}

	value, ok := capture.Value.Interface().(CaptureType)
	if !ok {
		return Capture[CaptureType]{
			Context: ctx,
		}, false
	}

	return Capture[CaptureType]{
		Context: ctx,
		Value:   value,
	}, true
}

func (ctx Context[CaptureType]) Capture() (Capture[CaptureType], Context[CaptureType]) {
	ctx.tb.Helper()

	raw, next := ctx.capture()

	capture, ok := setCaptureType[CaptureType](raw)
	if !ok {
		ctx.tb.Fatalf("invalid type, requested %T but held %T", *new(CaptureType), raw.Value.Interface())
	}

	return capture, next
}
