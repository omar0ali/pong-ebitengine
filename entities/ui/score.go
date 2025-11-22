package ui

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/omar0ali/pong/game"
	"golang.org/x/image/font"
)

type ScoreUI struct {
	Player      int
	CPU         int
	Font        font.Face
	Color       color.Color
	screenWidth int
}

func NewScoreUI(font font.Face, gc *game.GameContext) *ScoreUI {
	ui := &ScoreUI{
		Player: 0,
		CPU:    0,
		Font:   font,
		Color:  color.White,
	}

	gc.EventBus.Subscribe(game.ScoreCPU, func() {
		ui.CPU++
	})

	gc.EventBus.Subscribe(game.ScorePlayer, func() {
		ui.Player++
	})

	return ui
}

func (s *ScoreUI) Draw(screen *ebiten.Image) {
	txtScore := fmt.Sprintf("%d - %d", s.Player, s.CPU)

	// Measure text width in pixels
	bounds := text.BoundString(s.Font, txtScore)
	textWidth := bounds.Size().X

	// Center position in pixels
	screenCenter := s.screenWidth / 2
	x := screenCenter - textWidth/2

	// Draw the text
	text.Draw(screen, txtScore, s.Font, x, 20, s.Color)
}

func (s *ScoreUI) Update(gc *game.GameContext) {
	s.screenWidth = gc.WindowSize.Width
}

func (s *ScoreUI) GetType() string {
	return "score"
}
