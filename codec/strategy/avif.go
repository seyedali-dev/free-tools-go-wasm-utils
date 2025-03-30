package codecstrategy

import (
	"fmt"
	"image"
	"io"
)

// AVIFCodec implements codec.ImageCodec for AVIF format.
type AVIFCodec struct{}

// Encode encodes an AVIF image to the given writer. It is not implemented yet.
func (avifCodec *AVIFCodec) Encode(writer io.Writer, img image.Image, options map[string]interface{}) error {
	// TODO: integrate with an AVIF encoder library.
	return fmt.Errorf("AVIF encoding not implemented")
}
