package tests

// This file contains a bunch of interfaces that allow for custom equality and diff checks. These interfaces are used
// in the tests package to allow for custom equality and diff checks.

// EqualityInterface is an interface that allows for custom equality checks.
type EqualityInterface[E any] interface {
	Equals(E) bool
}

// DiffInterface is an interface that allows for custom diff checks.
type DiffInterface[E any] interface {
	Diff(E) string
}
