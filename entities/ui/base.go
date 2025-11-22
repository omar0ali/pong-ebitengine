package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/omar0ali/pong/game"
)

type UIState interface {
	Update(ui *UIBase, gc *game.GameContext)
	Draw(screen *ebiten.Image)
	GetType() string
}

type UIBase struct {
	CurrentState UIState
}

func (u *UIBase) Update(gc *game.GameContext) {
	// additional action
	u.CurrentState.Update(u, gc)
}

func (u *UIBase) Draw(screen *ebiten.Image) {
	u.CurrentState.Draw(screen)
}

func (u *UIBase) GetType() string {
	return u.CurrentState.GetType()
}

func (u *UIBase) SetUIState(state UIState) {
	u.CurrentState = state
}
