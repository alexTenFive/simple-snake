package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"snaketest/pkg/vec"
)

type (
	bodyPart struct {
		Direction      vec.Vector
		Position       vec.Vector
		GoalPositions  []vec.Vector
		GoalDirections []vec.Vector
		img            *ebiten.Image
	}

	Snake struct {
		//directionChangeLastTime time.Time
		Direction vec.Vector
		body      []*bodyPart
	}
)

const (
	partSize   = 16
	speed      = 16
	cornerEdge = 16
)

func newSnake(startPos vec.Vector, size int) *Snake {
	body := make([]*bodyPart, 0, size)
	position := startPos.X
	var im *ebiten.Image

	for i := 0; i < size; i++ {

		if i == 0 {
			im = snakeParts[SnakeTailPart]
		} else if i == size-1 {
			im = snakeParts[SnakeHeadPart]
		} else {
			im = snakeParts[SnakeBodyPart]
		}
		body = append(body, &bodyPart{
			Direction: vec.Vector{X: 1, Y: 0},
			Position:  vec.Vector{X: position, Y: startPos.Y},
			img:       im,
		})
		position += partSize
	}

	s := &Snake{
		Direction: vec.Vector{X: 1, Y: 0},
		body:      body,
	}
	return s
}
func (x *Snake) setDirection(dir vec.Vector) {
	if dir.X == 1 && x.Direction.X == -1 || dir.X == -1 && x.Direction.X == 1 ||
		dir.Y == -1 && x.Direction.Y == 1 || dir.Y == 1 && x.Direction.Y == -1 {
		return
	}

	x.Direction = dir
	x.body[len(x.body)-1].Direction = x.Direction

	for i := len(x.body) - 2; i >= 0; i-- {
		x.body[i].GoalPositions = append(x.body[i].GoalPositions, x.body[len(x.body)-1].Position)
		x.body[i].GoalDirections = append(x.body[i].GoalDirections, x.Direction)
	}
}
func (x *Snake) updateDirections() {
	for i := len(x.body) - 2; i >= 0; i-- {
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
func (x *Snake) addPart() {
	var partPosition vec.Vector
	if len(x.body) == 0 {
		return
	}

	partPosition = x.body[0].Position

	switch {
	case x.body[0].Direction.X == 1:
		partPosition.X = x.body[0].Position.X - partSize
	case x.body[0].Direction.X == -1:
		partPosition.X = x.body[0].Position.X + partSize
	case x.body[0].Direction.Y == 1:
		partPosition.Y = x.body[0].Position.Y - partSize
	case x.body[0].Direction.Y == -1:
		partPosition.Y = x.body[0].Position.Y + partSize
	}

	x.body[0].img = snakeParts[SnakeBodyPart]
	x.body = append([]*bodyPart{
		{
			Direction:      x.body[0].Direction,
			Position:       partPosition,
			GoalPositions:  x.body[0].GoalPositions,
			GoalDirections: x.body[0].GoalDirections,
			img:            snakeParts[SnakeTailPart],
		},
	}, x.body...)
}
