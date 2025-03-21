package errors

import "fmt"

// CustomError defines a structured error with a code, a message, and an underlying error.
type CustomError struct {
	Code    string
	Message string
	Err     error
}

// Error returns the string representation of the custom error.
// If the underlying error is nil, only the code and message are returned.
// If the underlying error is not nil, the error message is prefixed with the code and message.
func (ce *CustomError) Error() string {
	if ce.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", ce.Code, ce.Message, ce.Err)
	}
	return fmt.Sprintf("[%s] %s", ce.Code, ce.Message)
}

// WrapErr wraps an existing error with a CustomError.
// It appends the underlying error to the custom error.
// This is useful for wrapping errors that have already been handled by the code
// but still needs to be propagated to the caller.
func (ce *CustomError) WrapErr(err error) error {
	ce.Err = err
	return fmt.Errorf("%v: %w", ce.Message, err)
}

// Wrap wraps an existing error with a CustomError.
// It appends the underlying error to the custom error.
// This is useful for wrapping errors that have already been handled by the code
// but still needs to be propagated to the caller.
func Wrap(err error, code, message string) error {
	return &CustomError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// Global error variables for standardized error reporting.
var (
	ErrDecodeImage        = &CustomError{Code: "DecodeImageError", Message: "failed to decode image"}
	ErrEncodeImage        = &CustomError{Code: "EncodeImageError", Message: "failed to encode image"}
	ErrInvalidCropOptions = &CustomError{Code: "InvalidCropOptions", Message: "invalid crop options"}
	ErrUnsupportedFormat  = &CustomError{Code: "UnsupportedFormat", Message: "format is not supported"}
	Err                   *CustomError
)
