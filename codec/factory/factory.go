package codecfactory

import (
	"github.com/seyedali-dev/free-tools-go-wasm-utils/codec"
	codecstrategy "github.com/seyedali-dev/free-tools-go-wasm-utils/codec/strategy"
	"github.com/seyedali-dev/free-tools-go-wasm-utils/errors"
)

// GetEncoderFactory returns an encoder factory for the given format.
//
// Supported formats:
//  1. "jpeg"/"jpg": JPEG format with default quality (90%)
//  2. "png": PNG format with default compression (png.DefaultCompression)
//  3. "webp": WEBP format with default extended format (false)
//  4. "avif": AVIF format not implemented yet
//  5. "tiff": TIFF format with default compression (uncompressed)
//  6. "bmp": BMP format
//  7. "gif": GIF format
//  8. "ico": ICO format with default dimension (codecstrategy.DefaultICOSize)
func GetEncoderFactory(format string) (codec.ImageCodec, error) {
	switch format {
	case "png":
		return &codecstrategy.PNGCodec{}, nil
	case "jpeg", "jpg":
		return &codecstrategy.JPEGCodec{}, nil
	case "webp":
		return &codecstrategy.WEBPCodec{}, nil
	case "avif":
		return &codecstrategy.AVIFCodec{}, nil
	case "tiff":
		return &codecstrategy.TIFFCodec{}, nil
	case "bmp":
		return &codecstrategy.BMPCodec{}, nil
	case "gif":
		return &codecstrategy.GIFCodec{}, nil
	case "ico":
		return &codecstrategy.ICOCodec{}, nil
	default:
		return nil, errors.ErrUnsupportedFormat
	}
}
