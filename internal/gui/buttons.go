package gui

import (
	"fyne.io/fyne/v2/widget"
	"github.com/kglaus/geodienste-cli/internal/stac"
	"github.com/kglaus/geodienste-cli/internal/stac/models"
)

func NewStacButton(bind models.State, collections models.Collections, selectEntry *widget.SelectEntry) *widget.Button {
	buttonGenerate := widget.NewButton("Search", func() {
		collections = stac.GetCollections(selectEntry.Text)
		items := make([]interface{}, len(collections.Collections))
		for i, c := range collections.Collections {
			items[i] = c
		}
		bind.CompleteList.Set(items)
		bind.FilteredList.Set(items)
	})

	return buttonGenerate
}
