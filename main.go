package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const (
	screenWidth = 1280
	screenHeight = 720
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("GUI Editor")
	ebiten.SetCursorMode(ebiten.CursorModeVisible)

	g := NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}