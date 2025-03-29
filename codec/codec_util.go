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
	"syscall/js"

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
func ApplyBackgroundColor(img image.Image, bgColor color.Color) image.Image {
	// Check if the image already has no transparency (no need to process)
	if !hasTransparency(img) {
		return img
	}

	// Create a new RGBA image with the background color
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	// Fill with background color first
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			newImg.Set(x, y, bgColor)
		}
	}

	// Then overlay the original image, respecting alpha
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			origColor := img.At(x, y)
			r1, g1, b1, a1 := origColor.RGBA()

			// If fully opaque, just use the original color
			if a1 == 0xffff {
				newImg.Set(x, y, origColor)
				continue
			}

			// If fully transparent, keep the background color
			if a1 == 0 {
				continue
			}

			// Otherwise blend with background
			r2, g2, b2, _ := bgColor.RGBA()

			// Alpha blending formula
			alpha := float64(a1) / 0xffff
			r := uint8((float64(r1)*alpha + float64(r2)*(1-alpha)) / 0x101)
			g := uint8((float64(g1)*alpha + float64(g2)*(1-alpha)) / 0x101)
			b := uint8((float64(b1)*alpha + float64(b2)*(1-alpha)) / 0x101)

			newImg.Set(x, y, color.RGBA{R: r, G: g, B: b, A: 0xff})
		}
	}

	return newImg
}

// hasTransparency detects if an image has transparency.
func hasTransparency(img image.Image) bool {
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			_, _, _, a := img.At(x, y).RGBA()
			if a < 0xffff {
				return true
			}
		}
	}
	return false
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

// CopyBytesToGo takes a JS Uint8Array and returns a go byte slice.
func CopyBytesToGo(jsData js.Value) []byte {
	byteLength := jsData.Get("length").Int()
	fileBytes := make([]byte, byteLength)
	_ = js.CopyBytesToGo(fileBytes, jsData)

	return fileBytes
}

// CopyBytesToJS takes a byte slice and returns a JS Uint8Array.
func CopyBytesToJS(data []byte) js.Value {
	jsArray := js.Global().Get("Uint8Array").New(len(data))
	js.CopyBytesToJS(jsArray, data)

	return jsArray
}
