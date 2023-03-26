package tests

import (
	"reflect"
	"testing"
)

type Context[CaptureType any] struct {
	tb   testing.TB
	fail func(format string, args ...any)

	args []reflect.Value
	ix   int
}

func ExecFn(tb testing.TB, function any, args ...any) Context[interface{}] {
	tb.Helper()
	return ExecFnT[interface{}](tb, function, args...)
}

func ExecFnT[CaptureType any](tb testing.TB, function any, args ...any) Context[CaptureType] {
	tb.Helper()

	f := reflect.ValueOf(function)
	if f.Type().Kind() != reflect.Func {
		tb.Fatalf("expected function type but found %s", f.Type().Kind())
	}

	if len(args) != f.Type().NumIn() {
		tb.Fatalf("expected %d arguments but found %d", f.Type().NumIn(), len(args))
	}

	var values []reflect.Value
	for ix, arg := range args {
		exp := f.Type().In(ix)

		var value reflect.Value
		if arg == nil {
			value = reflect.New(exp).Elem()
		} else {
			value = reflect.ValueOf(arg)
		}

		if !value.CanConvert(exp) {
			tb.Fatalf("cannot convert input %d (%s) into %s", ix, value.Type(), exp)
		}

		values = append(values, value.Convert(exp))
	}

	returns := f.Call(values)
	return Context[CaptureType]{
		tb:   tb,
		fail: tb.Errorf,
		args: returns,
		ix:   len(returns) - 1,
	}
}

func (ctx Context[CaptureType]) Fatal() Context[CaptureType] {
	ctx.fail = ctx.tb.Fatalf
	return ctx
}

func (ctx Context[CaptureType]) NonFatal() Context[CaptureType] {
	ctx.fail = ctx.tb.Errorf
	return ctx
}

func (ctx Context[CaptureType]) Skip() Context[CaptureType] {
	ctx.ix--
	return ctx
}

func (ctx Context[CaptureType]) value() (reflect.Value, Context[CaptureType]) {
	ctx.tb.Helper()

	if ctx.ix < 0 {
		ctx.tb.Fatalf("received %d return values but attempted to process too many", len(ctx.args))
	}

	value := ctx.args[ctx.ix]
	ctx.ix--
	return value, ctx
}

func (ctx Context[CaptureType]) capture() (Capture[reflect.Value], Context[CaptureType]) {
	ctx.tb.Helper()
	raw, next := ctx.value()
	return Capture[reflect.Value]{
		Context: switchCaptureType[CaptureType, reflect.Value](ctx),
		Value:   raw,
	}, next
}

func switchCaptureType[OriginalCaptureType, CaptureType any](ctx Context[OriginalCaptureType]) Context[CaptureType] {
	return Context[CaptureType]{
		tb:   ctx.tb,
		fail: ctx.fail,
		args: ctx.args,
		ix:   ctx.ix,
	}
}
