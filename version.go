// This package provides a workaround for the problem that ldflags does not work during `go install` and cannot display embedded versions.
package version

import (
	"runtime/debug"
)

// Return runtime version of main.
// If failed to get version, return false.
func Version() (string, bool) {
	v, ok := debug.ReadBuildInfo()
	if !ok {
		return "", ok
	}
	return v.Main.Version, true
}
