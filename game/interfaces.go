package game

import "github.com/hajimehoshi/ebiten/v2"

type GameObject interface {
	Update()
	Draw(screen *ebiten.Image)
}

type Collidable interface {
	GetRect() (x1, y1, x2, y2 float64)
	OnCollision(a Collidable) // what do you want after it triggers
}
