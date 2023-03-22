package tests

import "testing"

func Test_TwoArgs(t *testing.T) {
	f := func(input string) (string, error) {
		return input, nil
	}

	output, _ := Run(t, f, "value").
		NoError().
		Capture()

	output.Equals("value")
}

func Test_ThreeArgs(t *testing.T) {
	f := func(one string, two string) (string, string, error) {
		return one, two, nil
	}

	output, remainder := Run(t, f, "one", "two").
		NoError().
		Capture()

	remainder.Equals("one")
	output.Equals("two")
}

func Test_ThreeArgsSkip(t *testing.T) {
	f := func(one string, two string) (string, string, error) {
		return one, two, nil
	}

	output, _ := Run(t, f, "one", "two").
		NoError().
		Skip().
		Capture()

	output.Equals("one")
}
