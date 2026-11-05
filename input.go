package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var KEYBINDS = map[string]string{
	" sv": "split panel vertically",
}

func Input_Update(WMRoot *WMRoot, focusedPanel *string, currentChord *string) {
	if rl.IsKeyPressed(rl.KeySpace) {
		*currentChord += " "
	}
	if rl.IsKeyPressed(rl.KeyS) {
		*currentChord += "s"
	}
	if *currentChord == " s" && rl.IsKeyPressed(rl.KeyV) {
		*currentChord = ""
		SplitFocusedPanel(&WMRoot.Node, focusedPanel)
	}
}