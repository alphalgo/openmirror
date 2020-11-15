package openmirror

import (
	"testing"

	a "github.com/i0Ek3/asrt"
)

func TestParing(t *testing.T) {
	om := New()
	got := om.Paired()

	assertParing(t, got, nil)
}

func assertParing(t *testing.T, got, want error) {
	a.Asrt(t, got, want)
}
