package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/catppuccin/fyne"
)

func main() {
	app := app.New()
	ctp := catppuccin.New()
	ctp.SetFlavor(catppuccin.Mocha)
	app.Settings().SetTheme(ctp)

	window := app.NewWindow("Slate")
	window.Resize(fyne.NewSize(1800, 1000))

	t1 := widget.NewButton("1", func() {})
	t2 := widget.NewButton("2", func() {})
	t3 := widget.NewButton("3", func() {})
	grid := container.New(layout.NewGridLayout(2), t1, t2, t3)

	window.SetContent(grid)
	window.ShowAndRun()
}