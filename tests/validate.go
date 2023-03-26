package tests

import "fmt"

const (
	ValidateSuccess string = "__success__"
)

type ValidateFn[CaptureType any] func(value CaptureType) bool
type ValidateFnf[CaptureType any] func(value CaptureType) string

func (ctx Context[CaptureType]) Validate(validate ValidateFn[CaptureType]) Context[CaptureType] {
	ctx.tb.Helper()
	return ctx.Validatef(func(value CaptureType) string {
		if !validate(value) {
			return fmt.Sprintf("%v", value)
		}
		return ValidateSuccess
	})
}

func (ctx Context[CaptureType]) Validatef(validate ValidateFnf[CaptureType]) Context[CaptureType] {
	ctx.tb.Helper()

	capture, next := ctx.Capture()
	if msg := validate(capture.Value); msg != ValidateSuccess {
		ctx.fail(msg)
	}

	return next
}
