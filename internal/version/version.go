package version

import (
	"fmt"
	"runtime"
)

// Version of the binaries
const Version = "0.0.1"

// String returns our formatted version string
func String(app string) string {
	return fmt.Sprintf("%s v%s (built w/%s)", app, Version, runtime.Version())
}
