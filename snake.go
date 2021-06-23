package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
	"snaketest/pkg/vec"
	"time"
)

type (
	bodyPart struct {
		Direction vec.Vector
		Position  vec.Vector
		GoalPositions []vec.Vector
		GoalDirections []vec.Vector
		img       *ebiten.Image
	}
	Snake struct {
		Direction vec.Vector
		body      []*bodyPart
	}
)

const(
	partSize = 16
	speed = 4
)

func newSnake(startPos vec.Vector, size int) *Snake {
	body := make([]*bodyPart, 0, size)
	position := startPos.X
	var im *ebiten.Image

	for i := 0; i < size; i++ {
		im = ebiten.NewImage(partSize, partSize)
		im.Fill(colornames.Snow)
		if i == size-1 {
			im.Fill(colornames.Red)
		}
		body = append(body, &bodyPart{
			Direction: vec.Vector{X: 1},
			Position:  vec.Vector{X: position, Y: startPos.Y},
			img:       im,
		})
		position += partSize
	}

	s := &Snake{
		body: body,
	}
	return s
}
func (x *Snake) setDirection(dir vec.Vector) {
	x.Direction = dir
	x.body[len(x.body)-1].Direction = x.Direction

	for i := len(x.body)-2; i >= 0; i-- {
		x.body[i].GoalPositions = append(x.body[i].GoalPositions, x.body[len(x.body)-1].Position)
		x.body[i].GoalDirections = append(x.body[i].GoalDirections, x.Direction)
	}
}

func (x *Snake) updateDirections() {
	for i := len(x.body)-2; i >= 0; i-- {
		if len(x.body[i].GoalPositions) == 0 {
			continue
		}
		if x.body[i].Position == x.body[i].GoalPositions[0] {
			x.body[i].Direction = x.body[i].GoalDirections[0]
			x.body[i].GoalPositions = x.body[i].GoalPositions[1:]
			x.body[i].GoalDirections = x.body[i].GoalDirections[1:]
		}
	}
}


func (x *Snake) Start() {

}

func (x *Snake) Update() error {
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyW):
		x.setDirection(vec.Vector{0, -1})
	case inpututil.IsKeyJustPressed(ebiten.KeyS):
		x.setDirection(vec.Vector{X: 0, Y: 1})
	case inpututil.IsKeyJustPressed(ebiten.KeyA):
		x.setDirection(vec.Vector{X: -1, Y: 0})
	case inpututil.IsKeyJustPressed(ebiten.KeyD):
		x.setDirection(vec.Vector{X: 1, Y: 0})
	}
	x.updateDirections()

	return nil
}

func (x *Snake) Draw(screen *ebiten.Image) {
	var op *ebiten.DrawImageOptions
	for _, b := range x.body {
		op = &ebiten.DrawImageOptions{}
		switch {
		case b.Direction.X == 1:
			b.Position.X += speed
		case b.Direction.X == -1:
			b.Position.X -= speed
		case b.Direction.Y == 1:
			b.Position.Y += speed
		case b.Direction.Y == -1:
			b.Position.Y -= speed
		}

		op.GeoM.Translate(b.Position.X, b.Position.Y)

		screen.DrawImage(b.img, op)
		time.Sleep(time.Microsecond * 30)
	}
}


