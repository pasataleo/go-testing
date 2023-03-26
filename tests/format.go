package tests

const (
	ActualFormatKey = "%actual%"
)

func fail[CaptureType any](ctx Context[CaptureType], actual interface{}, format string, args ...any) {
	ctx.fail(format, replaceActualFormatKey(actual, args...)...)
}

func replaceActualFormatKey(actual any, args ...any) []any {
	var replaced []any
	for _, arg := range args {
		if arg == ActualFormatKey {
			replaced = append(replaced, actual)
		} else {
			replaced = append(replaced, arg)
		}
	}
	return replaced
}
