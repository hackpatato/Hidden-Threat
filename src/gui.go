package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)
func Gui() {
myApp := app.New()
	myWindow := myApp.NewWindow("dowlanding...")


	myWindow.Resize(fyne.NewSize(500, 600))
	myWindow.SetFixedSize(true)

	logo := canvas.NewImageFromFile("logo.png")
	logo.FillMode = canvas.ImageFillContain
	logo.SetMinSize(fyne.NewSize(100, 100))


	statusLabel := widget.NewLabel("The system is being prepared, please wait...")
	statusLabel.Alignment = fyne.TextAlignCenter

	progress := widget.NewProgressBarInfinite()


	content := container.NewVBox(
		container.NewPadded(logo),
		widget.NewSeparator(),
		container.NewPadded(statusLabel),
		container.NewPadded(progress),
	)

	centeredContent := container.NewCenter(content)

	myWindow.SetContent(centeredContent)
	myWindow.ShowAndRun()
}