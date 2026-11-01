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

func main() {
	panels = []Panel{
		Panel{AppID: "author.app1", ID: "main", Title: "Panel 1", Workspace: 0},
		Panel{AppID: "author.app2", ID: "main", Title: "Panel 2", Workspace: 0},
		Panel{AppID: "author.app3", ID: "main", Title: "Panel 3", Workspace: 0},
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
				rl.DrawRectangle(int32(x), int32(y), int32(500), int32(500), rl.Color{R: 37, G: 39, B: 57, A: 255})
				rl.DrawRectangleLinesEx(rl.Rectangle{X: float32(x), Y: float32(y), Width: float32(500), Height: float32(500)}, 0.5, rl.Color{R: 0, G: 0, B: 255, A: 255})
				rl.DrawRectangle(
					int32(x),
					int32(y),
					int32(500),
					int32(30),
					rl.Color{R: 31, G: 32, B: 47, A: 255},
				)
				rl.DrawText(
					panels[i].Title,
					int32(int32(x + 250) - (rl.MeasureText(panels[i].Title, int32(12)) / int32(2))),
					int32(y + 10),
					int32(12),
					rl.Color{R: 255, G: 255, B: 255, A: 255},
				)
				i+=1
			}
		}

		rl.EndDrawing()
	}
}