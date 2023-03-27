package internal

import (
	"reflect"
)

type CallResults struct {
	ReturnValues []reflect.Value

	Success               bool
	ErrorMessageFmt       string
	ErrorMessageArguments []interface{}
}

func success(values []reflect.Value) CallResults {
	return CallResults{
		ReturnValues: values,
		Success:      true,
	}
}

func failure(format string, args ...any) CallResults {
	return CallResults{
		Success:               false,
		ErrorMessageFmt:       format,
		ErrorMessageArguments: args,
	}
}

func Call(function reflect.Value, args []reflect.Value) CallResults {

	convert := func(have reflect.Value, want reflect.Type) (reflect.Value, bool) {
		if !have.IsValid() {
			have = reflect.New(want).Elem()
		}

		if !have.CanConvert(want) {
			return have, false
		}

		return have.Convert(want), true
	}

	if function.Type().IsVariadic() {
		var arguments []reflect.Value
		for ix := 0; ix < function.Type().NumIn()-1; ix++ {
			if ix >= len(args) {
				return failure("expected at least %d arguments for variadic function but found %d", function.Type().NumIn()-1, len(args))
			}
			expected := function.Type().In(ix)
			value, ok := convert(args[ix], expected)
			if !ok {
				return failure("cannot convert input %d (%s) into %s", ix, value.Type(), expected)
			}
			arguments = append(arguments, value)
		}

		variadicIndex := function.Type().NumIn() - 1
		variadicType := function.Type().In(variadicIndex)
		elementType := variadicType.Elem()

		switch {
		case variadicIndex == len(args):
			// There are no variadic arguments.
			arguments = append(arguments, reflect.MakeSlice(variadicType, 0, 0))
		default:
			// There is at least one.
			argument := reflect.MakeSlice(variadicType, 0, len(args)-variadicIndex)
			for ix := variadicIndex; ix < len(args); ix++ {
				value, ok := convert(args[ix], elementType)
				if !ok {
					return failure("cannot convert input %d (%s) into %s", ix, value.Type(), variadicType)
				}
				argument = reflect.Append(argument, value)
			}
			arguments = append(arguments, argument)
		}

		return success(function.CallSlice(arguments))
	} else {
		if len(args) != function.Type().NumIn() {
			return failure("expected %d arguments but found %d", function.Type().NumIn(), len(args))
		}

		var arguments []reflect.Value
		for ix, argument := range args {
			expected := function.Type().In(ix)
			value, ok := convert(argument, expected)
			if !ok {
				return failure("cannot convert input %d (%s) into %s", ix, value.Type(), expected)
			}
			arguments = append(arguments, value)
		}
		return success(function.Call(arguments))
	}
}
