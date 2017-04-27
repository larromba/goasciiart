// by mo2zie

package goasciiart

import (
    "github.com/nfnt/resize"

    "os"
    "bytes"
    "image"
    "reflect"
    "image/color"
    _ "image/png"
    _ "image/jpeg"
)

var ASCIISTR = "MND8OZ$7I?+=~:,.."

func LoadImage(path string) (image.Image, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    img, _, err := image.Decode(f)
    return img, err
}

func ScaleImage(img image.Image, w int) (image.Image, int, int) {
    sz := img.Bounds()
    h := (sz.Max.Y * w * 10) / (sz.Max.X * 16)
    img = resize.Resize(uint(w), uint(h), img, resize.Lanczos3)
    return img, w, h
}

func Convert2Ascii(img image.Image, w, h int) []byte {
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
