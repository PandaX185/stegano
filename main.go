package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/PandaX185/stegano/utils"
	"github.com/fatih/color"
)

func main() {
	// figure.NewColorFigure("Stegano CLI", "doom", "green", true).Print()
	// fmt.Println()

	embedCmd := flag.NewFlagSet("embed", flag.ExitOnError)
	extractCmd := flag.NewFlagSet("ext", flag.ExitOnError)

	inputFile := embedCmd.String("i", "", "Input image file")
	outputFile := embedCmd.String("o", "output.png", "Output image file")
	secretText := embedCmd.String("text", "", "Text to hide")

	extractFile := extractCmd.String("i", "", "Image to extract text from")

	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()

	fmt.Println()
	fmt.Println()
	fmt.Println(green("███████╗████████╗███████╗ ██████╗  █████╗ ███╗   ██╗ ██████╗  "))
	fmt.Println(green("██╔════╝╚══██╔══╝██╔════╝██╔════╝ ██╔══██╗████╗  ██║██╔═══██╗ "))
	fmt.Println(green("███████╗   ██║   █████╗  ██║  ███╗███████║██╔██╗ ██║██║   ██║ "))
	fmt.Println(green("╚════██║   ██║   ██╔══╝  ██║   ██║██╔══██║██║╚██╗██║██║   ██║ "))
	fmt.Println(green("███████║   ██║   ███████╗╚██████╔╝██║  ██║██║ ╚████║╚██████╔╝ "))
	fmt.Println(green("╚══════╝   ╚═╝   ╚══════╝ ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝  "))
	fmt.Println()

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  stegano embed -i image.png -o output.png -text 'hidden message'")
		fmt.Println("  stegano ext -i output.png")
		os.Exit(0)
	}

	switch os.Args[1] {
	case "embed":
		embedCmd.Parse(os.Args[2:])
		if *inputFile == "" || *secretText == "" {
			log.Fatal(red("❌ Missing required arguments for embed: -i and -text"))
		}

		img, err := utils.LoadImage(*inputFile)
		if err != nil {
			log.Fatal(red("❌ Error loading image: "), err)
		}

		hiddenImg := utils.EmbedText(img, *secretText)
		err = utils.SaveImage(*outputFile, hiddenImg)
		if err != nil {
			log.Fatal(red("❌ Error saving image: "), err)
		}
		fmt.Println(green("\n\n✅ Text successfully embedded into"), bold(*outputFile))

	case "ext":
		extractCmd.Parse(os.Args[2:])
		if *extractFile == "" {
			log.Fatal(red("❌ Missing required argument: -i"))
		}

		img, err := utils.LoadImage(*extractFile)
		if err != nil {
			log.Fatal(red("❌ Error loading image: "), err)
		}

		secret := utils.ExtractText(img)
		if secret == "" {
			log.Fatal(red("❌ No text found in image"))
		}
		fmt.Println(green("\n\n🔎 Extracted Text:"), bold(secret))
	default:
		fmt.Println("Unknown command. Use 'embed' or 'extract'.")
	}

	fmt.Println()
}
