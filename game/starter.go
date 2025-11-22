package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type WindowSize struct {
	Width, Height int
}

type Starter struct {
	Scale    int
	Entities []GameObject
	Ctx      *GameContext
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
	for _, e := range g.Entities {
		e.Update(g.Ctx)
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
	return g.Ctx.WindowSize.Width, g.Ctx.WindowSize.Height
}

func (g *Starter) GetWindowSize() (int, int) {
	return g.Ctx.WindowSize.Width * g.Scale, g.Ctx.WindowSize.Height * g.Scale
}
