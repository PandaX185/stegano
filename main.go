package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PandaX185/stegano/utils"
	"github.com/fatih/color"
)

func main() {

	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()

	if len(os.Args) < 2 {
		fmt.Println()
		fmt.Println()
		fmt.Println(green("███████╗████████╗███████╗ ██████╗  █████╗ ███╗   ██╗ ██████╗  "))
		fmt.Println(green("██╔════╝╚══██╔══╝██╔════╝██╔════╝ ██╔══██╗████╗  ██║██╔═══██╗ "))
		fmt.Println(green("███████╗   ██║   █████╗  ██║  ███╗███████║██╔██╗ ██║██║   ██║ "))
		fmt.Println(green("╚════██║   ██║   ██╔══╝  ██║   ██║██╔══██║██║╚██╗██║██║   ██║ "))
		fmt.Println(green("███████║   ██║   ███████╗╚██████╔╝██║  ██║██║ ╚████║╚██████╔╝ "))
		fmt.Println(green("╚══════╝   ╚═╝   ╚══════╝ ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝  "))
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  stegano embed 'hidden message' into image.png as output.png")
		fmt.Println("  stegano embed image hidden.png into image.png as output.png")
		fmt.Println("  stegano ext from image.png")
		fmt.Println("  stegano ext image from image.png")
		fmt.Println("  stegano ext image from image.png as output.png")
		fmt.Println()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "embed":
		if len(os.Args) == 6 || len(os.Args) == 8 {
			hidden, err := utils.LoadImage(os.Args[3])
			if err != nil {
				log.Fatal(red("❌ Error loading image: "), err)
			}

			img, err := utils.LoadImage(os.Args[5])
			if err != nil {
				log.Fatal(red("❌ Error loading image: "), err)
			}

			hiddenImg := utils.EmbedImage(img, hidden)
			outputFile := "output.png"
			if len(os.Args) == 8 {
				outputFile = os.Args[7]
			}

			err = utils.SaveImage(outputFile, hiddenImg)
			if err != nil {
				log.Fatal(red("❌ Error saving image: "), err)
			}
			fmt.Println(green("\n\n✅ Image successfully embedded into"), bold(outputFile))

		} else {
			if len(os.Args) < 5 {
				log.Fatal(red("❌ No hidden message proivded"))
			}

			img, err := utils.LoadImage(os.Args[4])
			if err != nil {
				log.Fatal(red("❌ Error loading image: "), err)
			}

			hiddenImg := utils.EmbedText(img, os.Args[2])
			outputFile := "output.png"
			if len(os.Args) == 7 {
				outputFile = os.Args[6]
			}
			err = utils.SaveImage(outputFile, hiddenImg)
			if err != nil {
				log.Fatal(red("❌ Error saving image: "), err)
			}
			fmt.Println(green("\n\n✅ Text successfully embedded into"), bold(outputFile))
		}

	case "ext":
		if len(os.Args) >= 5 {
			img, err := utils.LoadImage(os.Args[4])
			if err != nil {
				log.Fatal(red("❌ Error loading image: "), err)
			}

			hidden := utils.ExtractImage(img)
			if hidden == nil {
				log.Fatal(red("❌ No image found in image"))
			}

			outputFile := "output.png"
			if len(os.Args) == 7 {
				outputFile = os.Args[6]
			}

			err = utils.SaveImage(outputFile, hidden)
			if err != nil {
				log.Fatal(red("❌ Error saving image: "), err)
			}

			fmt.Println(green("\n\n✅ Image successfully extracted into"), bold(outputFile))
		} else {
			if len(os.Args) < 4 {
				log.Fatal(red("❌ No image provided"))
			}

			img, err := utils.LoadImage(os.Args[3])
			if err != nil {
				log.Fatal(red("❌ Error loading image: "), err)
			}

			secret := utils.ExtractText(img)
			if secret == "" {
				log.Fatal(red("❌ No text found in image"))
			}
			fmt.Println(green("\n\n🔎 Extracted Text:"), bold(secret))
		}
	default:
		fmt.Println("Unknown command. Use 'embed' or 'extract'.")
	}

	fmt.Println()
}
