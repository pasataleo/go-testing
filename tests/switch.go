package tests

func Switch[CaptureType any](ctx Context[interface{}]) Context[CaptureType] {
	return switchCaptureType[interface{}, CaptureType](ctx)
}

func SwitchT[OriginalCaptureType, CaptureType any](ctx Context[OriginalCaptureType]) Context[CaptureType] {
	return switchCaptureType[OriginalCaptureType, CaptureType](ctx)
}
