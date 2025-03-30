package codecstrategy

import (
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"io"

	"github.com/seyedali-dev/free-tools-go-wasm-utils/codec/types"
	"github.com/seyedali-dev/free-tools-go-wasm-utils/errors"
)

// GIFCodec implements codec.ImageCodec for GIF format.
type GIFCodec struct{}

// Encode encodes an GIF image to the given writer.
// It takes types.NumColors, types.Quantizer, and types.Drawer as options.
//
// It either sets the options or none. There is no default.
func (gifCodec *GIFCodec) Encode(writer io.Writer, img image.Image, options map[string]interface{}) error {
	var gifOpts *gif.Options

	// Set options
	if options[types.NumColors] != nil {
		if numbColor, ok := options[types.NumColors]; ok {
			if numbColor, ok := numbColor.(int); ok {
				gifOpts.NumColors = numbColor
			} else {
				return errors.ErrInvalidArgument.WrapErr(fmt.Errorf("invalid num colors: %+v", numbColor))
			}
		} else {
			return errors.ErrInvalidArgument.WrapErr(fmt.Errorf("num colors is not set"))
		}
	}
	if options[types.Quantizer] != nil {
		if quantizer, ok := options[types.Quantizer]; ok {
			if quantizer, ok := quantizer.(draw.Quantizer); ok {
				gifOpts.Quantizer = quantizer.(draw.Quantizer)
			} else {
				return errors.ErrInvalidArgument.WrapErr(fmt.Errorf("invalid quantizer: %+v", quantizer))
			}
		} else {
			return errors.ErrInvalidArgument.WrapErr(fmt.Errorf("quantizer is not set"))
		}
	}
	if options[types.Drawer] != nil {
		if drawer, ok := options[types.Drawer]; ok {
			if drawer, ok := drawer.(draw.Drawer); ok {
				gifOpts.Drawer = drawer
			} else {
				return errors.ErrInvalidArgument.WrapErr(fmt.Errorf("invalid drawer: %+v", drawer))
			}
		} else {
			return errors.ErrInvalidArgument.WrapErr(fmt.Errorf("drawer is not set"))
		}
	}

	// Encode
	return gif.Encode(writer, img, gifOpts)
}
