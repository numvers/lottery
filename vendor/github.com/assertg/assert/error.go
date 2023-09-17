package assert

import "testing"

func NoError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		return
	}
	t.Errorf(`actual: %#v but expected nil`, err)
}
