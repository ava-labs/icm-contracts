package utils

import (
	"os"

	"github.com/ava-labs/libevm/log"
)

// ConfigureDefaultLoggingForTests sets up the default logger to show INFO level and above.
// This should be called once at the beginning of test suites to ensure log.Info() statements are visible.
func ConfigureDefaultLoggingForTests() {
	// Configure log level to ensure Info logs are visible in tests
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelInfo, true)))
}
