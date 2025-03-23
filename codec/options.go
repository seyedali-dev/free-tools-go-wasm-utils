package codec

type Option func(*Config)

type Config struct {
	Quality  int  // For JPEG/WebP compression
	Lossless bool // For WebP/AVIF
	BitDepth int  // For PNG/BMP
}

func WithQuality(quality int) Option {
	return func(config *Config) { config.Quality = quality }
}

func WithLossless(lossless bool) Option {
	return func(config *Config) { config.Lossless = lossless }
}
