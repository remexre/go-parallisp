package debug

import "fmt"

// Log logs to the debug system.
func Log(namespace, format string, others ...interface{}) {
	Chan <- D{namespace, fmt.Sprintf(format, others...)}
}
