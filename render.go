package main

import (
	"sort"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Draw(WMRoot *WMRoot, focusedPanel string, currentChord *string) {
	drawnPanels := handleSplit(WMRoot.Node, 0, 0, SCREEN_WIDTH, SCREEN_HEIGHT)
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

	keybindKeys := make([]string, 0, len(KEYBINDS))
	for k := range KEYBINDS {
		keybindKeys = append(keybindKeys, k)
	}
	sort.Strings(keybindKeys)

	keybindHints := ""
	for _, keybind := range keybindKeys {
		explanation := KEYBINDS[keybind]
		if strings.HasPrefix(keybind, *currentChord) && *currentChord != "" {
			keybindHints += " | '" + strings.ReplaceAll(keybind, " ", "<space>") + "': " + explanation
		}
	}

	if keybindHints != "" || *currentChord != "" {
		rl.DrawText(
			strings.ReplaceAll(*currentChord, " ", "<space>") + "\n" + keybindHints,
			35,
			SCREEN_HEIGHT - 55,
			int32(25),
			rl.Color{R: 200, G: 200, B: 255, A: 255},
		)
	}
}