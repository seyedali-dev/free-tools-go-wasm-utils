package codecstrategy

import (
	"image"
	"image/jpeg"
	"io"

	"github.com/seyedali-dev/free-tools-go-wasm-utils/codec/types"
)

// JPEGCodec implements codec.ImageCodec for JPEG format.
type JPEGCodec struct{}

// Encode encodes a JPEG image to the given writer. It takes types.Quality as an option to specify the quality of the JPEG image.
func (jpegCodec *JPEGCodec) Encode(writer io.Writer, img image.Image, options map[string]interface{}) error {
	var quality int
	if qualityStr, ok := options[types.Quality]; ok {
		if qInt, ok := qualityStr.(int); ok {
			quality = qInt
		}
	}

	opts := &jpeg.Options{}
	if quality > 0 {
		opts.Quality = quality
	}

	//Encode
	return jpeg.Encode(writer, img, opts)
}
