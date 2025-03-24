package codecfactory

import (
	"github.com/seyedali-dev/free-tools-go-wasm-utils/codec"
	codecstrategy "github.com/seyedali-dev/free-tools-go-wasm-utils/codec/strategy"
	"github.com/seyedali-dev/free-tools-go-wasm-utils/errors"
)

// GetEncoderFactory returns an encoder factory for the given format.
//
// Supported formats:
//  1. codec.PNG: PNG format with default compression (png.DefaultCompression)
//  2. codec.JPEG/codec.JPG: JPEG format with default quality (90%)
//  3. codec.WEBP: WEBP format with default extended format (false)
//  4. codec.AVIF: AVIF format not implemented yet
//  5. codec.TIFF: TIFF format with default compression (uncompressed)
//  6. codec.BMP: BMP format
//  7. codec.GIF: GIF format
//  8. codec.ICO: ICO format with default dimension (codecstrategy.DefaultICOSize)
func GetEncoderFactory(format codec.SupportedCodecFormat) (codec.ImageCodec, error) {
	normalizeFormat := codec.NormalizeJPEGFormat(format)

	switch normalizeFormat {
	case codec.PNG:
		return &codecstrategy.PNGCodec{}, nil
	case codec.JPEG:
		return &codecstrategy.JPEGCodec{}, nil
	case codec.WEBP:
		return &codecstrategy.WEBPCodec{}, nil
	case codec.AVIF:
		return &codecstrategy.AVIFCodec{}, nil
	case codec.TIFF:
		return &codecstrategy.TIFFCodec{}, nil
	case codec.BMP:
		return &codecstrategy.BMPCodec{}, nil
	case codec.GIF:
		return &codecstrategy.GIFCodec{}, nil
	case codec.ICO:
		return &codecstrategy.ICOCodec{}, nil
	default:
		return nil, errors.ErrUnsupportedFormat
	}
}
