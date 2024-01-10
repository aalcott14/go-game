package game

import (
	"fmt"
	"image/color"
	"time"

	"github.com/aalcott14/go-game/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	screenWidth        = 800
	screenHeight       = 600
	baseMeteorVelocity = 0.25
	meteorSpawnTime    = 1 * time.Second
)

type Game struct {
	player           *Player
	bullets          []*Bullet
	meteors          []*Meteor
	baseVelocity     float64
	meteorSpawnTimer *Timer
	score            int
}

func NewGame() *Game {
	g := &Game{
		baseVelocity:     baseMeteorVelocity,
		meteorSpawnTimer: NewTimer(meteorSpawnTime),
		score:            0,
	}
	g.player = NewPlayer(g)

	return g
}

func (g *Game) Update() error {
	g.player.Update()

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()

		m := NewMeteor(g.baseVelocity)
		g.meteors = append(g.meteors, m)
	}

	for _, b := range g.bullets {
		b.Update()
	}

	for _, m := range g.meteors {
		m.Update()
	}

	for i, m := range g.meteors {
		for j, b := range g.bullets {
			if m.Collider().Intersects(b.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.bullets = append(g.bullets[:j], g.bullets[j+1:]...)
				g.score++
			}
		}
	}

	for _, m := range g.meteors {
		if m.Collider().Intersects(g.player.Collider()) {
			g.Reset()
			break
		}
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

	text.Draw(screen, fmt.Sprintf("%s\n%03d", "SCORE", g.score), assets.ScoreFont, screenWidth/2+90, 30, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) AddBullet(b *Bullet) {
	g.bullets = append(g.bullets, b)
}

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.meteors = nil
	g.bullets = nil
	g.score = 0
	g.meteorSpawnTimer.Reset()
	g.baseVelocity = baseMeteorVelocity
}
