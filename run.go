package tests

import "testing"

type RunOutput struct {
	tb   testing.TB
	fail func(format string, args ...any)

	args []any
	ix   int
}

func Run(tb testing.TB, args ...any) RunOutput {
	return RunE(tb, len(args), args...)
}

func RunE(tb testing.TB, expected int, args ...any) RunOutput {
	if len(args) > expected {
		tb.Helper()
		tb.Fatalf("expected %d return values but found %d", expected, len(args))
	}

	return RunOutput{
		tb:   tb,
		fail: tb.Errorf,
		args: args,
		ix:   len(args) - 1,
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

func (output RunOutput) validate() (RunOutput, bool) {
	output.tb.Helper()

	if output.ix < 0 {
		output.fail("received %d return values but attempted to process %d too many", len(output.args), abs(output.ix))
		output.ix--
		return output, false
	}

	return output, true
}

func (output *RunOutput) value() interface{} {
	value := output.args[output.ix]
	output.ix--
	return value
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
