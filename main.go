package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/omar0ali/pong/entities"
	"github.com/omar0ali/pong/entities/effects"
	"github.com/omar0ali/pong/entities/paddle"
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
			entities.NewBall(windowSize, 5),
			paddle.NewPaddle("0.png", paddle.PlayerBehavior{}),
			paddle.NewPaddle("1.png", paddle.CPUBehavior{Speed: 3.8}),
			entities.NewScoreUI(utils.LoadFont("fonts/pixel_font_2.ttf", 18)),
		},
	}
	ebiten.SetWindowSize(game.GetWindowSize())
	ebiten.SetWindowTitle("Pong Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
