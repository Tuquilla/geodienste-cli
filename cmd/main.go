package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/kglaus/geodienste-cli/pkg/stac"
	"github.com/kglaus/geodienste-cli/pkg/stac/models"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("geodienst-cli")
	myWindow.SetContent(widget.NewLabel("geodienste-cli2"))
	myWindow.Resize(fyne.NewSize(1050, 400))

	stateBindings := models.State{
		CompleteList: binding.NewUntypedList(),
		FilteredList: binding.NewUntypedList(),
		Search:       binding.NewString(),
	}

	//contentBottom := newMainFrame(stateBindings)
	//contentBottomWrapper := container.NewVScroll(contentBottom)
	contentBottomWrapper := newMainFrame(stateBindings)

	var collections models.Collections

	// Search bar text
	inputBar := widget.NewEntry()
	inputBar.SetPlaceHolder("Enter text...")
	inputBar.OnChanged = func(text string) {
		stateBindings.Search.Set(text)
		items := make([]any, 0, len(collections.Collections))
		completeList, _ := stateBindings.CompleteList.Get()
		for _, element := range completeList {
			if collection, ok := element.(models.Collection); ok {
				if strings.Contains(strings.ToLower(collection.Title), strings.ToLower(text)) {
					items = append(items, collection)
				}
			}
			if feature, ok := element.(models.Feature); ok {
				assetKeys := slices.Sorted(maps.Keys(feature.Assets))
				for _, assetKey := range assetKeys {
					if strings.Contains(strings.ToLower(feature.Assets[assetKey].Title), strings.ToLower(text)) {
						items = append(items, feature)
					}
				}

			}
		}
		stateBindings.FilteredList.Set(items)
	}

	buttonGenerate := widget.NewButton("click me", func() {
		collections = stac.GetCollections()
		items := make([]interface{}, len(collections.Collections))
		for i, c := range collections.Collections {
			items[i] = c
		}
		stateBindings.CompleteList.Set(items)
		stateBindings.FilteredList.Set(items)
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

func newMainFrame(bind models.State) *container.Scroll {
	grid := container.New(layout.NewGridWrapLayout(fyne.Size{Width: 200, Height: 125}))
	scrollableGrid := container.NewVScroll(grid)

	refresh := func() {
		grid.Objects = nil

		state, _ := bind.FilteredList.Get()
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
					bind.CompleteList.Set(items)
					bind.FilteredList.Set(items)
				})

				stack := container.NewStack(button, label)

				grid.Add(stack)

			} else if feature, ok := item.(models.Feature); ok {

				assetKeys := slices.Sorted(maps.Keys(feature.Assets))
				for _, assetKey := range assetKeys {
					var label *widget.Label
					if len(feature.Assets[assetKey].Title) > 0 {
						label = widget.NewLabel(feature.Assets[assetKey].Title)
					} else {
						label = widget.NewLabel(assetKey)
					}
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

	bind.FilteredList.AddListener(binding.NewDataListener(func() {
		refresh()
		// reset scroll position to top
		scrollableGrid.Offset = fyne.NewPos(0, 0)
	}))

	refresh()

	return scrollableGrid
}
