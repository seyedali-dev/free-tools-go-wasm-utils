// Package codec provides utilities for encoding and decoding images in various formats.
// It defines strategy interfaces for image encoding (ImageEncoder) and decoding (ImageDecoder),
// along with implementations for common formats such as JPEG and PNG.
// This package is designed to support extensible and configurable image processing workflows.
package codec

import (
	"bytes"
	"image"
	"image/png"
)

// PNGEncoder implements ImageEncoder for PNG format with default compression.
type PNGEncoder struct{}

// Encode converts image to PNG bytes using lossless compression.
func (e PNGEncoder) Encode(img image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	return buf.Bytes(), err
}
