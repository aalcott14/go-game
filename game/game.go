package game

import (
	"game/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	playerPosition Vector
}

func NewGame() *Game {
	g := &Game{
		playerPosition: Vector{X: 100, Y: 100},
	}

	return g
}

func (g *Game) Update() error {
	speed := 5.0

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.playerPosition.Y += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.playerPosition.Y -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.playerPosition.X -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.playerPosition.X += speed
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.playerPosition.X, g.playerPosition.Y)
	screen.DrawImage(assets.PlayerSprite, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
