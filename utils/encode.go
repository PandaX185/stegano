package utils

import (
	"image"
	"image/color"
)

func EmbedText(img image.Image, text string) image.Image {
	bounds := img.Bounds()
	res := image.NewRGBA(bounds)

	textBytes := []byte(text)
	textLen := len(textBytes)
	bitIndex := 0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)

			if bitIndex < 32 {
				bit := uint8((textLen >> (31 - bitIndex)) & 1)
				r8 = (r8 & 0xfe) | bit
			} else {
				byteIndex := (bitIndex - 32) / 8
				if byteIndex < len(textBytes) {
					bit := (textBytes[byteIndex] >> (7 - ((bitIndex - 32) % 8))) & 1
					r8 = (r8 & 0xfe) | bit
				}
			}
			bitIndex++
			res.Set(x, y, color.RGBA{r8, g8, b8, uint8(a >> 8)})
		}
	}

	showProgressBar("Embedding text...")

	return res
}
