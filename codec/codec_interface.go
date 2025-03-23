// Package codec provides utilities for encoding and decoding images in various formats.
// It defines strategy interfaces for image encoding (ImageEncoder) and decoding (ImageDecoder),
// along with implementations for common formats such as JPEG and PNG.
// This package is designed to support extensible and configurable image processing workflows.
package codec

import (
	"bytes"
	"image"
)

// ImageEncoder defines the strategy interface for image encoding operations.
// Implementations should handle specific image formats.
type ImageEncoder interface {
	// Encode converts an image.Image to its byte representation.
	// Returns:
	//   - []byte: Encoded image data
	//   - error: Encoding failure if any
	Encode(image.Image) ([]byte, error)
}

// ImageDecoder defines the strategy interface for image decoding operations.
type ImageDecoder interface {
	// Decode converts byte data to an image.Image.
	// Returns:
	//   - image.Image: Decoded image
	//   - string: Detected image format
	//   - error: Decoding failure if any
	Decode([]byte) (image.Image, string, error)
}

// DefaultDecoder implements ImageDecoder using Go's standard image decoding.
type DefaultDecoder struct{}

// Decode detects image format and decodes from byte data.
func (defaultDecoder DefaultDecoder) Decode(data []byte) (image.Image, string, error) {
	img, format, err := image.Decode(bytes.NewReader(data))
	return img, format, err
}
