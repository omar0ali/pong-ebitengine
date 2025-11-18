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
		Width:  720,
		Height: 480,
		Scale:  2,
		Entities: []game.GameObject{
			&entities.Ball{
				ObjectBase: entities.ObjectBase{
					Position:     entities.Point{X: 360, Y: 220, VX: 0, VY: 5},
					CurrentFrame: utils.LoadImage("ball/0.png"),
				},
				MaxSpeed: 5,
			},
			&entities.Paddle{
				ObjectBase: entities.ObjectBase{
					Position:     entities.Point{X: 1, Y: 1},
					CurrentFrame: utils.LoadImage("paddles/0.png"),
				},
			},
			&entities.PaddleOppose{
				ObjectBase: entities.ObjectBase{
					Position:     entities.Point{X: 1, Y: 1},
					CurrentFrame: utils.LoadImage("paddles/1.png"),
				},
				Speed: 3.8,
			},
		},
	}
	ebiten.SetWindowSize(game.GetWindowSize())
	ebiten.SetWindowTitle("Pong Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
