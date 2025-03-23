package codec

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/draw"
	"strings"

	"github.com/seyedali-dev/free-tools-go-wasm-utils/errors"
)

// Base64ImgToImg converts a base64-encoded image string to an image.Image.
func Base64ImgToImg(base64Image string) (image.Image, error) {
	if len(base64Image) == 0 {
		return nil, errors.Err.WrapErr(fmt.Errorf("base64 image data is empty"))
	}

	baseImg, _, err := decodeImage(base64Image)
	if err != nil {
		return nil, errors.Err.WrapErr(err)
	}

	return normalizeImage(baseImg), nil
}

// decodeImage handles Base64 decoding and image format detection.
//
// Parameters:
//   - data: Base64 encoded image string (optionally with data URI)
//
// Returns:
//   - image.Image: Decoded image
//   - string: Detected image format
//   - error: Decoding failure if any
func decodeImage(data string) (image.Image, string, error) {
	cleanData := sanitizeBase64(data)
	imgBytes, err := base64.StdEncoding.DecodeString(cleanData)
	if err != nil {
		return nil, "", errors.ErrDecodeImage.WrapErr(err)
	}

	decodedImg, format, err := Decode(bytes.NewBuffer(imgBytes))
	if err != nil || decodedImg == nil {
		return nil, "", errors.ErrDecodeImage.WrapErr(fmt.Errorf("error !!! %w", err))
	}

	return decodedImg, format, nil
}

// sanitizeBase64 removes data URI prefix from Base64 strings if present.
//
// Parameters:
//   - data: Raw input string
//
// Returns:
//   - string: Clean Base64 data without URI prefix
func sanitizeBase64(data string) string {
	if idx := strings.Index(data, ","); idx != -1 {
		return data[idx+1:]
	}
	return data
}

// normalizeImage converts any image format to RGBA for consistent processing.
//
// Parameters:
//   - img: Source image to convert
//
// Returns:
//   - *image.RGBA: Converted image in RGBA format
func normalizeImage(img image.Image) *image.RGBA {
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)
	return rgba
}
