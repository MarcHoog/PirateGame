package main

import "github.com/hajimehoshi/ebiten/v2"

type ButtonEvent struct {
	ID string
}

type ButtonData struct {
	main []interface{}
	alt  []interface{}
}

type Button struct {
	ID            string
	V2            Vector2
	Width, Height int
	state         *EditorState
	pressed       bool
	image         *ebiten.Image
	buttonData    *ButtonData
	index         int
	active        bool
}

func (b *Button) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.V2.AsFloat64())
	screen.DrawImage(b.image, op)
}

func (b *Button) Update() {}

func (b *Button) IsMouseOver(MouseX, MouseY int) bool {
	x, y := b.V2.AsInt()
	return MouseX > x && MouseX < x+b.Width &&
		MouseY > y && MouseY < y+b.Height
}
