package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Keybind struct {
	Explanation string
	Action      func()
}

var KEYBINDS = map[string]*Keybind{
	" sv": &Keybind{Explanation: "Split Panel Vertically", Action: nil},
	" qq": &Keybind{Explanation: "Quit Slate", Action: nil},
}

func Input_Update(WMRoot *WMRoot, focusedPanel *string, currentChord *string, quitSlate *bool) {
	KEYBINDS[" sv"].Action = func() { SplitFocusedPanel(&WMRoot.Node, focusedPanel) }
	KEYBINDS[" qq"].Action = func() { *quitSlate = true }

	if *currentChord != "" {
		if rl.IsKeyPressed(rl.KeyEscape) {
			*currentChord = ""
		}
		if rl.IsKeyPressed(rl.KeyBackspace) {
			*currentChord = (*currentChord)[:len(*currentChord)-1]
		}
	}

	if char := rl.GetCharPressed(); char > 0 {
		c := rune(char)
		*currentChord += string(c)
	}

	for bind, obj := range KEYBINDS {
		if *currentChord == bind {
			*currentChord = ""
			obj.Action();
		}
	}
}