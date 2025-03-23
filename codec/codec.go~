package codec

import (
	"image"
	"io"
)

// ImageCodec is a common interface for encoding/decoding images.
type ImageCodec interface {
	Encode(writer io.Writer, img image.Image, options map[string]interface{}) error // Encode encodes an image to the given writer.
}

// Decode decodes an image with default image.Decode decoder from the given reader. It returns the decoded image, format and error if occurred.
func Decode(reader io.Reader, options map[string]interface{}) (image.Image, string, error) {
	return image.Decode(reader)
}
