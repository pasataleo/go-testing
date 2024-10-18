package tests

import "github.com/google/go-cmp/cmp"

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

func CmpOption(opt cmp.Option) Opt {
	return &cmpOpt{opt: opt}
}

type cmpOpt struct {
	opt cmp.Option
}

func (o *cmpOpt) transform(context *validator) *validator {
	// this option does not modify the context
	return context
}

func filterCmpOpts(opt ...Opt) ([]Opt, []cmp.Option) {
	var opts []Opt
	var cmpOpts []cmp.Option
	for _, o := range opt {
		if co, ok := o.(*cmpOpt); ok {
			cmpOpts = append(cmpOpts, co.opt)
		} else {
			opts = append(opts, o)
		}
	}
	return opts, cmpOpts
}
