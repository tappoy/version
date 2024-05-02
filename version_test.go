package version

import (
	"testing"
)

func TestVersion(t *testing.T) {
	v, ok := Version()
	if !ok {
		t.Errorf("Version() should return true, but got false")
	}

	if v != "" {
		t.Errorf("Version() should return (development), but got <%s>", v)
	}
}
