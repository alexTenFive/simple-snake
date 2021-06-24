package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/exp/rand"
	"snaketest/pkg/vec"
)

type (
	Apple struct {
		img *ebiten.Image
		R float64
		Position vec.Vector
	}
)

func generateApple() *Apple {
	xr := rand.Intn(screenWidth)
	yr := rand.Intn(screenHeight)
	im := snakeParts[ApplePart]

	return &Apple{Position: vec.Vector{X: float64(xr-(xr%partSize)), Y: float64(yr-(yr%partSize))}, img: im}
}
