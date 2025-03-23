package codecstrategy

import (
	"image"
	"io"

	"github.com/seyedali-dev/free-tools-go-wasm-utils/codec/types"
	"github.com/vldrus/golang/image/ico"
	"golang.org/x/image/draw"
)

var (
	MinimumICOSize = 16
	MaximumICOSize = 256
	DefaultICOSize = 256
	ValidICOSizes  = []int{16, 24, 32, 48, 64, 128, 256}
)

// ICOCodec implements codec.ImageCodec for ICO format.
type ICOCodec struct{}

// Encode encodes an ICO image to the given writer. It takes types.Dimension as an option to specify the size of the ICO image.
func (icoCodec *ICOCodec) Encode(writer io.Writer, img image.Image, options map[string]interface{}) error {
	dimension := DefaultICOSize
	dimensionStr := options[types.Dimension]
	if dimensionStr != "" {
		if d, ok := dimensionStr.(int); ok {
			dimension = d
		}
	}

	resizedImg := ResizeImage(img, dimension)

	return ico.Encode(writer, resizedImg)
}

// ResizeImage resizes the given image to the specified size.
func ResizeImage(img image.Image, size int) image.Image {
	srcBounds := img.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, size, size))
	draw.CatmullRom.Scale(dst, dst.Bounds(), img, srcBounds, draw.Over, nil)
	return dst
}

// Decode decodes an ICO image from the given reader. It uses the default image.Decode function.
func (icoCodec *ICOCodec) Decode(reader io.Reader, _ map[string]interface{}) (image.Image, error) {
	img, _, err := image.Decode(reader)

	return img, err
}
