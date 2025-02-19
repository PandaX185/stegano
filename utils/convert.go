package utils

import (
	"bytes"
	"image"
	"image/png"
)

func imageToBytes(img image.Image) ([]byte, int, int, error) {
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		return nil, 0, 0, err
	}
	bounds := img.Bounds()
	return buf.Bytes(), bounds.Dx(), bounds.Dy(), nil
}
