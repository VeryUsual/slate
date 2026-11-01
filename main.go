package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	catppuccin "github.com/catppuccin/fyne"
)

type Panel struct {
	AppID      string
	ID         string
	Title      string
	Content    fyne.CanvasObject
	Workspace  int
}

var panels = []Panel{}

func ClosePanel(app_id string, id string) []Panel {
	newPanel := []Panel{}

	for _, p := range panels {
		if !(p.AppID == app_id && p.ID == id) {
			newPanel = append(newPanel, p)
		}
	}

	return newPanel
}

func DrawPanel(panel *Panel, onClose func(fyne.CanvasObject)) fyne.CanvasObject {
	var p fyne.CanvasObject
	
	title_label := canvas.NewText(panel.Title, color.White)
	title_label.TextStyle.Bold = true
	title_label.TextStyle.Italic = true
	title_label.TextSize = 11
	title_label.Alignment = fyne.TextAlignCenter

	closeButton := widget.NewButton("x", func() {
		onClose(p)
	})

	header := container.NewStack(
		canvas.NewRectangle(color.NRGBA{R: 0, G: 170, B:0, A: 255}),
		container.NewCenter(container.NewPadded(title_label)),
		container.NewBorder(nil, nil, nil, closeButton),
	)

	p = container.NewBorder(header, nil, nil, nil, panel.Content)
	return p
}

func main() {
	app := app.New()
	ctp := catppuccin.New()
	ctp.SetFlavor(catppuccin.Mocha)
	app.Settings().SetTheme(ctp)

	window := app.NewWindow("Slate")
	window.Resize(fyne.NewSize(1800, 1000))

	panels = []Panel{
		Panel{AppID: "author.app1", ID: "main", Title: "1", Content: widget.NewButton("1", func() {}), Workspace: 0},
		Panel{AppID: "author.app2", ID: "main", Title: "2", Content: widget.NewButton("2", func() {}), Workspace: 0},
		Panel{AppID: "author.app3", ID: "main", Title: "3", Content: widget.NewButton("3", func() {}), Workspace: 0},
	}

	grid := container.New(layout.NewGridLayout(2))

	for i := range panels {
		panel := panels[i]

		appid, id := panel.AppID, panel.ID

		paneldrawn := DrawPanel(&panel, func(panel fyne.CanvasObject) {
			grid.Remove(panel)
			panels = ClosePanel(appid, id)
			grid.Refresh()
		})

		grid.Add(paneldrawn)
	}

	window.SetContent(grid)
	window.ShowAndRun()
}