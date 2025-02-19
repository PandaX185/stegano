package utils

import (
	"image"
	"image/color"
	"log"

	cliColor "github.com/fatih/color"
	"github.com/nfnt/resize"
)

func EmbedText(img image.Image, text string) image.Image {
	bounds := img.Bounds()
	res := image.NewRGBA(bounds)

	textBytes := []byte(text)
	textLen := len(textBytes)
	bitIndex := 0

	showProgressBar("Embedding text...")

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

	return res
}

func EmbedImage(img image.Image, hidden image.Image) image.Image {
	red := cliColor.New(cliColor.FgRed).SprintFunc()
	bounds := img.Bounds()
	res := image.NewRGBA(bounds)
	bitIndex := 0

	hostW, hostH := bounds.Dx(), bounds.Dy()
	resizedImage := resize.Resize(uint(hostW)/4, uint(hostH)/4, hidden, resize.Lanczos3)

	hiddenBytes, w, h, err := imageToBytes(resizedImage)
	if err != nil {
		log.Fatal(red("❌ Error converting hidden image to bytes: "), err)
	}

	requiredPixels := (len(hiddenBytes) * 8) / 3
	if bounds.Dx()*bounds.Dy() < requiredPixels {
		log.Fatal(red("❌ Error: Host image is too small to embed the hidden image."))
	}

	showProgressBar("Embedding image...")

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)

			for i := 0; i < 3; i++ {
				var bit uint8

				if bitIndex < 32 {
					bit = uint8((w >> (31 - bitIndex)) & 1)
				} else if bitIndex < 64 {
					bit = uint8((h >> (63 - bitIndex)) & 1)
				} else {
					byteIndex := (bitIndex - 64) / 8
					if byteIndex < len(hiddenBytes) {
						shift := 7 - ((bitIndex - 64) % 8)
						bit = (hiddenBytes[byteIndex] >> shift) & 1
					} else {
						break
					}
				}

				if i == 0 {
					r8 = (r8 & 0xFE) | bit
				} else if i == 1 {
					g8 = (g8 & 0xFE) | bit
				} else if i == 2 {
					b8 = (b8 & 0xFE) | bit
				}
				bitIndex++
			}

			res.Set(x, y, color.RGBA{r8, g8, b8, uint8(a >> 8)})
		}
	}

	return res
}
