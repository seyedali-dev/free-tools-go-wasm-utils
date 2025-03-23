package codecstrategy

import (
	"image"
	"image/color"
	"image/png"
	"io"

	"github.com/seyedali-dev/free-tools-go-wasm-utils/codec"
	"github.com/seyedali-dev/free-tools-go-wasm-utils/codec/types"
)

// PNGCodec implements codec.ImageCodec for PNG format.
type PNGCodec struct{}

// Encode encodes a PNG image to the given writer. It takes types.CompressionLevel as an option to specify the compression level.
func (pngCodec *PNGCodec) Encode(writer io.Writer, img image.Image, options map[string]interface{}) error {
	encoder := png.Encoder{}
	encoder.CompressionLevel = png.DefaultCompression

	// Check if CompressionLevel option is set
	if compLvlInfc, ok := options[types.CompressionLevel]; ok {
		if level, ok := compLvlInfc.(png.CompressionLevel); ok {
			encoder.CompressionLevel = level
		}
	}

	// Apply background color
	bgColor := color.Color(color.White) // default white
	if bgColorInfc, ok := options[types.BackgroundColor]; ok {
		if bgColorStr, ok := bgColorInfc.(string); ok {
			if bgColorHex, err := codec.ParseHexColor(bgColorStr); err == nil {
				bgColor = bgColorHex
			}
		}
	}
	img = codec.ApplyBackgroundColor(img, bgColor)

	return encoder.Encode(writer, img)
}

// Decode decodes a PNG image from the given reader. It uses the png.Decode function.
func (pngCodec *PNGCodec) Decode(reader io.Reader, _ map[string]interface{}) (image.Image, error) {
	return png.Decode(reader)
}
