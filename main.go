package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Panel struct {
	AppID      string
	ID         string
	Title      string
	Workspace  int
}

var panels = []Panel{}

/*
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
*/

func main() {
	panels = []Panel{
		Panel{AppID: "author.app1", ID: "main", Title: "1", Workspace: 0},
		Panel{AppID: "author.app2", ID: "main", Title: "2", Workspace: 0},
		Panel{AppID: "author.app3", ID: "main", Title: "3", Workspace: 0},
	}

	rl.InitWindow(1800, 1000, "Slate")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		i := 0
		for row := 0; row < 2; row++ {
			for col := 0; col < 2; col++ {
				if len(panels) - 1 < i {
					break
				}
				x := col * 500
				y := row * 500
				rl.DrawRectangle(int32(x), int32(y), int32(500), int32(500), rl.Color{R: 0, G: 255, B: 0, A: 255})
				rl.DrawText(panels[i].Title, int32(x + 75), int32(y + 25), int32(20), rl.Color{R: 255, G: 0, B: 0, A: 255})
				i+=1
			}
		}

		rl.EndDrawing()
	}
}