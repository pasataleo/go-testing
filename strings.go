package tests

func (output *RunOutput) string() (string, bool) {
	output.tb.Helper()

	raw := output.value()
	value, ok := raw.(string)
	if !ok {
		output.fail("expected string type but found %T", raw)
	}
	return value, ok
}

func (output RunOutput) CaptureString() (string, RunOutput) {
	output.tb.Helper()

	var value string
	if output, ok := output.validate(); !ok {
		return value, output
	}

	if value, ok := output.string(); ok {
		return value, output
	}

	return value, output
}

func (output RunOutput) StringEquals(expected string) RunOutput {
	output.tb.Helper()

	if output, ok := output.validate(); !ok {
		return output
	}

	value, output := output.CaptureString()
	if output.tb.Failed() {
		return output
	}

	if value != expected {
		output.fail("expected \"%s\" but was \"%s\"", expected, value)
	}

	return output
}
