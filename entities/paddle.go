package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/omar0ali/pong/game"
)

type Paddle struct {
	ObjectBase
}

func (p *Paddle) Update(gc *game.GameContext) {
	w, h := p.CurrentFrame.Size()
	y := float64(gc.Height - h)
	mx, _ := ebiten.CursorPosition()

	x := float64(mx) - float64(w)/2

	if x < 0 {
		x = 0
	}
	if x > float64(gc.Width-w) {
		x = float64(gc.Width - w)
	}

	p.SetPosition(x, y)
}

func (p *Paddle) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(p.GetPosition())
	screen.DrawImage(p.CurrentFrame, opts)
}

func (b *Paddle) OnCollision(a game.Collidable) {
	// collision
}
