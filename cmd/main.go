package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/kglaus/geodienste-cli/pkg/stac"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("geodienst-cli")
	myWindow.SetContent(widget.NewLabel("geodienste-cli2"))
	myWindow.Resize(fyne.NewSize(300, 600))

	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

	canvasObjects := []fyne.CanvasObject{}

	text1 := canvas.NewText("Part1", green)
	text2 := canvas.NewText("Part2", green)
	text3 := canvas.NewText("Part3", color.White)
	text4 := canvas.NewText("Part4", color.White)

	canvasObjects = append(canvasObjects, text1, text2, text3, text4)

	content_bottom := container.New(layout.NewGridLayout(2), canvasObjects...)

	button_generate := widget.NewButton("click me", func() {
		content_bottom.Objects = []fyne.CanvasObject{canvas.NewText("Part4", color.White), canvas.NewText("Part5", color.White)}
		content_bottom.Refresh()
		stac.TestStac()
	})

	content_top := container.New(layout.NewGridLayout(1), button_generate)

	content := container.New(layout.NewGridLayout(1), content_top, content_bottom)

	myWindow.SetContent(content)

	myWindow.Show()
	myApp.Run()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}
