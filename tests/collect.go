package tests

import (
	"reflect"
)

func (ctx Context[CaptureType]) Collect(args ...any) Context[CaptureType] {
	var values []reflect.Value
	for _, arg := range args {
		values = append(values, reflect.ValueOf(arg))
	}

	return Context[CaptureType]{
		tb:   ctx.tb,
		fail: ctx.fail,
		args: values,
		ix:   len(values) - 1,
	}
}
