package tests

func (output *RunOutput) error(raw interface{}) (error, bool) {
	output.tb.Helper()

	value, ok := raw.(error)
	if !ok {
		output.fail("expected error type but found %T", raw)
	}
	return value, ok
}

func (output RunOutput) CaptureError() (error, RunOutput) {
	output.tb.Helper()

	if output, ok := output.validate(); !ok {
		return nil, output
	}

	value := output.value()
	if value == nil {
		return nil, output
	}

	if err, ok := output.error(value); ok {
		return err, output
	}
	return nil, output
}

func (output RunOutput) NoError() RunOutput {
	output.tb.Helper()

	err, output := output.CaptureError()
	if output.tb.Failed() {
		return output
	}

	if err != nil {
		output.fail("expected no error but found \"%s\"", err.Error())
	}

	return output
}

func (output RunOutput) Error(expected string) RunOutput {
	output.tb.Helper()

	err, output := output.CaptureError()
	if output.tb.Failed() {
		return output
	}

	if err == nil {
		output.fail("expected \"%s\" but found no error", expected)
	} else if err.Error() != expected {
		output.fail("expected \"%s\" but was \"%s\"", expected, err.Error())
	}

	return output
}
