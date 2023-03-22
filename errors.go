package tests

import (
	"reflect"
	"strings"
)

var (
	err = reflect.TypeOf((*error)(nil)).Elem()
)

type ErrorOutput struct {
	RunOutput

	Value error
}

func (output RunOutput) error() (ErrorOutput, RunOutput) {
	output.tb.Helper()

	raw, next := output.value()
	if !raw.Type().ConvertibleTo(err) {
		output.fail("expected error type but found %s", raw.Kind())
		return ErrorOutput{
			RunOutput: output,
		}, next
	}

	iface := raw.Interface()
	if iface == nil {
		return ErrorOutput{
			RunOutput: output,
		}, next
	}

	return ErrorOutput{
		RunOutput: output,
		Value:     iface.(error),
	}, next
}

func (output RunOutput) CaptureError() (ErrorOutput, RunOutput) {
	output.tb.Helper()

	if next, ok := output.validate(); !ok {
		return ErrorOutput{
			RunOutput: output,
		}, next
	}

	return output.error()
}

func (output RunOutput) NoError() RunOutput {
	output.tb.Helper()
	return output.NoErrorf(ActualFormatKey)
}

func (output RunOutput) NoErrorf(format string, args ...any) RunOutput {
	output.tb.Helper()

	err, output := output.CaptureError()
	if output.tb.Failed() {
		return output
	}

	if err.Value != nil {
		fmt := strings.ReplaceAll(format, ActualFormatKey, err.Value.Error())
		output.fail(fmt, args...)
	}

	return output
}

func (output RunOutput) ErrorMatches(expected string) RunOutput {
	output.tb.Helper()
	return output.ErrorMatchesf(expected, ActualFormatKey)
}

func (output RunOutput) ErrorMatchesf(expected string, format string, args ...any) RunOutput {
	output.tb.Helper()

	err, output := output.CaptureError()
	if output.tb.Failed() {
		return output
	}

	if err.Value == nil {
		fmt := strings.ReplaceAll(format, ActualFormatKey, "nil")
		output.fail(fmt, args...)
	} else if err.Value.Error() != expected {
		fmt := strings.ReplaceAll(format, ActualFormatKey, err.Value.Error())
		output.fail(fmt, args...)
	}
	return output
}

func (output RunOutput) Error() RunOutput {
	output.tb.Helper()
	return output.Errorf("no error")
}

func (output RunOutput) Errorf(format string, args ...any) RunOutput {
	output.tb.Helper()

	err, output := output.CaptureError()
	if output.tb.Failed() {
		return output
	}

	if err.Value == nil {
		output.fail(format, args...)
	}

	return output
}
