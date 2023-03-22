package tests

import (
	"reflect"
	"testing"
)

const (
	ActualFormatKey = "%actual%"
)

type RunOutput struct {
	tb   testing.TB
	fail func(format string, args ...any)

	args []reflect.Value
	ix   int
}

func Run(tb testing.TB, function any, args ...any) RunOutput {
	tb.Helper()

	f := reflect.ValueOf(function)
	if f.Type().Kind() != reflect.Func {
		tb.Fatalf("invalid function: %s", f.Type().Kind())
	}

	if len(args) != f.Type().NumIn() {
		tb.Fatalf("expected %d arguments but found %d", f.Type().NumIn(), len(args))
	}

	var values []reflect.Value
	for _, arg := range args {
		values = append(values, reflect.ValueOf(arg))
	}

	returns := f.Call(values)
	return RunOutput{
		tb:   tb,
		fail: tb.Errorf,
		args: returns,
		ix:   len(returns) - 1,
	}
}

func (output RunOutput) Fatal() RunOutput {
	output.fail = output.tb.Fatalf
	return output
}

func (output RunOutput) Skip() RunOutput {
	output.ix--
	return output
}

func (output RunOutput) Capture() (ValueOutput, RunOutput) {
	output.tb.Helper()

	if next, ok := output.validate(); !ok {
		return ValueOutput{
			RunOutput: output,
		}, next
	}

	raw, next := output.value()
	return ValueOutput{
		RunOutput: output,
		Value:     raw.Interface(),
	}, next
}

func (output RunOutput) validate() (RunOutput, bool) {
	output.tb.Helper()

	if output.ix < 0 {
		output.fail("received %d return values but attempted to process %d too many", len(output.args), abs(output.ix))
		output.ix--
		return output, false
	}

	return output, true
}

func (output RunOutput) value() (reflect.Value, RunOutput) {
	value := output.args[output.ix]
	output.ix--
	return value, output
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
