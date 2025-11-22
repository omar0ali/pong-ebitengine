package effects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/omar0ali/pong/entities"
)

type EffectInstance struct {
	Position  entities.Point
	Frames    []*ebiten.Image
	FrameIdx  int
	Active    bool
	FrameTime float64
	Timer     float64
}

func (e *EffectInstance) Update(delta float64) {
	if !e.Active {
		return
	}
	e.Timer += delta
	if e.Timer >= e.FrameTime {
		e.Timer -= e.FrameTime
		e.FrameIdx++

		if e.FrameIdx >= len(e.Frames) {
			e.Active = false
		}
	}
}

func (e *EffectInstance) Draw(screen *ebiten.Image) {
	if !e.Active {
		return
	}
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(e.Position.X, e.Position.Y)
	screen.DrawImage(e.Frames[e.FrameIdx], opts)
}
