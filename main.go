package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/omar0ali/pong/entities"
	"github.com/omar0ali/pong/game"
	"github.com/omar0ali/pong/utils"
)

func main() {
	game := &game.Starter{
		Width:  320,
		Height: 240,
		Scale:  2,
		Entities: []game.GameObject{
			&entities.Ball{
				ObjectBase: entities.ObjectBase{
					Position:     entities.Point{X: 30, Y: 30, VX: 3, VY: 6},
					CurrentFrame: utils.LoadImage("ball/ball.png"),
				},
			},
		},
	}
	ebiten.SetWindowSize(game.GetWindowSize())
	ebiten.SetWindowTitle("Pong Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
