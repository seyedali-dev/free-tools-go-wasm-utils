package codec

import (
	"image"
	"io"
)

// ImageCodec is a common interface for encoding/decoding images.
type ImageCodec interface {
	Encode(writer io.Writer, img image.Image, options map[string]interface{}) error // Encode encodes an image to the given writer.
	Decode(reader io.Reader, options map[string]interface{}) (image.Image, error)   // Decode decodes an image from the given reader.
}
