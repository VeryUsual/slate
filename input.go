package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var KEYBINDS = map[string]string{
	" sv": "split panel vertically",
	" qq": "quit slate",
}

func Input_Update(WMRoot *WMRoot, focusedPanel *string, currentChord *string, quitSlate *bool) {
	if rl.IsKeyPressed(rl.KeyEscape) && *currentChord != "" {
		*currentChord = ""
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		*currentChord += " "
	}
	if rl.IsKeyPressed(rl.KeyS) {
		*currentChord += "s"
	}
	if rl.IsKeyPressed(rl.KeyV) {
		*currentChord += "v"
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		*currentChord += "q"
	}
	
	if *currentChord == " qq" {
		*quitSlate = true
	}
	if *currentChord == " sv" {
		*currentChord = ""
		SplitFocusedPanel(&WMRoot.Node, focusedPanel)
	}
}