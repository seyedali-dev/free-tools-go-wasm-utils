package codecstrategy

import (
	"image"
	"image/draw"
	"image/gif"
	"io"

	"github.com/seyedali-dev/free-tools-go-wasm-utils/codec/types"
)

// GIFCodec implements codec.ImageCodec for GIF format.
type GIFCodec struct{}

// Encode encodes an GIF image to the given writer. It takes types.NumColors, types.Quantizer, and types.Drawer as options.
func (gifCodec *GIFCodec) Encode(writer io.Writer, img image.Image, options map[string]interface{}) error {
	var gifOpts *gif.Options

	numColor := options[types.NumColors]
	if numColor != nil {
		gifOpts.NumColors = numColor.(int)
	}

	quantizer := options[types.Quantizer]
	if quantizer != nil {
		gifOpts.Quantizer = quantizer.(draw.Quantizer)
	}

	drawer := options[types.Drawer]
	if drawer != nil {
		gifOpts.Drawer = drawer.(draw.Drawer)
	}

	return gif.Encode(writer, img, gifOpts)
}
