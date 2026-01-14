package gui

import (
	"maps"
	"slices"
	"strings"

	"fyne.io/fyne/v2/widget"
	"github.com/kglaus/geodienste-cli/internal/stac/models"
)

func NewSearchBar(bind models.State, collections models.Collections) *widget.Entry {
	searchBar := widget.NewEntry()
	searchBar.SetPlaceHolder("Search Assets")
	searchBar.OnChanged = func(text string) {
		bind.Search.Set(text)
		items := make([]any, 0, len(collections.Collections))
		completeList, _ := bind.CompleteList.Get()
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
		bind.FilteredList.Set(items)
	}
	return searchBar
}

func NewSelectEntry(selectOptions []string) *widget.SelectEntry {
	selectEntry := widget.NewSelectEntry(selectOptions)
	// Preselect first option
	selectEntry.Text = selectOptions[0]
	return selectEntry
}
