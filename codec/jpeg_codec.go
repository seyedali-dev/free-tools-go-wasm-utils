// Package codec provides utilities for encoding and decoding images in various formats.
// It defines strategy interfaces for image encoding (ImageEncoder) and decoding (ImageDecoder),
// along with implementations for common formats such as JPEG and PNG.
// This package is designed to support extensible and configurable image processing workflows.
package codec

import (
	"bytes"
	"image"
	"image/jpeg"
)

// JPEGEncoder implements ImageEncoder for JPEG format with configurable quality.
// Fields:
//   - Quality: JPEG compression quality (1-100)
type JPEGEncoder struct{ Quality int }

// Encode converts image to JPEG bytes using specified quality setting.
func (e JPEGEncoder) Encode(img image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, img, &jpeg.Options{Quality: e.Quality})
	return buf.Bytes(), err
}
