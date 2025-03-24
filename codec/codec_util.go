package codec

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"strconv"
	"strings"

	"github.com/seyedali-dev/free-tools-go-wasm-utils/errors"
)

// Base64ImgToImg converts a base64-encoded image string to an image.Image.
func Base64ImgToImg(base64Image string) (image.Image, string, error) {
	if len(base64Image) == 0 {
		return nil, "", errors.Err.WrapErr(fmt.Errorf("base64 image data is empty"))
	}

	baseImg, format, err := decodeImage(base64Image)
	if err != nil {
		return nil, "", errors.Err.WrapErr(err)
	}

	return normalizeImage(baseImg), format, nil
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

// ApplyBackgroundColor applies a background color to the source image.
func ApplyBackgroundColor(img image.Image, bg color.Color) image.Image {
	// Create a new RGBA image with the same dimensions
	bounds := img.Bounds()
	bgImg := image.NewRGBA(bounds)

	// Fill the background with the specified color
	draw.Draw(bgImg, bounds, &image.Uniform{C: bg}, image.Point{}, draw.Src)

	// Draw the source image on top
	draw.Draw(bgImg, bounds, img, bounds.Min, draw.Over)

	return bgImg
}

// ParseHexColor converts a hex string to a color.Color.
func ParseHexColor(s string) (color.Color, error) {
	// Validate the color string.
	if s == "" {
		// return default white color if not provided.
		return color.White, nil
	}
	if s[0] == '#' {
		s = s[1:]
	}
	if len(s) != 6 {
		return nil, errors.ErrParseHexColor.WrapErr(fmt.Errorf("color string must be 6 characters"))
	}

	// Parse string to RGB colors.
	r, err := strconv.ParseUint(s[0:2], 16, 8)
	if err != nil {
		return nil, errors.ErrParseHexColor.WrapErr(fmt.Errorf("error occurred parsing [0:2] color hex: %w", err))
	}

	g, err := strconv.ParseUint(s[2:4], 16, 8)
	if err != nil {
		return nil, errors.ErrParseHexColor.WrapErr(fmt.Errorf("error occurred parsing [2:4] color hex: %w", err))
	}

	b, err := strconv.ParseUint(s[4:6], 16, 8)
	if err != nil {
		return nil, errors.ErrParseHexColor.WrapErr(fmt.Errorf("error occurred parsing [4:6] color hex: %w", err))
	}

	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}, nil
}

// NormalizeJPEGFormat normalizes the format to a JPEG family image formats if the input format is JPEG.
func NormalizeJPEGFormat(format SupportedCodecFormat) SupportedCodecFormat {
	switch format {
	case JPEG, JPG, JFIF, JP2, JPEGXR, JPE:
		return JPEG
	default:
		return format
	}
}
