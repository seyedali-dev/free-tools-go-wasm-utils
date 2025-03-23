package shared

import (
	"fmt"
	"syscall/js"

	"github.com/seyedali-dev/free-tools-go-wasm-utils/errors"
	"github.com/seyedali-dev/free-tools-go-wasm-utils/logger"
)

// Logger is the shared instance for logging operations.
// It is used across the application for consistent logging behavior.
var Logger *logger.Logger

// init initializes the shared logger instance.
//
// It sets the logger to the default logging level (Debug).
func init() {
	Logger = logger.New(logger.Debug)
}

func RecoverAndRejectJS(reject js.Value) {
	if r := recover(); r != nil {
		errMsg := errors.Err.WrapErr(fmt.Errorf("(unexpected error: %v)", r)).Error()
		Logger.Error(errMsg)

		reject.Invoke(errMsg)
	}
}
