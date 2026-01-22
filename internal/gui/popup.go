package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/image/colornames"
)

func NewErrorPopup(errorText string, myCanvas fyne.Canvas) *widget.PopUp {
	label := canvas.NewText(errorText, colornames.Red)
	popup := widget.NewPopUp(label, myCanvas)
	return popup
}

func NewSuccessfulPopup(message string, myCanvas fyne.Canvas) *widget.PopUp {
	label := canvas.NewText(message, colornames.Green)
	popup := widget.NewPopUp(label, myCanvas)
	return popup
}
