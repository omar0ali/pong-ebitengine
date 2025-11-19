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

func (PlayerBehavior) OnCollision(p *PaddleBase, a game.Collidable) {
	ball, ok := a.(*Ball)
	if !ok {
		return
	}

	pw, _ := p.GetSize()
	bw, bh := ball.GetSize()

	ball.Position.VY = -ball.Position.VY

	ballCenterX := ball.Position.X + float64(bw)/2
	paddleCenterX := p.Position.X + float64(pw)/2

	distance := ballCenterX - paddleCenterX
	factor := 0.1

	ball.Position.VX += distance * factor

	if ball.Position.VX < -ball.MaxSpeed {
		ball.Position.VX = -ball.MaxSpeed
	} else if ball.Position.VX > ball.MaxSpeed {
		ball.Position.VX = ball.MaxSpeed
	}

	ball.Position.Y = p.Position.Y - float64(bh)
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

	pw, _ := p.CurrentFrame.Size()
	paddleCenter := p.Position.X + float64(pw)/2

	bw, _ := ball.CurrentFrame.Size()
	ballCenter := ball.Position.X + float64(bw)/2

	// How far is the ball center from paddle center
	dx := ballCenter - paddleCenter

	if dx > float64(c.Speed) {
		p.Position.X += c.Speed
	} else if dx < -c.Speed {
		p.Position.X -= c.Speed
	} else {
		p.Position.X += dx
	}

}

func (CPUBehavior) OnCollision(p *PaddleBase, a game.Collidable) {
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

func (CPUBehavior) GetType() string {
	return "cpu"
}
