package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Starter struct {
	LayoutWidth, LayoutHeight int
}

func (g *Starter) Update() error {
	return nil
}

func (g *Starter) Draw(screen *ebiten.Image) {
}

func (g *Starter) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.LayoutWidth, g.LayoutHeight
}

func (g *Starter) GetWindowSize() (int, int) {
	return g.LayoutWidth * 2, g.LayoutHeight * 2
}
