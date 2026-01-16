package gui

import (
	"fmt"
	"maps"
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/kglaus/geodienste-cli/internal/helpers"
	"github.com/kglaus/geodienste-cli/internal/stac"
	"github.com/kglaus/geodienste-cli/internal/stac/models"
)

func NewMainFrame(bind models.State, myWindow fyne.Window) *container.Scroll {
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
					go func() {
						featureCollection, err := stac.GetItems(collection.GetItemsLink().Href)
						fyne.Do(func() {
							if err != nil {
								popup := NewErrorPopup(err.Error(), myWindow.Canvas())
								popup.Show()
							}
							items := make([]interface{}, len(featureCollection.Features))
							for i, c := range featureCollection.Features {
								items[i] = c
							}
							bind.CompleteList.Set(items)
							bind.FilteredList.Set(items)
						})
					}()
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
						// TODO Remove Printf
						fmt.Printf("call %s\n", feature.Assets[assetKey].Href)
						// Save assets
						helpers.SaveAsset()
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
