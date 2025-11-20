package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/omar0ali/pong/entities"
	"github.com/omar0ali/pong/entities/effects"
	"github.com/omar0ali/pong/game"
	"github.com/omar0ali/pong/utils"
)

func main() {
	windowSize := game.WindowSize{
		Width:  875,
		Height: 480,
	}

	game := &game.Starter{
		WindowSize: windowSize,
		Scale:      2,
		Entities: []game.GameObject{
			effects.NewEffectManager(), // animation
			&entities.Ball{
				ObjectBase: entities.ObjectBase{
					Position: entities.Point{
						X:  float64(windowSize.Width) / 2,
						Y:  float64(windowSize.Height) / 2,
						VX: 5,
						VY: 0,
					},
					CurrentFrame: utils.LoadImage("ball/0.png"),
				},
				MaxSpeed: 5,
			},
			&entities.PaddleBase{
				ObjectBase: entities.ObjectBase{
					Position:     entities.Point{X: 0, Y: 0},
					CurrentFrame: utils.LoadImage("paddles/0.png"),
				},
				Behavior: entities.PlayerBehavior{},
			},
			&entities.PaddleBase{
				ObjectBase: entities.ObjectBase{
					Position:     entities.Point{X: 1, Y: 1},
					CurrentFrame: utils.LoadImage("paddles/1.png"),
				},
				Behavior: entities.CPUBehavior{
					Speed: 3.8,
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
