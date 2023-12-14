package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth        = 800
	screenHeight       = 600
	baseMeteorVelocity = 0.25
)

type Game struct {
	player       *Player
	bullets      []*Bullet
	meteors      []*Meteor
	baseVelocity float64
}

func NewGame() *Game {
	g := &Game{
		baseVelocity: baseMeteorVelocity,
	}
	g.player = NewPlayer(g)

	return g
}

func (g *Game) Update() error {
	g.player.Update()

	for _, b := range g.bullets {
		b.Update()
	}

	for _, m := range g.meteors {
		m.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, b := range g.bullets {
		b.Draw(screen)
	}

	for _, m := range g.meteors {
		m.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) AddBullet(b *Bullet) {
	g.bullets = append(g.bullets, b)
}
