package tests

import (
	"github.com/google/go-cmp/cmp"
)

type ValueOutput struct {
	RunOutput

	Value interface{}
}

func (output RunOutput) Equals(expected any) RunOutput {
	output.tb.Helper()

	value, next := output.Capture()
	if output.tb.Failed() {
		return next
	}

	if diff := cmp.Diff(expected, value.Value); len(diff) > 0 {
		output.fail(diff)
	}

	return next
}
