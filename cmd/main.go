package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/kglaus/geodienste-cli/pkg/gui"
	"github.com/kglaus/geodienste-cli/pkg/stac"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("geodienst-cli")
	myWindow.SetContent(widget.NewLabel("geodienste-cli2"))
	myWindow.Resize(fyne.NewSize(300, 600))

	canvasObjects := []fyne.CanvasObject{}
	//green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	//text1 := canvas.NewText("Part1", green)
	//text2 := canvas.NewText("Part2", green)
	//text3 := canvas.NewText("Part3", color.White)
	//text4 := canvas.NewText("Part4", color.White)
	//
	//canvasObjects = append(canvasObjects, text1, text2, text3, text4)

	contentBottom := container.New(layout.NewGridLayout(4), canvasObjects...)

	buttonGenerate := widget.NewButton("click me", func() {
		collections := stac.GetCollections()
		collectionObjects := []fyne.CanvasObject{}

		for _, element := range collections.Collections {
			collectionObjects = append(collectionObjects, gui.CollectionButton(element, contentBottom))
		}

		contentBottom.Objects = collectionObjects
		contentBottom.Refresh()
	})

	contentTop := container.New(layout.NewGridLayout(1), buttonGenerate)

	content := container.New(layout.NewGridLayout(1), contentTop, contentBottom)

	myWindow.SetContent(content)

	myWindow.Show()
	myApp.Run()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}
