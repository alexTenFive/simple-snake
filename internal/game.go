package internal

import (
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"io/ioutil"
	"math"
	"snaketest/pkg/vec"
	"time"
)

type Game struct {
	snake           *Snake
	apple           *Apple
	appleCounter    int64
	delta           int64
	lastMove        time.Time
	gameOverFont    font.Face
	infoFont        font.Face
	isGameOver      bool
	isGameOverBlink bool
	lastBlink       time.Time
}

const (
	gameSpeed      = 26
	snakeStartSize = 5
	blinkSpeed     = time.Millisecond * 10
	screenWidth    = 1280
	screenHeight   = 720
)

func NewGame() *Game {
	fontData, err := ioutil.ReadFile("../assets/font/Etastro.ttf")
	if err != nil {
		panic("cannot load asset font")
	}

	ttfFont, err := truetype.Parse(fontData)
	if err != nil {
		panic("cannot parse asset font")
	}

	if err := loadSnakeSprites(); err != nil {
		panic(fmt.Sprintf("cannot load snake sprites: %s", err))
	}

	return &Game{
		snake:        newSnake(vec.Vector{X: partSize * 3, Y: partSize * 3}, snakeStartSize),
		apple:        generateApple(),
		appleCounter: 0,
		isGameOver:   false,
		gameOverFont: truetype.NewFace(ttfFont, &truetype.Options{
			Size:    72,
			DPI:     72,
			Hinting: font.HintingFull,
		}),
		infoFont: truetype.NewFace(ttfFont, &truetype.Options{
			Size:    18,
			DPI:     72,
			Hinting: font.HintingFull,
		}),
	}
}

func (g *Game) Update() error {
	if g.isGameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			g.snake = newSnake(vec.Vector{X: partSize * 3, Y: partSize * 3}, snakeStartSize)
			g.apple = generateApple()
			g.isGameOver = false
			g.appleCounter = 0
		}
		return nil
	}

	// TODO: two directions simultaniously leads to bug and eat snakeself
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyW):
		g.snake.setDirection(vec.Vector{X: 0, Y: -1})
	case inpututil.IsKeyJustPressed(ebiten.KeyS):
		g.snake.setDirection(vec.Vector{X: 0, Y: 1})
	case inpututil.IsKeyJustPressed(ebiten.KeyA):
		g.snake.setDirection(vec.Vector{X: -1, Y: 0})
	case inpututil.IsKeyJustPressed(ebiten.KeyD):
		g.snake.setDirection(vec.Vector{X: 1, Y: 0})
	}
	g.snake.updateDirections()

	if time.Since(g.lastMove) < time.Second/gameSpeed {
		return nil
	}
	g.lastMove = time.Now()

	head := g.snake.body[len(g.snake.body)-1]

	for i, b := range g.snake.body {
		if i != len(g.snake.body)-1 && head.Position == b.Position {
			g.isGameOver = true
			break
		}

		b.Position.X += b.Direction.X * speed
		b.Position.Y += b.Direction.Y * speed

		if b.Position.X > screenWidth {
			b.Position.X = float64(int(b.Position.X)%screenWidth) - cornerEdge
		}
		if b.Position.Y > screenHeight {
			b.Position.Y = float64(int(b.Position.Y)%screenHeight) - cornerEdge
		}

		if b.Position.X < -cornerEdge {
			b.Position.X += screenWidth + cornerEdge
		}
		if b.Position.Y < -cornerEdge {
			b.Position.Y += screenHeight + cornerEdge
		}
	}

	if g.apple != nil {
		if head.Position.X >= g.apple.Position.X && head.Position.X < g.apple.Position.X+partSize &&
			head.Position.Y >= g.apple.Position.Y && head.Position.Y < g.apple.Position.Y+partSize {
			g.snake.addPart()
			g.apple = generateApple()
			g.appleCounter++
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var op *ebiten.DrawImageOptions
	for _, b := range g.snake.body {
		op = &ebiten.DrawImageOptions{}

		op.GeoM.Translate(-float64(partSize/2), -float64(partSize/2))

		var degrees float64
		switch {
		case b.Direction.X == 1:
		case b.Direction.X == -1:
			degrees = 180
		case b.Direction.Y == 1:
			degrees = 90
		case b.Direction.Y == -1:
			degrees = 270
		}

		op.GeoM.Rotate(degrees * (2 * math.Pi / 360))
		op.GeoM.Translate(b.Position.X, b.Position.Y)

		screen.DrawImage(b.img, op)
	}
	// draw apple
	if g.apple != nil {
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-float64(partSize/2), -float64(partSize/2))
		op.GeoM.Rotate(g.apple.R * (2 * math.Pi / 360))
		op.GeoM.Translate(g.apple.Position.X, g.apple.Position.Y)
		screen.DrawImage(g.apple.img, op)
		g.apple.R++
		if g.apple.R > 360 {
			g.apple.R = 0
		}
	}

	if g.isGameOver {
		if time.Since(g.lastBlink) < blinkSpeed {
			return
		}
		if g.isGameOverBlink {
			text.Draw(screen, "GAME OVER", g.gameOverFont, screenWidth/3, screenHeight/2, colornames.Black)
			text.Draw(screen, "Press \"R\" to play another one", g.infoFont, screenWidth/2-(screenWidth/12), screenHeight/2+30, colornames.Black)
			g.isGameOverBlink = false
			g.lastBlink = time.Now()
			return
		}
		text.Draw(screen, "GAME OVER", g.gameOverFont, screenWidth/3, screenHeight/2, colornames.Indianred)
		text.Draw(screen, "Press \"R\" to play another one", g.infoFont, screenWidth/2-(screenWidth/12), screenHeight/2+30, colornames.Indianred)
		g.isGameOverBlink = true
		g.lastBlink = time.Now()
		return
	}

	text.Draw(screen, fmt.Sprintf("scores: %d", g.appleCounter), g.infoFont, 10, 20, colornames.Lightyellow)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
