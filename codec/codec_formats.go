package codec

// SupportedCodecFormat represents a supported image format.
type SupportedCodecFormat string

const (
	PNG SupportedCodecFormat = "png"

	// JPEG Family
	JPEG   SupportedCodecFormat = "jpeg"
	JPG    SupportedCodecFormat = "jpg"
	JFIF   SupportedCodecFormat = "jfif"
	JP2    SupportedCodecFormat = "jp2"
	JPEGXR SupportedCodecFormat = "jxr"
	JPE    SupportedCodecFormat = "jpe"

	GIF  SupportedCodecFormat = "gif"
	BMP  SupportedCodecFormat = "bmp"
	TIFF SupportedCodecFormat = "tiff"
	ICO  SupportedCodecFormat = "ico"
	AVIF SupportedCodecFormat = "avif"
	WEBP SupportedCodecFormat = "webp"
)

// SupportedFormats is a list of supported image formats.
var SupportedFormats = []SupportedCodecFormat{PNG, JPEG, GIF, BMP, TIFF, ICO, AVIF, WEBP}
