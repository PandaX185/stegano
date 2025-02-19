package utils

import (
	"time"

	cliColor "github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
)

func showProgressBar(text string) {
	blue := cliColor.New(cliColor.FgBlue).SprintFunc()
	bar := progressbar.Default(-1, blue(text))
	for i := 0; i < 30; i++ {
		time.Sleep(10 * time.Millisecond)
		bar.Add(1)
	}
}
