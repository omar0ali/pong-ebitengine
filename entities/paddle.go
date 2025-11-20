package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/omar0ali/pong/game"
)

type Behavior interface {
	Update(p *PaddleBase, gc *game.GameContext)
	OnCollision(p *PaddleBase, a game.Collidable)
	GetType() string
}

type PaddleBase struct {
	ObjectBase
	Behavior Behavior
}

func (p *PaddleBase) Update(gc *game.GameContext) {
	p.Behavior.Update(p, gc)
}

func (p *PaddleBase) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(p.GetPosition())
	screen.DrawImage(p.CurrentFrame, opts)
}

func (p *PaddleBase) OnCollision(collision game.Collidable) {
	p.Behavior.OnCollision(p, collision)
}

func (p *PaddleBase) GetType() string {
	return p.Behavior.GetType()
}

// ------------------------------ Player ------------------------------

type PlayerBehavior struct{}

func (PlayerBehavior) Update(p *PaddleBase, gc *game.GameContext) {
	// left side of the screen.
	// w, h := p.CurrentFrame.Size()
	h := p.CurrentFrame.Bounds().Size().Y
	_, my := ebiten.CursorPosition()

	y := float64(my) - float64(h)/2

	if y < 0 {
		y = 0
	}
	if y > float64(gc.Height-h) {
		y = float64(gc.Height - h)
	}

	p.SetPosition(0, y)

}

func (PlayerBehavior) OnCollision(p *PaddleBase, a game.Collidable) {
	ball, ok := a.(*Ball)
	if !ok {
		return
	}

	_, ph := p.GetSize()
	_, bh := ball.GetSize()

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
}

func (PlayerBehavior) GetType() string {
	return "player"
}

// --------------------------------Comp---------------------------------

type CPUBehavior struct {
	Speed float64
}

func (c CPUBehavior) Update(p *PaddleBase, gc *game.GameContext) {
	ballEntity := gc.Starter.GetEntity("ball")
	ball, ok := ballEntity.(*Ball)
	if !ok {
		return
	}

	ph := p.CurrentFrame.Bounds().Size().Y
	paddleCenter := p.Position.Y + float64(ph)/2

	bh := ball.CurrentFrame.Bounds().Size().Y
	ballCenter := ball.Position.Y + float64(bh)/2

	// set paddle position
	p.Position.X = float64(gc.Width) - float64(p.CurrentFrame.Bounds().Size().X)

	// How far is the ball center from paddle center
	dh := ballCenter - paddleCenter

	// limit the paddle to not go over the screen
	if p.Position.Y < 5 {
		p.Position.Y = 5
	}
	if p.Position.Y > float64(gc.Height-ph)-5 {
		p.Position.Y = float64(gc.Height-ph) - 5
	}

	if dh > float64(c.Speed) {
		p.Position.Y += c.Speed
	} else if dh < -c.Speed {
		p.Position.Y -= c.Speed
	} else {
		p.Position.Y += dh
	}
}

func (CPUBehavior) OnCollision(p *PaddleBase, a game.Collidable) {
	ball, ok := a.(*Ball)
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
