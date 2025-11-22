package game

type EventName int

const (
	ScoreCPU EventName = iota
	ScorePlayer
	BallSplashAniamtionBottom
	BallSplashAnimationTop
	BallSplashAnimationLeft
	BallSplashAnimationRight
)

type EventBus struct {
	handlers map[EventName][]func()
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: map[EventName][]func(){},
	}
}

func (b *EventBus) Subscribe(event EventName, handler func()) {
	b.handlers[event] = append(b.handlers[event], handler)
}

func (b *EventBus) Publish(event EventName) {
	for _, handler := range b.handlers[event] {
		handler()
	}
}
