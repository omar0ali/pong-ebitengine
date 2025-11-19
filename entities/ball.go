package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/omar0ali/pong/game"
)

type Ball struct {
	ObjectBase
	MaxSpeed float64
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
		b.Position.X = 0
		b.Position.VX = -b.Position.VX
	}

	// reset the ball in the middle
	if b.Position.Y > sh-float64(bh) || b.Position.Y < -float64(bh) {
		//TODO: reset the ball speed and position, also set a max speed so it does not glitch out.
		b.Position.VX = 0
		b.Position.VY = b.MaxSpeed
		b.Position.Y = sh / 2
		b.Position.X = sw / 2
	}
}

func (b *Ball) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(b.GetPosition())
	screen.DrawImage(b.CurrentFrame, opts)
}

func (b *Ball) OnCollision(a game.Collidable) {
	// ------- Increasing the speed after each hit by the paddle
	if b.Position.VX > 0 {
		b.Position.VX += 0.2
	}
	if b.Position.VY > 0 {
		b.Position.VY += 0.2
	}
	// --------------
}

func (b *Ball) GetType() string {
	return "ball"
}
