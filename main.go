package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const SCREEN_WIDTH = 1850;
const SCREEN_HEIGHT = 950;

func main() {
	WMRoot := WMRoot{}
	WMRoot.Node = &WMSplit{
		Type: "h",
		Nodes: []any{
			&WMNode{Panel: Panel{AppID: "author.app1", ID: "main", Title: "Panel 1", Workspace: 0}},
			&WMSplit{
				Type: "h",
				Nodes: []any{
					&WMNode{Panel: Panel{AppID: "author.app2", ID: "main", Title: "Panel 2", Workspace: 0}},
					&WMNode{Panel: Panel{AppID: "author.app3", ID: "main", Title: "Panel 3", Workspace: 0}},
				},
			},
		},
	}

	focusedPanel := "author.app1/main"

	currentChord := ""

	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Slate")
	rl.SetExitKey(rl.KeyNull)
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		Input_Update(&WMRoot, &focusedPanel, &currentChord)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		Draw(&WMRoot, focusedPanel, &currentChord)
		rl.EndDrawing()
	}
}