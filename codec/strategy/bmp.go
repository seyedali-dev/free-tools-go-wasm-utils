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

// Decode decodes a BMP image from the given reader. It uses the bmp.Decode function.
func (bmpCodec *BMPCodec) Decode(reader io.Reader, _ map[string]interface{}) (image.Image, error) {
	return bmp.Decode(reader)
}
