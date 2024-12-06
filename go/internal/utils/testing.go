package utils

import "testing"

func ExpectEqInt(got, want int, t *testing.T) {
	if got != want {
		t.Errorf("Got %d instead of %d", got, want)
	}
}
