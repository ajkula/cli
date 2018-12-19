package main

import (
	"fmt"
	"net/http"

	"github.com/nfnt/resize"

	"bytes"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"reflect"
)

var ASCIISTR = "MND8OZ$7I?+=~:,.."

func byPassErrors(e error) error {
	if e != nil {
		fmt.Printf("[Cannot recognize image format]\n")
		return e
	}
	return nil
}

func fromUrlAndSize(url string, width int) (image.Image, int, error) {
	res, err := http.Get(url)
	if byPassErrors(err) != nil {
		return nil, 0, err
	}

	defer res.Body.Close()
	img, _, err := image.Decode(res.Body)
	if byPassErrors(err) != nil {
		return nil, 0, err
	}

	return img, width, nil
}

func ScaleImage(img image.Image, w int) (image.Image, int, int) {
	sz := img.Bounds()
	h := (sz.Max.Y * w * 10) / (sz.Max.X * 16)
	img = resize.Resize(uint(w), uint(h), img, resize.Lanczos3)
	return img, w, h
}

// Convert2Ascii(ScaleImage(fromUrlAndSize(res.UrlToImage, 80)))
func Convert2Ascii(url string, width int) []byte {
	if img, width, err := fromUrlAndSize(url, width); err == nil {
		return ConvertImg2Ascii(ScaleImage(img, width))
	} else {
		return []byte{}
	}
}

func ConvertImg2Ascii(img image.Image, w, h int) []byte {
	table := []byte(ASCIISTR)
	buf := new(bytes.Buffer)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			g := color.GrayModel.Convert(img.At(j, i))
			y := reflect.ValueOf(g).FieldByName("Y").Uint()
			pos := int(y * 16 / 255)
			_ = buf.WriteByte(table[pos])
		}
		_ = buf.WriteByte('\n')
	}
	return buf.Bytes()
}

// func main() {
//     p := Convert2Ascii(ScaleImage(Init()))
//     fmt.Print(string(p))
// }
