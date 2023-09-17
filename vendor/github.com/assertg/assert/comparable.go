package assert

import (
	"testing"
)

func Equals[C comparable](t *testing.T, actual, expect C) {
	t.Helper()
	if actual == expect {
		return
	}
	t.Errorf(`actual: %#v but expected: %#v`, expect, actual)
}
