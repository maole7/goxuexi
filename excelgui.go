package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello, World!")
	questionButton := widget.NewButton("Ask a Question", func() {
		fmt.Println("You clicked the button!")
	})
  content := container.NewVBox(
		questionButton,
	)
	w.SetContent(content)
	w.ShowAndRun()
}

