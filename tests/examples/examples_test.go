package examples

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/pasataleo/go-testing/tests"
)

func TestOne(t *testing.T) {
	tests.ExecFn(t, One, "value").Equals("value")
}

func TestOneWithCapture(t *testing.T) {
	capture, _ := tests.ExecFn(t, One, "value").Capture()
	capture.Equals("value")

	var value interface{}
	value = capture.Value
	_ = fmt.Sprintf("%v", value)
}

func TestOneWithCaptureT(t *testing.T) {
	capture, _ := tests.ExecFnT[string](t, One, "value").Capture()
	capture.Equals("value")

	var value string
	value = capture.Value
	_ = fmt.Sprintf("%s", value)
}

func TestTwo(t *testing.T) {
	tests.ExecFn(t, Two, "one", "two").
		Equals("two").
		Equals("one")
}

func TestTwoWithCapture(t *testing.T) {
	capture, _ := tests.ExecFn(t, Two, "one", "two").
		Equals("two").
		Capture()
	capture.Equals("one")

	var value interface{}
	value = capture.Value
	_ = fmt.Sprintf("%v", value)
}

func TestTwoWithCaptureT(t *testing.T) {
	capture, _ := tests.ExecFnT[string](t, Two, "one", "two").
		Equals("two").
		Capture()
	capture.Equals("one")

	var value string
	value = capture.Value
	_ = fmt.Sprintf("%s", value)
}

func TestTwoWithCaptureTFirst(t *testing.T) {
	capture, rest := tests.ExecFnT[string](t, Two, "one", "two").
		Capture()
	capture.Equals("two")
	rest.Equals("one")

	var value string
	value = capture.Value
	_ = fmt.Sprintf("%s", value)
}

func TestThree(t *testing.T) {
	tests.ExecFn(t, Three, "one", "two", "three").
		Equals("three").
		Equals("two").
		Equals("one")
}

func TestThreeWithSkip(t *testing.T) {
	tests.ExecFn(t, Three, "one", "two", "three").
		Equals("three").
		Skip().
		Equals("one")
}

func TestOneE(t *testing.T) {
	tests.ExecFn(t, OneE, "value", nil).
		NoError().
		Equals("value")
}

func TestOneEWithError(t *testing.T) {
	tests.ExecFn(t, OneE, "value", errors.New("anything")).
		Error().
		Equals("value")
}

func TestOneEWithErrorMatches(t *testing.T) {
	tests.ExecFn(t, OneE, "value", errors.New("specific")).
		ErrorMatches("specific").
		Equals("value")
}

func TestTwoE(t *testing.T) {
	tests.ExecFn(t, TwoE, "one", "two", nil).
		NoError().
		Equals("two").
		Equals("one")
}

func TestTwoTMultipleTypes(t *testing.T) {
	tests.ExecFn(t, TwoT[string, int], "value", 0, nil).
		NoError().
		Equals(0).
		Equals("value")
}

func TestTwoTMultipleTypesWithCapture(t *testing.T) {
	capture, rest := tests.ExecFnT[int](t, TwoT[string, int], "value", 0, nil).
		NoError().
		Capture()
	capture.Equals(0)
	rest.Equals("value")
}

func TestTwoTMultipleTypesWithMultipleCaptureTypes(t *testing.T) {
	capture, rest := tests.ExecFnT[int](t, TwoT[string, int], "value", 0, nil).
		NoError().
		Capture()
	capture.Equals(0)

	var intValue int
	intValue = capture.Value
	_ = fmt.Sprintf("%d", intValue)

	capture2, _ := tests.SwitchT[int, string](rest).Capture()
	capture2.Equals("value")

	var stringValue string
	stringValue = capture2.Value
	_ = fmt.Sprintf("%s", stringValue)
}

func TestTwoTMultipleTypesWithCaptureError(t *testing.T) {
	capture, rest := tests.ExecFn(t, TwoT[string, int], "value", 0, nil).
		CaptureError()
	capture.NoError()
	rest.Equals(0).Equals("value")
}

func TestCustomType(t *testing.T) {
	value := ExampleStruct{
		String:  "value",
		Integer: 1,
	}
	tests.ExecFn(t, OneT[ExampleStruct], value, nil).
		NoError().
		Equals(value)
}

func TestCustomTypePtr(t *testing.T) {
	value := ExampleStruct{
		String:  "value",
		Integer: 1,
	}
	capture, _ := tests.ExecFn(t, OneT[*ExampleStruct], &value, nil).
		NoError().
		Capture()

	capture.IsNotNil()
	capture.Equals(&value)
}

func TestCustomTypePtrNil(t *testing.T) {
	capture, _ := tests.ExecFn(t, OneT[*ExampleStruct], nil, nil).
		NoError().
		Capture()

	capture.IsNil()
}

func TestCustomTypeValidate(t *testing.T) {
	value := ExampleStruct{
		String:  "value",
		Integer: 1,
	}
	tests.ExecFn(t, OneT[ExampleStruct], value, nil).
		NoError().
		Validate(func(actual interface{}) bool {
			if diff := cmp.Diff(actual, value); len(diff) > 0 {
				return false // validate failed.
			}
			return true
		})
}

func TestCustomTypeSwitchAndValidate(t *testing.T) {
	value := ExampleStruct{
		String:  "value",
		Integer: 1,
	}
	capture, rest := tests.ExecFn(t, OneT[ExampleStruct], value, nil).
		CaptureError()
	capture.NoError()

	tests.Switch[ExampleStruct](rest).Validate(func(actual ExampleStruct) bool {
		if diff := cmp.Diff(actual.String, value.String); len(diff) > 0 {
			return false
		}
		return true
	})
}

func TestCustomTypeValidateT(t *testing.T) {
	value := ExampleStruct{
		String:  "value",
		Integer: 1,
	}
	tests.ExecFnT[ExampleStruct](t, OneT[ExampleStruct], value, nil).
		NoError().
		Validatef(func(actual ExampleStruct) string {
			if diff := cmp.Diff(actual.String, value.String); len(diff) > 0 {
				return diff
			}
			return tests.ValidateSuccess
		})
}

func TestVariadic(t *testing.T) {
	tests.ExecFn(t, Variadic).Equals([]string{})
	tests.ExecFn(t, Variadic, "one").Equals([]string{"one"})
	tests.ExecFn(t, Variadic, "one", "two").Equals([]string{"one", "two"})
}
