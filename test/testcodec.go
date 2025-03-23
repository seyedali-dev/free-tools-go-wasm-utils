package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/seyedali-dev/free-tools-go-wasm-utils/codec"
	codecfactory "github.com/seyedali-dev/free-tools-go-wasm-utils/codec/factory"
)

func main() {
	/*// Create a dummy image (a red rectangle)
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
	)*/

	// Read the image into image.Image
	inputImg := fmt.Sprintf("./test/input.%s", "png")
	file, err := os.Open(inputImg)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

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

	testDecode(err, img)
}

func testEncode(err error, jpgCodec codec.ImageCodec, img image.Image) {
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
	output := "./test/output.png"
	//output := "output.avif"
	err = os.WriteFile(output, buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Image saved as " + output)
}

func testDecode(err error, img image.Image) {
	byt, err := convertImageToPNGBytes(img)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(byt)
	decode, _, err := codec.Decode(buffer)
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
