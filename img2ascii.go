package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"

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

func Convert2Ascii(url string, width int) []byte {
	if img, width, err := fromUrlAndSize(url, width); err == nil {
		return ConvertImg2Ascii(ScaleImage(img, width))
	}
	return []byte{}
}

func ReadImgFile(uri string, width string) []byte {
	size := 80
	var err error

	if uri == "" {
		return []byte{}
	}
	if width != "" {
		if size, err = strconv.Atoi(width); err != nil {
			check(err)
		}
	}

	reader, err := os.Open(uri)
	check(err)
	defer reader.Close()
	// if img, _, err := image.Decode(reader); err != nil {
	img, _, err := image.Decode(reader)
	// return ConvertImg2Ascii(ScaleImage(img, size))
	// } else {
	check(err)
	return ConvertImg2Ascii(ScaleImage(img, size))
	// }
	// return []byte{}
}

func DisplayAsciiFromLocalFile(uri string, size string) {
	asciiArt := ReadImgFile(uri, size)
	fmt.Println(string(asciiArt))
}

// func ConvertImg2Ascii(img image.Image, w, h int) []byte {
// 	table := []byte(ASCIISTR)
// 	buf := new(bytes.Buffer)

// 	for i := 0; i < h; i++ {
// 		for j := 0; j < w; j++ {
// 			g := color.GrayModel.Convert(img.At(j, i))
// 			y := reflect.ValueOf(g).FieldByName("Y").Uint()
// 			pos := int(y * 16 / 255)
// 			_ = buf.WriteByte(table[pos])
// 		}
// 		_ = buf.WriteByte('\n')
// 	}
// 	return buf.Bytes()
// }

func ConvertImg2Ascii(img image.Image, w, h int) []byte {
	table := []byte(ASCIISTR)
	buf := new(bytes.Buffer)

	for i := 0; i < h; i++ {
		var wg sync.WaitGroup
		wg.Add(1)
		func() {
			for j := 0; j < w; j++ {
				g := color.GrayModel.Convert(img.At(j, i))
				y := reflect.ValueOf(g).FieldByName("Y").Uint()
				pos := int(y * 16 / 255)
				_ = buf.WriteByte(table[pos])
			}
			_ = buf.WriteByte('\n')
			wg.Done()
		}()
		wg.Wait()
	}
	return buf.Bytes()
}
