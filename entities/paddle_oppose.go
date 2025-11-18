package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/omar0ali/pong/game"
)

type PaddleOppose struct {
	ObjectBase
	Speed float64
}

func (p *PaddleOppose) Update(gc *game.GameContext) {
	ballEntity := gc.Starter.GetEntity("ball")
	ball, ok := ballEntity.(*Ball)
	if !ok {
		return
	}

	pw, _ := p.CurrentFrame.Size()
	paddleCenter := p.Position.X + float64(pw)/2

	bw, _ := ball.CurrentFrame.Size()
	ballCenter := ball.Position.X + float64(bw)/2

	// How far is the ball center from paddle center
	dx := ballCenter - paddleCenter

	if dx > float64(p.Speed) {
		p.Position.X += p.Speed
	} else if dx < -p.Speed {
		p.Position.X -= p.Speed
	} else {
		p.Position.X += dx
	}
}
func (p *PaddleOppose) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(p.GetPosition())
	screen.DrawImage(p.CurrentFrame, opts)
}

func (p *PaddleOppose) OnCollision(a game.Collidable) {
	ball, ok := a.(*Ball)
	if !ok {
		return
	}

	pw, _ := p.GetSize()
	bw, _ := ball.GetSize()

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

}

func (p *PaddleOppose) GetType() string {
	return "paddle_oppose"
}
