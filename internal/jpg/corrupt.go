package jpg

import (
	"bytes"
	"fmt"
	"image"
	"math/rand"
	"time"

	// "image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"strings"

	"github.com/AvraamMavridis/randomcolor"
)

func Corrupt(file *File) (*File, error) {
	// make a place for the new image to be stored
	newJPG := File{
		Path: strings.Split(file.Path, ".jpg")[0] + "-corrupted.jpg",
	}

	width := file.Image.Bounds().Size().X
	height := file.Image.Bounds().Size().Y

	img := image.NewRGBA(file.Image.Bounds())
	draw.Draw(img, img.Bounds(), file.Image, image.Point{}, draw.Src)

	// create a rectangle
	rectCount := randInt(2, 25)
	for i := 0; i < rectCount; i++ {
		newRandomRect(img, height-i, width-i)
	}

	// encode the bytes into a jpg
	buff := bytes.NewBuffer([]byte{})

	jpeg.Encode(buff, img, &jpeg.Options{Quality: 100})

	// decode the bytes into a new image
	newImage, err := jpeg.Decode(bytes.NewBuffer(buff.Bytes()))
	if err != nil {
		return nil, err
	}
	newJPG.Image = newImage

	return &newJPG, nil
}

func newRandomRect(img *image.RGBA, height, width int) {
	top := randInt(0, height)
	right := randInt(0, width)
	bottom := randInt(0, height)
	left := randInt(0, width)
	bgColor := newRandomColorRGBA()
	bgColor.A = 1
	overlay := image.NewRGBA(image.Rect(left, top, right, bottom))
	draw.Draw(img, overlay.Bounds(), &image.Uniform{C: bgColor}, image.Point{}, draw.Over)
}

func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(max-min+1) + min
	return i
}

func newRandomColorRGBA() color.RGBA {
	rgba, _ := hexToRGBA(newRandomColorHex())
	return rgba
}

func newRandomColorHex() string {
	hex := randomcolor.GetRandomColorInHex()
	return hex
}

func hexToRGBA(hex string) (color.RGBA, error) {
	var (
		rgba             color.RGBA
		err              error
		errInvalidFormat = fmt.Errorf("invalid")
	)
	rgba.A = 0xff
	if hex[0] != '#' {
		return rgba, errInvalidFormat
	}
	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}
	switch len(hex) {
	case 7:
		rgba.R = hexToByte(hex[1])<<4 + hexToByte(hex[2])
		rgba.G = hexToByte(hex[3])<<4 + hexToByte(hex[4])
		rgba.B = hexToByte(hex[5])<<4 + hexToByte(hex[6])
	case 4:
		rgba.R = hexToByte(hex[1]) * 17
		rgba.G = hexToByte(hex[2]) * 17
		rgba.B = hexToByte(hex[3]) * 17
	default:
		err = errInvalidFormat
	}
	return rgba, err
}
