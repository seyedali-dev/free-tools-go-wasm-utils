package codecstrategy

import (
	"image"
	"io"

	"github.com/HugoSmits86/nativewebp"
	"github.com/seyedali-dev/free-tools-go-wasm-utils/codec/types"
)

// WEBPCodec implements codec.ImageCodec for WEBP format.
//
// Ref:
//   - The reddit post :: https://www.reddit.com/r/golang/comments/1hnc57a/native_webp_encoder_for_go/
//   - The actual repo :: https://github.com/HugoSmits86/nativewebp,
type WEBPCodec struct{}

// Encode encodes a WEBP image to the given writer. It takes types.UseExtendedFormat as an option.
func (webpCodec *WEBPCodec) Encode(writer io.Writer, img image.Image, options map[string]interface{}) error {
	var webpOpts *nativewebp.Options

	if uef, ok := options[types.UseExtendedFormat]; ok {
		if uefBool, ok := uef.(bool); ok {
			webpOpts.UseExtendedFormat = uefBool
		}
	}

	return nativewebp.Encode(writer, img, webpOpts)
}

// Decode decodes a WEBP image from the given reader. It uses the nativewebp.Decode function.
func (webpCodec *WEBPCodec) Decode(reader io.Reader, _ map[string]interface{}) (image.Image, error) {
	return nativewebp.Decode(reader)
}
