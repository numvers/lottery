package assert

import "testing"

type singed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func Positive[C singed](t *testing.T, actual C) {
	t.Helper()
	if actual > 0 {
		return
	}
	t.Errorf(`actual: %#v is negative but expected positive`, actual)
}

func Negative[C singed](t *testing.T, actual C) {
	t.Helper()
	if actual < 0 {
		return
	}
	t.Errorf(`actual: %#v is positive but expected negative`, actual)
}
