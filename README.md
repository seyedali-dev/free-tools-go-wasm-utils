# Go WASM Utils

Reusable logging, error, codec (image) handling utilities for Go/WASM projects. Provides structured logging with panic recovery and standardized error handling for WebAssembly environments.

## Features
- 🚀 WASM-friendly logging to JavaScript console
- 🛠️ Multi-level logging (Debug, Info, Error)
- 🧨 Panic recovery for WASM modules
- 📦 Standardized error types with codes
- 🔍 Context-rich error wrapping

## Installation
```bash
go get https://github.com/seyedali-dev/free-tools-go-wasm-utils
```

## Usage

### Logger Initialization
```go
// `internal/shared/logger.go` or anywhere you like!
package shared

import "github.com/seyedali-dev/crop-wasm/internal/logger"

var Logger *logger.Logger

func init() {
    // Initialize with desired log level; Options: Debug, Info, Error
    Logger = logger.New(logger.Debug) // Default log is debug but you can use other levels with the same `Logger` instance
}
```

### Logging Examples
```go
// Debug logging (shows only when log level is Debug)
shared.Logger.Debug("Processing image with dimensions: 800x600")

// Info logging (shows in Info and Debug levels)
shared.Logger.Info("Image cropping started")

// Error logging with error wrapping
err := processImage()
if err != nil {
    shared.Logger.Error(fmt.Sprintf("Processing failed: %v", err))
    return errors.ErrProcessingFailed.Wrap(err)
}
```

### Error Handling Patterns
```go
// Using predefined errors
if err := json.Unmarshal(data, &opts); err != nil {
    return errors.ErrInvalidCropOptions.Wrap(err)
}

// Creating custom errors
myErr := &errors.CustomError{
    Code:    "ImageTooSmall",
    Message: "Image dimensions below minimum requirement",
}.Wrap(err)

// Handling errors with context
result, err := cropImage(opts)
if err != nil {
    return errors.ErrCropFailed.Wrap(fmt.Errorf("crop failed for image %s: %w", imageName, err))
}
```

### Image Codec Examples
```go
var data []byte // base64 encoded image
image, format, err := codec.DefaultDecoder.Decode(data)
if err != nil {
    return errors.ErrDecodeImage.Wrap(err)
}

// Do something with image and format...
```

### WASM Panic Recovery
```go
func cropImageJS(_ js.Value, args []js.Value) interface{} {
    return js.Global().Get("Promise").New(js.FuncOf(func(_ js.Value, promiseArgs []js.Value) interface{} {
        resolve, reject := promiseArgs[0], promiseArgs[1]
        go func() {
            defer func() {
                if r := recover(); r != nil {
                    errMsg := fmt.Sprintf("Panic recovered: %v", r)
                    shared.Logger.Error(errMsg)
                    reject.Invoke(errMsg)
                }
            }()
            
            // ... your processing logic ...
        }()
        return nil
    }))
}
```

## Error Types
Predefined error codes for consistent error handling:

| Error Code            | Description                          | Usage Example               |
|-----------------------|--------------------------------------|-----------------------------|
| `DecodeImageError`    | Image decoding failure               | `errors.ErrDecodeImage`     |
| `EncodeImageError`    | Image encoding failure               | `errors.ErrEncodeImage`     |
| `InvalidCropOptions`  | Invalid cropping parameters          | `errors.ErrInvalidCropOptions` |
| `UnsupportedFormat`   | Unsupported image format             | `errors.ErrUnsupportedFormat` |
| `ProcessingFailed`    | Generic processing failure           | `errors.ErrProcessingFailed` |

## WASM Considerations
1. Logging automatically outputs to browser console with `[WASM]` prefix
2. Panic recovery prevents module crashes
3. Error messages are properly marshaled to JavaScript
4. All operations are async-safe for browser environments

## Contributing
1. Fork repository
2. Create feature branch (`git checkout -b feature/logging-enhancements`)
3. Commit changes (`git commit -am 'Add debug level filtering'`)
4. Push to branch (`git push origin feature/logging-enhancements`)
5. Create Pull Request
