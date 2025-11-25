package gui

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kglaus/geodienste-cli/pkg/stac"
	"github.com/kglaus/geodienste-cli/pkg/stac/models"
)

func CollectionButton(collection models.Collection, contentBottom *fyne.Container, inputBar *widget.Entry) *fyne.Container {

	label := widget.NewLabel(collection.Title)
	label.Wrapping = fyne.TextWrapWord

	collectionButton := widget.NewButton("", func() {
		featureCollection := stac.GetItems(collection.GetItemsLink().Href)

		var featureCollectionObjects []fyne.CanvasObject

		for _, element := range featureCollection.Features {

			assetKeys := slices.Sorted(maps.Keys(element.Assets))

			for _, assetKey := range assetKeys {
				featureCollectionObjects = append(featureCollectionObjects, DownloadButton(element.Assets[assetKey]))
			}
		}

		contentBottom.Objects = featureCollectionObjects

		inputBar.OnChanged = func(text string) {
			contentBottom.Objects = FilterCanvasObjects(featureCollectionObjects, text)
			contentBottom.Refresh()
		}

		contentBottom.Refresh()

	})
	stack := container.NewStack(collectionButton, label)

	return stack
}

func DownloadButton(assetObject models.AssetObject) *fyne.Container {
	label := widget.NewLabel(assetObject.Title)
	label.Wrapping = fyne.TextWrapWord

	downloadButton := widget.NewButton("", func() {
		fmt.Printf("call %s\n", assetObject.Href)
	})
	stack := container.NewStack(downloadButton, label)
	return stack
}

func FilterCanvasObjects(canvasObjects []fyne.CanvasObject, text string) []fyne.CanvasObject {
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
