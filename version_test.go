package version

import (
	"testing"
)

func TestVersion(t *testing.T) {
	v := Version()

	if v != "" {
		t.Errorf("Version() should return empty string, but got %s", v)
	}
}
