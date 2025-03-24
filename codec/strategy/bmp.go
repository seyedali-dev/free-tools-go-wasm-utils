package codecstrategy

import (
	"image"
	"io"

	"golang.org/x/image/bmp"
)

// BMPCodec implements codec.ImageCodec for BMP format.
type BMPCodec struct{}

// Encode encodes a BMP image to the given writer. It takes no options.
func (bmpCodec *BMPCodec) Encode(writer io.Writer, img image.Image, _ map[string]interface{}) error {
	return bmp.Encode(writer, img)
}
