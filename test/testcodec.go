package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"os"

	"github.com/seyedali-dev/free-tools-go-wasm-utils/codec"
	codecfactory "github.com/seyedali-dev/free-tools-go-wasm-utils/codec/factory"
)

func main() {
	// Create a dummy image (a red rectangle)
	rect := image.Rect(0, 0, 100, 100)
	img := image.NewRGBA(rect)
	draw.Draw(
		img,
		rect,
		&image.Uniform{C: color.RGBA{
			R: 255,
			A: 255},
		},
		image.Point{},
		draw.Src,
	)

	// Get a JPEG codec with quality option set to 90.
	//jpgCodec, err := codecfactory.GetEncoderFactory(codec.JPEG)
	//jpgCodec, err := codecfactory.GetEncoderFactory(codec.PNG)
	//jpgCodec, err := codecfactory.GetEncoderFactory(codec.ICO)
	//jpgCodec, err := codecfactory.GetEncoderFactory(codec.TIFF)
	//jpgCodec, err := codecfactory.GetEncoderFactory(codec.BMP)
	//jpgCodec, err := codecfactory.GetEncoderFactory(codec.WEBP)
	jpgCodec, err := codecfactory.GetEncoderFactory(codec.GIF)
	//jpgCodec, err := codecfactory.GetEncoderFactory(codec.AVIF)
	if err != nil {
		panic(err)
	}

	// Encode image to buffer.
	buf := &bytes.Buffer{}
	err = jpgCodec.Encode(buf, img, map[string]interface{}{"quality": 90})
	if err != nil {
		panic(err)
	}

	// Write the encoded image to a file.
	//output := "output.jpg"
	//output := "output.png"
	//output := "output.ico"
	//output := "output.tiff"
	//output := "output.bmp"
	//output := "output.webp"
	output := "output.gif"
	//output := "output.avif"
	err = os.WriteFile(output, buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Image saved as " + output)
}
