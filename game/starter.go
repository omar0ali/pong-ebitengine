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

func (g *Starter) handleCollision() {
	for i := 0; i < len(g.Entities); i++ {
		a, ok := g.Entities[i].(Collidable)
		if !ok {
			continue
		}
		for j := i + 1; j < len(g.Entities); j++ {
			b, ok := g.Entities[j].(Collidable)
			if !ok {
				continue
			}
			CheckCollision(a, b)
		}
	}
}

func (g *Starter) GetEntity(s string) GameObject {
	for _, e := range g.Entities {
		if e.GetType() == s {
			return e
		}
	}
	return nil
}

func (g *Starter) Update() error {
	ctx := GameContext{
		Width:   g.Width,
		Height:  g.Height,
		Starter: g,
	}

	for _, e := range g.Entities {
		e.Update(&ctx)
	}

	g.handleCollision()

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
