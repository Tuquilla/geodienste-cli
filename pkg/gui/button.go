package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/kglaus/geodienste-cli/pkg/stac/models"
)

func CollectionButton(collection models.Collection, contentBottom *fyne.Container) *widget.Button {
	collectionButton := widget.NewButton(collection.Id, func() {

	})
	return collectionButton
}
