package entities

import "github.com/hajimehoshi/ebiten/v2"

type Point struct {
	X, Y, VX, VY float64
}

type ObjectBase struct {
	Position     Point
	CurrentFrame *ebiten.Image
}

func (o *ObjectBase) GetRect() (x1, y1, x2, y2 int) {
	w, h := o.CurrentFrame.Size()
	x1, y2 = int(o.Position.X), int(o.Position.Y)
	x2, y2 = int(o.Position.X)+w, int(o.Position.Y)+h
	return x1, y1, x2, y2
}

func (o *ObjectBase) GetSize() (width, height int) {
	return o.CurrentFrame.Size()
}

func (o *ObjectBase) GetPosition() (x, y float64) {
	return o.Position.X, o.Position.Y
}

func (o *ObjectBase) SetPosition(x, y float64) {
	o.Position.X = x
	o.Position.Y = y
}
