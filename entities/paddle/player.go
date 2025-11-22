package paddle

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/omar0ali/pong/entities"
	"github.com/omar0ali/pong/game"
)

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
	if y > float64(gc.WindowSize.Height-h) {
		y = float64(gc.WindowSize.Height - h)
	}

	p.SetPosition(0, y)

}

func (PlayerBehavior) OnCollision(p *PaddleBase, a game.Collidable) {
	ball, ok := a.(*entities.Ball)
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
