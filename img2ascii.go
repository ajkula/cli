package main

import (
	"net/http"

	"github.com/nfnt/resize"

	"bytes"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"reflect"
)

var ASCIISTR = "MND8OZ$7I?+=~:,.."

func fromUrlAndSize(url string, width int) (image.Image, int) {
	res, err := http.Get(url)
	check(err)

	defer res.Body.Close()
	img, _, err := image.Decode(res.Body)
	check(err)

	return img, width
}

func ScaleImage(img image.Image, w int) (image.Image, int, int) {
	sz := img.Bounds()
	h := (sz.Max.Y * w * 10) / (sz.Max.X * 16)
	img = resize.Resize(uint(w), uint(h), img, resize.Lanczos3)
	return img, w, h
}

// Convert2Ascii(ScaleImage(fromUrlAndSize(res.UrlToImage, 80)))
func Convert2Ascii(url string, width int) []byte {
	return ConvertImg2Ascii(ScaleImage(fromUrlAndSize(url, width)))
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
