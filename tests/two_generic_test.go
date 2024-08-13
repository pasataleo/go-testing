package tests

import (
	"errors"
	"testing"
)

func fn2[A any, B any](a A, b B) (A, B) {
	return a, b
}

func TestExecute2(t *testing.T) {
	Execute2(fn2(1, 2)).Equal(t, 2).Equal(t, 1)

	value, rest := Execute2(fn2(1, 2)).Capture()
	rest.Equal(t, 1)
	if value != 2 {
		t.Error("not equal")
	}

	Execute2E(fn2[int, error](1, nil)).NoError(t).Equal(t, 1)
	Execute2E(fn2(1, errors.New("error"))).MatchesError(t, "error").Equal(t, 1)
}
