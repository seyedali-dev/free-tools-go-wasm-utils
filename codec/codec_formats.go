package codec

// SupportedCodecFormat represents a supported image format.
type SupportedCodecFormat string

// Image formats.
const (
	PNG  SupportedCodecFormat = "png"
	GIF  SupportedCodecFormat = "gif"
	BMP  SupportedCodecFormat = "bmp"
	ICO  SupportedCodecFormat = "ico"
	AVIF SupportedCodecFormat = "avif"
	WEBP SupportedCodecFormat = "webp"

	// TIFF Family
	TIFF SupportedCodecFormat = "tiff"
	TIF  SupportedCodecFormat = "tif"

	// SVG Family
	SVG  SupportedCodecFormat = "svg"
	SVGZ SupportedCodecFormat = "svgz"

	// HEIF Family
	HEIF SupportedCodecFormat = "heif"
	HEIC SupportedCodecFormat = "heic"

	// JPEG Family
	JPEG   SupportedCodecFormat = "jpeg"
	JPG    SupportedCodecFormat = "jpg"
	JFIF   SupportedCodecFormat = "jfif"
	JP2    SupportedCodecFormat = "jp2"
	JPEGXR SupportedCodecFormat = "jxr"
	JPE    SupportedCodecFormat = "jpe"
	PJP    SupportedCodecFormat = "pjp"
	PJPEG  SupportedCodecFormat = "pjpeg"
)

// Audio Formats
const (
	// MP3 Family
	MP3  SupportedCodecFormat = "mp3"
	M4V  SupportedCodecFormat = "m4v"
	M4A  SupportedCodecFormat = "m4a"
	WAV  SupportedCodecFormat = "wav"
	FLAC SupportedCodecFormat = "flac"
	AIFF SupportedCodecFormat = "aiff"
	MIDI SupportedCodecFormat = "mid"
)

// Video formats.
const (
	MP4  SupportedCodecFormat = "mp4"
	MPG  SupportedCodecFormat = "mpg"
	MPEG SupportedCodecFormat = "mpeg"
	MOV  SupportedCodecFormat = "mov"
	AVI  SupportedCodecFormat = "avi"
	WMV  SupportedCodecFormat = "wmv"

	// WEBM Family
	WEBM SupportedCodecFormat = "webm"
	WEBA SupportedCodecFormat = "weba"

	// OGG Family
	OGG  SupportedCodecFormat = "ogg"
	OGA  SupportedCodecFormat = "oga"
	OPUS SupportedCodecFormat = "opus"
	OGV  SupportedCodecFormat = "ogv"
	OGM  SupportedCodecFormat = "ogm"
)

// Document formats.
const (
	PDF  SupportedCodecFormat = "pdf"
	DOC  SupportedCodecFormat = "doc"
	DOCX SupportedCodecFormat = "docx"
	XLS  SupportedCodecFormat = "xls"
	XLSX SupportedCodecFormat = "xlsx"
	CSV  SupportedCodecFormat = "csv"
)
