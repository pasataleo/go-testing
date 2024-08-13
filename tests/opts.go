package tests

var (
	// Fatal is an option that means if the test fails, the test will be stopped.
	Fatal = &opt{
		fn: func(context *validator) *validator {
			context.Fatal()
			return context
		},
	}
)

// Opt is an option that can be used to modify the behavior of the underlying validator.
type Opt interface {
	transform(context *validator) *validator
}

var _ Opt = &opt{}

type opt struct {
	fn func(context *validator) *validator
}

func (o *opt) transform(context *validator) *validator {
	return o.fn(context)
}
