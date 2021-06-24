package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
)

var snakeParts map[int64]*ebiten.Image

const (
	SnakeHeadPart = iota
	SnakeBodyPart
	SnakeTailPart
	ApplePart
)

func init() {
	snakeParts = make(map[int64]*ebiten.Image, 3)
}

func loadSnakeSprites() error {
	var err error
	snakeParts[SnakeHeadPart], _, err = ebitenutil.NewImageFromFile("assets/sprites/snake_head.png")
	if err != nil {
		return err
	}
	snakeParts[SnakeBodyPart], _, err = ebitenutil.NewImageFromFile("assets/sprites/snake_body.png")
	if err != nil {
		return err
	}
	snakeParts[SnakeTailPart], _, err = ebitenutil.NewImageFromFile("assets/sprites/snake_tail.png")
	if err != nil {
		return err
	}
	snakeParts[ApplePart], _, err = ebitenutil.NewImageFromFile("assets/sprites/an_apple.png")
	if err != nil {
		return err
	}

	return nil
}
