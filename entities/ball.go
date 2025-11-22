package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/omar0ali/pong/game"
	"github.com/omar0ali/pong/utils"
)

type Ball struct {
	ObjectBase
	MaxSpeed float64
}

func NewBall(gc *game.GameContext, maxSpeed float64) *Ball {

	ball := &Ball{
		ObjectBase: ObjectBase{
			Position: Point{
				X:  float64(gc.WindowSize.Width / 2),
				Y:  float64(gc.WindowSize.Height / 2),
				VX: float64(maxSpeed),
				VY: 0,
			},
			CurrentFrame: utils.LoadImage("ball/0.png"),
		},
		MaxSpeed: maxSpeed,
	}
	return ball
}

func (b *Ball) Update(gc *game.GameContext) {
	// get size of the screen
	sw, sh := float64(gc.WindowSize.Width), float64(gc.WindowSize.Height)

	// get size of the ball
	bw, bh := b.GetSize()

	// move
	b.Position.X -= b.Position.VX
	b.Position.Y -= b.Position.VY

	if b.Position.Y > sh-float64(bh) {
		gc.EventBus.Publish(game.BallSplashAniamtionBottom)
		b.Position.Y = sh - float64(bh)
		b.Position.VY = -b.Position.VY
	} else if b.Position.Y < 0 {
		gc.EventBus.Publish(game.BallSplashAnimationTop)
		b.Position.Y = 0
		b.Position.VY = -b.Position.VY
	}

	reset := func() {
		//TODO: reset the ball speed and position, also set a max speed so it does not glitch out.
		b.Position.VX = b.MaxSpeed
		b.Position.VY = 0
		b.Position.Y = sh / 2
		b.Position.X = sw / 2

	}

	// reset the ball in the middle
	if b.Position.X > sw-float64(bw) {
		gc.EventBus.Publish(game.ScoreCPU)
		reset()
	}
	if b.Position.X < -float64(bw) {
		gc.EventBus.Publish(game.ScorePlayer)
		reset()
	}
}

func (b *Ball) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(b.GetPosition())
	screen.DrawImage(b.CurrentFrame, opts)
}

func (b *Ball) OnCollision(a game.Collidable, gc *game.GameContext) {
	// ------- Increasing the speed after each hit by the paddle
	if b.Position.VX > 0 {
		b.Position.VX += 0.3
	}
	if b.Position.VY > 0 {
		b.Position.VY += 0.3
	}
	// --------------
}

func (b *Ball) GetType() string {
	return "ball"
}
