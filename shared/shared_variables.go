package shared

import (
	"github.com/seyedali-dev/free-tools-go-wasm-utils/codec"
	"github.com/seyedali-dev/free-tools-go-wasm-utils/logger"
)

// Encoders registry maps image formats to their corresponding encoder strategies.
//
// Supported formats:
//   - "jpeg"/"jpg": JPEG format with 90% quality
//   - "png": PNG format with default compression
var Encoders = map[string]codec.ImageEncoder{
	"jpeg": codec.JPEGEncoder{Quality: 90},
	"jpg":  codec.JPEGEncoder{Quality: 90},
	"png":  codec.PNGEncoder{},
}

// Decoder is the shared instance for image decoding operations.
var Decoder = codec.DefaultDecoder{}

// Logger is the shared instance for logging operations.
// It is used across the application for consistent logging behavior.
var Logger *logger.Logger

// init initializes the shared logger instance.
//
// It sets the logger to the default logging level (Debug).
func init() {
	Logger = logger.New(logger.Debug)
}
