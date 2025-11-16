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

func (p *Paddle) OnCollision(a game.Collidable) {
	ball, ok := a.(*Ball)
	if !ok {
		return
	}

	// ------- Increasing the speed after each hit by the paddle
	if ball.Position.VX > 0 {
		ball.Position.VX += 0.2
	}
	if ball.MaxSpeed > 0 {
		ball.MaxSpeed += 0.2
	}
	if ball.Position.VY > 0 {
		ball.Position.VY += 0.2
	}
	// --------------

	pw, _ := p.GetSize()
	bw, bh := ball.GetSize()

	ball.Position.VY = -ball.Position.VY

	ballCenterX := ball.Position.X + float64(bw)/2
	paddleCenterX := p.Position.X + float64(pw)/2

	distance := ballCenterX - paddleCenterX
	factor := 0.1

	ball.Position.VX += distance * factor

	// limit the angle how much it turn to X and -x base on the max speed
	if ball.Position.VX < -ball.MaxSpeed {
		ball.Position.VX = -ball.MaxSpeed
	} else if ball.Position.VX > ball.MaxSpeed {
		ball.Position.VX = ball.MaxSpeed
	}

	ball.Position.Y = p.Position.Y - float64(bh)
}

func (p *Paddle) GetType() string {
	return "paddle"
}
