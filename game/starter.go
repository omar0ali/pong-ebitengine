package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Starter struct {
	Width, Height int
	Scale         int
	Entities      []GameObject
}

func (g *Starter) Update() error {
	ctx := GameContext{
		Width:  g.Width,
		Height: g.Height,
		Delta:  1.0 / 60.0,
	}

	for _, e := range g.Entities {
		e.Update(&ctx)
	}
	return nil
}

func (g *Starter) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{111, 111, 111, 255})
	for _, e := range g.Entities {
		e.Draw(screen)
	}
}

func (g *Starter) Layout(_, _ int) (int, int) {
	return g.Width, g.Height
}

func (g *Starter) GetWindowSize() (int, int) {
	return g.Width * g.Scale, g.Height * g.Scale
}
