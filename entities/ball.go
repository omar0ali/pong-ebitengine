package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/omar0ali/pong/game"
)

type Ball struct {
	ObjectBase
}

func (b *Ball) Update(gc *game.GameContext) {
	// get size of the screen
	sw, sh := float64(gc.Width), float64(gc.Height)
	// get size of the ball
	bw, bh := b.GetSize()

	// move
	b.Position.X += b.Position.VX
	b.Position.Y += b.Position.VY

	// check
	if b.Position.X > sw-float64(bw) {
		b.Position.X = sw - float64(bw)
		b.Position.VX = -b.Position.VX
	} else if b.Position.X < 0 {
		b.Position.VX = -b.Position.VX
	}

	if b.Position.Y > sh-float64(bh) {
		b.Position.Y = sh - float64(bh)
		b.Position.VY = -b.Position.VY
	} else if b.Position.Y < 0 {
		b.Position.VY = -b.Position.VY
	}
}

func (b *Ball) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(b.GetPosition())
	screen.DrawImage(b.CurrentFrame, opts)
}

func (b *Ball) OnCollision(a game.Collidable) {
	// collides with paddles
}
