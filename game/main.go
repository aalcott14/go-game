package main

import (
	"github.com/aalcott14/go-game/game/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()
	ebiten.SetWindowTitle("Adam's Go Game")
	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
