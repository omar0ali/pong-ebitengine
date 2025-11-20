package paddle

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/omar0ali/pong/entities"
	"github.com/omar0ali/pong/game"
)

type Behavior interface {
	Update(p *PaddleBase, gc *game.GameContext)
	OnCollision(p *PaddleBase, a game.Collidable)
	GetType() string
}

type PaddleBase struct {
	entities.ObjectBase
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
