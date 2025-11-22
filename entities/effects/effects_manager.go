package effects

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/omar0ali/pong/entities"
	"github.com/omar0ali/pong/game"
	"github.com/omar0ali/pong/utils"
)

type EffectManager struct {
	Active   []*EffectInstance
	lastTime time.Time
}

func NewEffectManager(gc *game.GameContext) *EffectManager {

	effectManager := &EffectManager{
		Active:   []*EffectInstance{},
		lastTime: time.Now(),
	}

	gc.EventBus.Subscribe(game.BallSplashAniamtionBottom, func() {
		ball, ok := gc.Starter.GetEntity("ball").(*entities.Ball)
		if ok {
			effectManager.Spawn(
				ball.Position.X,
				ball.Position.Y,
				utils.LoadImagesFromFolder("ball/animation/splash_bottom"),
				0.1,
			)
		}
	})

	gc.EventBus.Subscribe(game.BallSplashAnimationTop, func() {
		ball, ok := gc.Starter.GetEntity("ball").(*entities.Ball)
		if ok {
			effectManager.Spawn(
				ball.Position.X,
				ball.Position.Y,
				utils.LoadImagesFromFolder("ball/animation/splash_top"),
				0.1,
			)
		}
	})

	gc.EventBus.Subscribe(game.BallSplashAnimationLeft, func() {
		ball, ok := gc.Starter.GetEntity("ball").(*entities.Ball)
		if ok {
			effectManager.Spawn(
				ball.Position.X,
				ball.Position.Y,
				utils.LoadImagesFromFolder("ball/animation/splash_left"),
				0.1,
			)
		}
	})

	gc.EventBus.Subscribe(game.BallSplashAnimationRight, func() {
		ball, ok := gc.Starter.GetEntity("ball").(*entities.Ball)
		if ok {
			effectManager.Spawn(
				ball.Position.X,
				ball.Position.Y,
				utils.LoadImagesFromFolder("ball/animation/splash_right"),
				0.1,
			)
		}
	})

	return effectManager
}

func (em *EffectManager) Spawn(x, y float64, frames []*ebiten.Image, frameTime float64) {
	em.Active = append(em.Active, &EffectInstance{
		Position: entities.Point{
			X: x,
			Y: y,
		},
		Frames:    frames,
		FrameTime: frameTime,
		Active:    true,
	})
}
func (em *EffectManager) Update(gc *game.GameContext) {
	now := time.Now()
	delta := now.Sub(em.lastTime).Seconds()
	em.lastTime = now

	alive := em.Active[:0] //clear
	for _, e := range em.Active {
		e.Update(delta)
		if e.Active {
			alive = append(alive, e)
		}
	}
	em.Active = alive
}

func (em *EffectManager) Draw(screen *ebiten.Image) {
	for _, e := range em.Active {
		e.Draw(screen)
	}
}

func (em *EffectManager) GetType() string {
	return "effects"
}
