package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/catppuccin/fyne"
)

func NewPanel(title string, content fyne.CanvasObject) fyne.CanvasObject {
	title_label := canvas.NewText(title, color.White)
	title_label.TextStyle.Bold = true
	title_label.TextStyle.Italic = true
	title_label.TextSize = 11
	title_label.Alignment = fyne.TextAlignCenter

	header := container.NewStack(
		canvas.NewRectangle(color.NRGBA{R: 0, G: 170, B:0, A: 255}),
		container.NewPadded(title_label),
	)

	return container.NewBorder(header, nil, nil, nil, content)
}

func main() {
	app := app.New()
	ctp := catppuccin.New()
	ctp.SetFlavor(catppuccin.Mocha)
	app.Settings().SetTheme(ctp)

	window := app.NewWindow("Slate")
	window.Resize(fyne.NewSize(1800, 1000))

	t1 := NewPanel("title", widget.NewButton("1", func() {}))
	t2 := NewPanel("title", widget.NewButton("2", func() {}))
	t3 := NewPanel("title", widget.NewButton("3", func() {}))
	grid := container.New(layout.NewGridLayout(2), t1, t2, t3)

	window.SetContent(grid)
	window.ShowAndRun()
}