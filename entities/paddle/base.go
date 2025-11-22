package paddle

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/omar0ali/pong/entities"
	"github.com/omar0ali/pong/game"
	"github.com/omar0ali/pong/utils"
)

type Behavior interface {
	Update(p *PaddleBase, gc *game.GameContext)
	OnCollision(p *PaddleBase, a game.Collidable, gc *game.GameContext)
	GetType() string
}

func NewPaddle(framePath string, behavior Behavior) *PaddleBase {
	return &PaddleBase{
		ObjectBase: entities.ObjectBase{
			Position:     entities.Point{X: 0, Y: 0},
			CurrentFrame: utils.LoadImage("paddles/" + framePath),
		},
		Behavior: behavior,
	}
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

func (p *PaddleBase) OnCollision(collision game.Collidable, gc *game.GameContext) {
	p.Behavior.OnCollision(p, collision, gc)
}

func (p *PaddleBase) GetType() string {
	return p.Behavior.GetType()
}
