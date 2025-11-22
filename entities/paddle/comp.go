package paddle

import (
	"github.com/omar0ali/pong/entities"
	"github.com/omar0ali/pong/game"
)

type CPUBehavior struct {
	Speed float64
}

func (c CPUBehavior) Update(p *PaddleBase, gc *game.GameContext) {
	ballEntity := gc.Starter.GetEntity("ball")
	ball, ok := ballEntity.(*entities.Ball)
	if !ok {
		return
	}

	ph := p.CurrentFrame.Bounds().Size().Y
	paddleCenter := p.Position.Y + float64(ph)/2

	bh := ball.CurrentFrame.Bounds().Size().Y
	ballCenter := ball.Position.Y + float64(bh)/2

	// set paddle position
	p.Position.X = float64(gc.WindowSize.Width) - float64(p.CurrentFrame.Bounds().Size().X)

	// How far is the ball center from paddle center
	dh := ballCenter - paddleCenter

	// limit the paddle to not go over the screen
	if p.Position.Y < 5 {
		p.Position.Y = 5
	}
	if p.Position.Y > float64(gc.WindowSize.Height-ph)-5 {
		p.Position.Y = float64(gc.WindowSize.Height-ph) - 5
	}

	if dh > float64(c.Speed) {
		p.Position.Y += c.Speed
	} else if dh < -c.Speed {
		p.Position.Y -= c.Speed
	} else {
		p.Position.Y += dh
	}
}

func (CPUBehavior) OnCollision(p *PaddleBase, a game.Collidable, gc *game.GameContext) {
	gc.EventBus.Publish(game.BallSplashAnimationRight)

	ball, ok := a.(*entities.Ball)
	if !ok {
		return
	}

	pw, ph := p.GetSize()
	bw, bh := ball.GetSize()

	// Reverse horizontal direction
	ball.Position.VX = -ball.Position.VX

	ballCenterY := ball.Position.Y + float64(bh)/2
	paddleCenterY := p.Position.Y + float64(ph)/2
	distance := paddleCenterY - ballCenterY
	factor := 0.1

	ball.Position.VY += distance * factor

	if ball.Position.VY < -ball.MaxSpeed {
		ball.Position.VY = -ball.MaxSpeed
	} else if ball.Position.VY > ball.MaxSpeed {
		ball.Position.VY = ball.MaxSpeed
	}

	leftEdge := p.Position.X - float64(pw)/2
	rightEdge := p.Position.X + float64(pw)/2

	ballCenterX := ball.Position.X + float64(bw)/2

	if ballCenterX < p.Position.X {
		ball.Position.X = leftEdge - float64(bw)
	} else {
		ball.Position.X = rightEdge
	}
}

func (CPUBehavior) GetType() string {
	return "cpu"
}
