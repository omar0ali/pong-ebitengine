package game

import "github.com/hajimehoshi/ebiten/v2"

type GameContext struct {
	WindowSize *WindowSize
	EventBus   *EventBus
	Starter    *Starter
}

type GameObject interface {
	Update(g *GameContext)
	Draw(screen *ebiten.Image)
	GetType() string
}

type Collidable interface {
	GetRect() (x1, y1, x2, y2 float64)
	OnCollision(a Collidable, gc *GameContext) // what do you want after it triggers
}

func CheckCollision(a, b Collidable, gc *GameContext) bool {
	ax1, ay1, ax2, ay2 := a.GetRect()
	bx1, by1, bx2, by2 := b.GetRect()
	if ax2 > bx1 && ax1 < bx2 && ay2 > by1 && ay1 < by2 {
		a.OnCollision(b, gc)
		b.OnCollision(a, gc)
		return true
	}
	return false
}
