package openmirror

import (
	"testing"
)

func TestParing(t *testing.T) {
	om := New()
	got := om.Paired()

	assertParing(t, got, nil)
}

func assertParing(t *testing.T, got, want error) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
