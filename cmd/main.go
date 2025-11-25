package main

import (
	"fmt"
	"strings"

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
	myWindow.Resize(fyne.NewSize(1050, 400))

	var canvasObjects []fyne.CanvasObject

	contentBottom := container.New(layout.NewGridWrapLayout(fyne.Size{Width: 200, Height: 125}), canvasObjects...)
	contentBottomWrapper := container.NewVScroll(contentBottom)

	var collectionObjects []fyne.CanvasObject

	buttonGenerate := widget.NewButton("click me", func() {
		collections := stac.GetCollections()

		for _, element := range collections.Collections {
			collectionObjects = append(collectionObjects, gui.CollectionButton(element, contentBottom))
		}

		contentBottom.Objects = collectionObjects
		contentBottom.Refresh()
	})

	inputBar := widget.NewEntry()
	inputBar.SetPlaceHolder("Enter text...")
	inputBar.OnChanged = func(text string) {
		contentBottom.Objects = filterObjects(collectionObjects, text)
		contentBottom.Refresh()
	}

	contentTop := container.New(layout.NewGridLayout(2), buttonGenerate, inputBar)

	content := container.NewBorder(contentTop, nil, nil, nil, contentBottomWrapper)

	myWindow.SetContent(content)

	myWindow.Show()
	myApp.Run()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}

func filterObjects(canvasObjects []fyne.CanvasObject, text string) []fyne.CanvasObject {
	var filteredObjects []fyne.CanvasObject

	for _, canvasObject := range canvasObjects {
		if c, ok := canvasObject.(*fyne.Container); ok {
			for _, object := range c.Objects {
				if v, ok := object.(*widget.Label); ok {
					if strings.Contains(strings.ToLower(v.Text), strings.ToLower(text)) {
						filteredObjects = append(filteredObjects, canvasObject)
					}
				}
			}
		}
	}
	return filteredObjects
}
