package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/kglaus/stac-client/configs"
	"github.com/kglaus/stac-client/internal/gui"
	"github.com/kglaus/stac-client/internal/stac/models"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("stac-client")
	myWindow.Resize(fyne.NewSize(1050, 400))
	myWindow.SetContent(widget.NewLabel("stac-client"))

	stateBindings := models.State{
		CompleteList: binding.NewUntypedList(),
		FilteredList: binding.NewUntypedList(),
		Search:       binding.NewString(),
	}

	selectOptions := configs.ReadConfig().BaseUrls

	// Mainframe
	contentBottomWrapper := gui.NewMainFrame(stateBindings, myWindow)

	// Search bar
	searchBar := gui.NewSearchBar(stateBindings)

	// Select entry
	selectEntry := gui.NewSelectEntry(selectOptions)

	// Dir bar
	dirBar := gui.NewDirBar()

	// Stac button
	buttonGenerate := gui.NewStacButton(stateBindings, selectEntry, myWindow)

	// Arrange widgets
	contentTop := container.New(layout.NewGridLayout(2), selectEntry, buttonGenerate, searchBar, dirBar)
	content := container.NewBorder(contentTop, nil, nil, nil, contentBottomWrapper)

	myWindow.SetContent(content)

	myWindow.Show()
	myApp.Run()
	onClose()
}

func onClose() {
	fmt.Println("Exited")
}
