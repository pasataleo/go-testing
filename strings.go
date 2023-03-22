package tests

import (
	"reflect"
	"strings"
)

var (
	str = reflect.TypeOf((*string)(nil)).Elem()
)

type StringOutput struct {
	RunOutput

	Value string
}

func (output RunOutput) string() (StringOutput, RunOutput) {
	output.tb.Helper()

	raw, next := output.value()
	if raw.Kind() != reflect.String {
		output.fail("expected string type but found %s", raw.Kind())
		return StringOutput{
			RunOutput: output,
		}, next
	}
	return StringOutput{
		RunOutput: output,
		Value:     raw.Interface().(string),
	}, next
}

func (output RunOutput) CaptureString() (StringOutput, RunOutput) {
	output.tb.Helper()

	if next, ok := output.validate(); !ok {
		return StringOutput{
			RunOutput: output,
		}, next
	}

	return output.string()
}

func (output RunOutput) StringEquals(expected string) RunOutput {
	output.tb.Helper()
	return output.StringEqualsf(expected, ActualFormatKey)
}

func (output RunOutput) StringEqualsf(expected string, format string, args ...any) RunOutput {
	output.tb.Helper()

	value, next := output.CaptureString()
	if output.tb.Failed() {
		return next
	}

	if value.Value != expected {
		fmt := strings.ReplaceAll(format, ActualFormatKey, value.Value)
		output.fail(fmt, args...)
	}

	return next
}

func (output RunOutput) StringLessThan(other string) RunOutput {
	output.tb.Helper()
	return output.StringLessThanf(other, ActualFormatKey)
}

func (output RunOutput) StringLessThanf(other string, format string, args ...any) RunOutput {
	output.tb.Helper()

	value, next := output.CaptureString()
	if output.tb.Failed() {
		return next
	}

	if value.Value < other {
		fmt := strings.ReplaceAll(format, ActualFormatKey, value.Value)
		output.fail(fmt, args...)
	}

	return next
}
