package main

import (
	"fmt"
	"maps"
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/kglaus/geodienste-cli/pkg/gui"
	"github.com/kglaus/geodienste-cli/pkg/stac"
	"github.com/kglaus/geodienste-cli/pkg/stac/models"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("geodienst-cli")
	myWindow.SetContent(widget.NewLabel("geodienste-cli2"))
	myWindow.Resize(fyne.NewSize(1050, 400))

	stateBindings := binding.NewUntypedList()
	contentBottom := newMainFrame(stateBindings)
	contentBottomWrapper := container.NewVScroll(contentBottom)

	var collectionObjects []fyne.CanvasObject

	// Search bar text
	inputBar := widget.NewEntry()
	inputBar.SetPlaceHolder("Enter text...")
	inputBar.OnChanged = func(text string) {
		contentBottom.Objects = gui.FilterCanvasObjects(collectionObjects, text)
		contentBottom.Refresh()
	}

	buttonGenerate := widget.NewButton("click me", func() {
		collections := stac.GetCollections()
		items := make([]interface{}, len(collections.Collections))
		for i, c := range collections.Collections {
			items[i] = c
		}
		stateBindings.Set(items)
	})

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

func newMainFrame(bind binding.UntypedList) *fyne.Container {
	grid := container.New(layout.NewGridWrapLayout(fyne.Size{Width: 200, Height: 125}))

	refresh := func() {
		grid.Objects = nil

		state, _ := bind.Get()
		for _, item := range state {
			if collection, ok := item.(models.Collection); ok {

				label := widget.NewLabel(collection.Title)
				label.Wrapping = fyne.TextWrapWord

				button := widget.NewButton("", func() {
					featureCollection := stac.GetItems(collection.GetItemsLink().Href)
					items := make([]interface{}, len(featureCollection.Features))
					for i, c := range featureCollection.Features {
						items[i] = c
					}
					bind.Set(items)
				})

				stack := container.NewStack(button, label)

				grid.Add(stack)

			} else if feature, ok := item.(models.Feature); ok {

				assetKeys := slices.Sorted(maps.Keys(feature.Assets))
				for _, assetKey := range assetKeys {
					label := widget.NewLabel(feature.Assets[assetKey].Title)
					label.Wrapping = fyne.TextWrapWord

					button := widget.NewButton("", func() {
						fmt.Printf("call %s\n", feature.Assets[assetKey].Href)
					})

					stack := container.NewStack(button, label)

					grid.Add(stack)
				}
			}
		}
	}

	bind.AddListener(binding.NewDataListener(func() {
		refresh()
	}))

	refresh()

	return grid
}
