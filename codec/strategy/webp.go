package codecstrategy

import (
	"fmt"
	"image"
	"io"

	"github.com/HugoSmits86/nativewebp"
	"github.com/seyedali-dev/free-tools-go-wasm-utils/codec/types"
	"github.com/seyedali-dev/free-tools-go-wasm-utils/errors"
)

// WEBPCodec implements codec.ImageCodec for WEBP format.
// It takes types.UseExtendedFormat as an option.
//
// It either sets the options or none. There is no default.
//
// Ref:
//   - The reddit post :: https://www.reddit.com/r/golang/comments/1hnc57a/native_webp_encoder_for_go/
//   - The actual repo :: https://github.com/HugoSmits86/nativewebp,
type WEBPCodec struct{}

// Encode encodes a WEBP image to the given writer. It takes types.UseExtendedFormat as an option.
func (webpCodec *WEBPCodec) Encode(writer io.Writer, img image.Image, options map[string]interface{}) error {
	var webpOpts *nativewebp.Options

	// Set options
	if options[types.UseExtendedFormat] != nil {
		if uef, ok := options[types.UseExtendedFormat]; ok {
			if uefBool, ok := uef.(bool); ok {
				webpOpts.UseExtendedFormat = uefBool
			} else {
				return errors.ErrInvalidArgument.WrapErr(fmt.Errorf("invalid use extended format: %+v", uef))
			}
		} else {
			return errors.ErrInvalidArgument.WrapErr(fmt.Errorf("use extended format is not set"))
		}
	}

	// Encode
	return nativewebp.Encode(writer, img, webpOpts)
}
