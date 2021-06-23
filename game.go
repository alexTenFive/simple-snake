package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"snaketest/pkg/vec"
)

type Game struct {
	Snake *Snake
}

func NewGame() *Game {
	return &Game{Snake: newSnake(vec.Vector{X: 100, Y: 100}, 5)}
}
func (g *Game) Start() {
	g.Snake.Start()
}

func (g *Game) Update() error {
	g.Snake.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Snake.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth / 2, screenHeight / 2
}



