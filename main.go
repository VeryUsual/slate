package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Panel struct {
	AppID      string
	ID         string
	Title      string
	Workspace  int
}

type WMRoot struct {
	Node       any // WMNode or WMSplit
}

type WMNode struct {
	Panel      Panel
}

type WMSplit struct {
	Type       string
	Nodes      []any // WMNode or WMSplit
}

type DrawnPanel struct {
	X          int
	Y          int
	Width      int
	Height     int
	Panel      Panel
}

func handleSplit(n any, x int, y int, width int, height int) []DrawnPanel {
	drawnPanels := []DrawnPanel{}

	switch n_typed := n.(type) {
	case *WMNode:
		drawnPanels = append(drawnPanels, DrawnPanel{X: x, Y: y, Width: width, Height: height, Panel: n_typed.Panel})
	case *WMSplit:
		for i, node := range n_typed.Nodes {
			drawnPanels = append(drawnPanels, handleSplit(node, x + i*(width / 2), y, width / 2, height)...)
		}
	default:
		fmt.Println("this shouldnt happen")
	}

	return drawnPanels
}

func SplitFocusedPanel(n *any, focusedPanel *string) bool {
	switch n_typed := (*n).(type) {
	case *WMNode:
		if *focusedPanel == n_typed.Panel.AppID + "/" + n_typed.Panel.ID {
			*n = &WMSplit{
				Type: "h",
				Nodes: []any{
					n_typed,
					&WMNode{
						Panel: Panel{
							AppID: "author.app99",
							ID: "main",
							Title: "panel that got split",
						},
					},
				},
			}

			return true
		}
	case *WMSplit:
		for i := range n_typed.Nodes {
			if SplitFocusedPanel(&n_typed.Nodes[i], focusedPanel) {
				return true
			}
		}
	}

	return false
}

func Update(WMRoot *WMRoot, focusedPanel *string) {
	if rl.IsKeyPressed(rl.KeyS) {
		SplitFocusedPanel(&WMRoot.Node, focusedPanel)
	}
}

func Draw(WMRoot *WMRoot, focusedPanel string) {
	drawnPanels := handleSplit(WMRoot.Node, 0, 0, 1850, 950)
	for _, p := range drawnPanels {
		rl.DrawRectangle(int32(p.X), int32(p.Y), int32(p.Width), int32(p.Height), rl.Color{R: 37, G: 39, B: 57, A: 255})
		
		if focusedPanel == p.Panel.AppID + "/" + p.Panel.ID {
			rl.DrawRectangleLinesEx(
				rl.Rectangle{X: float32(p.X), Y: float32(p.Y), Width: float32(p.Width), Height: float32(p.Height)},
				1,
				rl.Color{R: 0, G: 255, B: 0, A: 255},
			)
		} else {
			rl.DrawRectangleLinesEx(
				rl.Rectangle{X: float32(p.X), Y: float32(p.Y), Width: float32(p.Width), Height: float32(p.Height)},
				0.5,
				rl.Color{R: 128, G: 128, B: 128, A: 128},
			)
		}

		rl.DrawRectangle(
			int32(p.X),
			int32(p.Y),
			int32(p.Width),
			int32(30),
			rl.Color{R: 31, G: 32, B: 47, A: 255},
		)
		rl.DrawText(
			p.Panel.Title,
			int32(int32(p.X + (p.Width / 2)) - (rl.MeasureText(p.Panel.Title, int32(12)) / int32(2))),
			int32(p.Y + 10),
			int32(12),
			rl.Color{R: 255, G: 255, B: 255, A: 255},
		)
	}
}

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

	rl.InitWindow(1850, 950, "Slate")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		Update(&WMRoot, &focusedPanel)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		Draw(&WMRoot, focusedPanel)
		rl.EndDrawing()
	}
}