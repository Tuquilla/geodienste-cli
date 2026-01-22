package gui

import (
	"maps"
	"slices"
	"strings"

	"fyne.io/fyne/v2/widget"
	"github.com/kglaus/stac-client/internal/helpers"
	"github.com/kglaus/stac-client/internal/stac/models"
)

func NewSearchBar(bind models.State) *widget.Entry {
	searchBar := widget.NewEntry()
	searchBar.SetPlaceHolder("Search in Assets")
	searchBar.OnChanged = func(text string) {
		bind.Search.Set(text)
		completeList, _ := bind.CompleteList.Get()
		items := make([]any, 0, len(completeList))
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

func NewDirBar() *widget.Entry {
	dirBar := widget.NewEntry()
	_, assetRootPath, _ := helpers.OperatingSystem()
	dirBar.Text = assetRootPath
	// TODO implement this path for saving assets
	return dirBar
}
