package utils

import (
	"image"
)

func ExtractText(img image.Image) string {
	var (
		bounds    = img.Bounds()
		textBytes []byte
		textLen   uint32 = 0
		bitIndex         = 0
		curByte   byte   = 0
	)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, _, _, _ := img.At(x, y).RGBA()
			bit := byte(r & 1)

			if bitIndex < 32 {
				textLen = (textLen << 1) | uint32(bit)
			} else {
				curByte = (curByte << 1) | bit
				if (bitIndex-32+1)%8 == 0 {
					textBytes = append(textBytes, curByte)
					curByte = 0
					if uint32(len(textBytes)) >= textLen {
						showProgressBar("Extracting text...")
						return string(textBytes)
					}
				}
			}
			bitIndex++
		}
	}

	return ""
}
