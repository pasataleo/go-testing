package examples

type ExampleStruct struct {
	String  string
	Integer int
}

func One(one string) string {
	return one
}

func Two(one, two string) (string, string) {
	return one, two
}

func Three(one, two, three string) (string, string, string) {
	return one, two, three
}

func OneE(one string, err error) (string, error) {
	return one, err
}

func TwoE(one, two string, err error) (string, string, error) {
	return one, two, err
}

func OneT[Value any](value Value, err error) (Value, error) {
	return value, err
}

func TwoT[One, Two any](one One, two Two, err error) (One, Two, error) {
	return one, two, err
}

func Variadic(args ...string) []string {
	return args
}
