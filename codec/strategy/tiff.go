package codecstrategy

import (
	"image"
	"io"

	"github.com/seyedali-dev/free-tools-go-wasm-utils/codec/types"
	"golang.org/x/image/tiff"
)

// TIFFCodec implements codec.ImageCodec for TIFF format.
type TIFFCodec struct{}

// Encode encodes a TIFF image to the given writer. It takes types.CompressionType as an option to specify the compression type.
func (tiffCodec *TIFFCodec) Encode(writer io.Writer, img image.Image, options map[string]interface{}) error {
	compressionType := tiff.Uncompressed // Default compression type
	if compTypeStr, ok := options[types.CompressionType]; ok {
		if compType, ok := compTypeStr.(int); ok {
			compressionType = tiff.CompressionType(compType)
		}
	}

	return tiff.Encode(writer, img, &tiff.Options{
		Compression: compressionType,
		Predictor:   true,
	})
}
