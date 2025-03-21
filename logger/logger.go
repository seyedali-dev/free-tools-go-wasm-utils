package logger

import "syscall/js"

// JSLog writes messages to the JavaScript console in WASM environments.
// It prefixes each message with "[GO_WASM] ::" for identification.
func JSLog(message string) {
	// Access the global JavaScript object and call the console's log method
	// to output the message prefixed with "[GO_WASM] ::".
	js.Global().Get("console").Call("log", "[GO_WASM] :: "+message)
}

// LogLevel represents the different levels of logging severity.
type LogLevel int

// Log levels, in order of increasing severity.
const (
	// Debug level logs all messages, including debug information.
	Debug LogLevel = iota
	// Info level logs informational messages.
	Info
	// Error level logs error messages.
	Error
)

// Logger represents a logging instance with a configurable log level.
type Logger struct {
	// level is the minimum log level for logging messages.
	level LogLevel
}

// New creates a new logger instance.
func New(level LogLevel) *Logger {
	return &Logger{level: level}
}

// Debug logs debug messages if the current log level is Debug or lower.
func (l *Logger) Debug(msg string) {
	// Check if the logger's level allows for Debug messages.
	if l.level <= Debug {
		// Log the message to the JavaScript console with a debug prefix.
		JSLog("[DEBUG] :: " + msg)
	}
}

// Info logs informational messages
// It is used to log messages that are useful for debugging, but might not be
// of interest to end-users.
func (l *Logger) Info(msg string) {
	// Check if the logger's level allows for Info messages.
	if l.level <= Info {
		// Log the message to the JavaScript console with an info prefix.
		JSLog("[INFO] :: " + msg)
	}
}

// Error logs error messages.
// It is used to log error messages that indicate something went wrong.
func (l *Logger) Error(msg string) {
	// Check if the logger's level allows for Error messages.
	if l.level <= Error {
		// Log the message to the JavaScript console with an error prefix.
		JSLog("[ERROR] :: " + msg)
	}
}
