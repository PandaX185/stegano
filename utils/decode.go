package utils

import (
	"bytes"
	"image"
	"image/png"
	"log"

	cliColor "github.com/fatih/color"
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

func ExtractImage(img image.Image) image.Image {
	bounds := img.Bounds()
	bitIndex := 0
	var width, height uint32
	var hiddenBytes []byte
	var curByte byte
	red := cliColor.New(cliColor.FgRed).SprintFunc()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			for _, color := range []uint32{r, g, b} {
				bit := byte(color & 1)

				if bitIndex < 32 {
					width = (width << 1) | uint32(bit)
				} else if bitIndex < 64 {
					height = (height << 1) | uint32(bit)
				} else {
					curByte = (curByte << 1) | bit
					if (bitIndex-64)%8 == 7 { // After 8 bits, append the byte
						hiddenBytes = append(hiddenBytes, curByte)
						curByte = 0
					}
				}
				bitIndex++
			}
		}
	}
	if (bitIndex-63)%8 == 0 {
		hiddenBytes = append(hiddenBytes, curByte)
		curByte = 0
	}

	showProgressBar("Extracting image...")

	img, err := png.Decode(bytes.NewReader(hiddenBytes))
	if err != nil {
		log.Fatal(red("❌ Error extracting hidden image: "), err)
	}

	return img
}
