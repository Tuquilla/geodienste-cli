package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/kglaus/geodienste-cli/internal/stac"
	"github.com/kglaus/geodienste-cli/internal/stac/models"
)

func NewStacButton(bind models.State, selectEntry *widget.SelectEntry, myWindow fyne.Window) *widget.Button {
	buttonGenerate := widget.NewButton("Search", func() {
		go func() {
			collections, err := stac.GetCollections(selectEntry.Text)
			fyne.Do(func() {
				if err != nil {
					popup := NewErrorPopup(err.Error(), myWindow.Canvas())
					popup.Show()
				}
				items := make([]interface{}, len(collections.Collections))
				for i, c := range collections.Collections {
					items[i] = c
				}
				bind.CompleteList.Set(items)
				bind.FilteredList.Set(items)
			})
		}()
	})

	return buttonGenerate
}
