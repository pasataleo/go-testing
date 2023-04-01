package tests

import (
	"reflect"
	"testing"

	"github.com/pasataleo/go-testing/tests/internal"
)

type Context[CaptureType any] struct {
	tb   testing.TB
	fail func(format string, args ...any)

	args []reflect.Value
	ix   int
}

func Empty(tb testing.TB) Context[interface{}] {
	return EmptyT[interface{}](tb)
}

func EmptyT[CaptureType any](tb testing.TB) Context[CaptureType] {
	return Context[CaptureType]{
		tb:   tb,
		fail: tb.Errorf,
		ix:   -1,
	}
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

	var arguments []reflect.Value
	for _, arg := range args {
		arguments = append(arguments, reflect.ValueOf(arg))
	}

	result := internal.Call(f, arguments)
	if !result.Success {
		tb.Fatalf(result.ErrorMessageFmt, result.ErrorMessageArguments...)
	}

	return Context[CaptureType]{
		tb:   tb,
		fail: tb.Errorf,
		args: result.ReturnValues,
		ix:   len(result.ReturnValues) - 1,
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
