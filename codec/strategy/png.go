package codecstrategy

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"

	"github.com/seyedali-dev/free-tools-go-wasm-utils/codec"
	"github.com/seyedali-dev/free-tools-go-wasm-utils/codec/types"
	"github.com/seyedali-dev/free-tools-go-wasm-utils/errors"
)

// PNGCodec implements codec.ImageCodec for PNG format.
type PNGCodec struct{}

// Encode encodes a PNG image to the given writer.
// It takes types.CompressionLevel and types.BackgroundColor as an option.
//
// It either sets the options or none. There is no default.
func (pngCodec *PNGCodec) Encode(writer io.Writer, img image.Image, options map[string]interface{}) error {
	encoder := png.Encoder{}

	// Check if CompressionLevel option is set
	if options[types.CompressionLevel] != nil {
		if compLvlIntfc, ok := options[types.CompressionLevel]; ok {
			if level, ok := compLvlIntfc.(png.CompressionLevel); ok {
				encoder.CompressionLevel = level
			} else {
				return errors.ErrInvalidArgument.WrapErr(fmt.Errorf("invalid compression level: %+v", compLvlIntfc))
			}
		} else {
			return errors.ErrInvalidArgument.WrapErr(fmt.Errorf("compression level is not set"))
		}
	}

	// Apply background color
	var bgColor color.Color
	if options[types.BackgroundColor] != nil {
		if bgColorIntfc, ok := options[types.BackgroundColor]; ok {
			if bgColorStr, ok := bgColorIntfc.(string); ok {
				if bgColorHex, err := codec.ParseHexColor(bgColorStr); err == nil {
					bgColor = bgColorHex
				}
			} else {
				return errors.ErrInvalidArgument.WrapErr(fmt.Errorf("invalid background color: %+v", bgColorIntfc))
			}
		} else {
			return errors.ErrInvalidArgument.WrapErr(fmt.Errorf("background color is not set"))
		}

		img = codec.ApplyBackgroundColor(img, bgColor)
	}

	// Encode
	return encoder.Encode(writer, img)
}
