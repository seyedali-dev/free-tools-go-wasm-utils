package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
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
	//imgCodecTest, err := codecfactory.GetEncoderFactory(codec.JPEG)
	imgCodecTest, err := codecfactory.GetEncoderFactory(codec.PNG)
	//imgCodecTest, err := codecfactory.GetEncoderFactory(codec.ICO)
	//imgCodecTest, err := codecfactory.GetEncoderFactory(codec.TIFF)
	//imgCodecTest, err := codecfactory.GetEncoderFactory(codec.BMP)
	//imgCodecTest, err := codecfactory.GetEncoderFactory(codec.WEBP)
	//imgCodecTest, err := codecfactory.GetEncoderFactory(codec.GIF)
	//imgCodecTest, err := codecfactory.GetEncoderFactory(codec.AVIF)
	if err != nil {
		panic(err)
	}

	testEncode(err, imgCodecTest, img)

	testDecode(err, img, imgCodecTest)
}

func testEncode(err error, jpgCodec codec.ImageCodec, img *image.RGBA) {
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

func testDecode(err error, img *image.RGBA, imgCodecTest codec.ImageCodec) {
	byt, err := convertImageToPNGBytes(img)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(byt)
	decode, err := imgCodecTest.Decode(buffer, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decode: ", decode.Bounds())
}

// convertImageToPNGBytes converts an image.Image to a byte slice in PNG format.
func convertImageToPNGBytes(img image.Image) ([]byte, error) {
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
