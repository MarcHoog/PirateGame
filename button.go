package main

import "github.com/hajimehoshi/ebiten/v2"

type ButtonEventClickEvent struct {
	ButtonID string
}

type Button struct {
	ID                  string
	X, Y, Width, Height int
	image               *ebiten.Image
}

func (b *Button) Draw() {}

func (b *Button) Update() {}

func (b *Button) IsMouseOver(MouseX, MouseY int) bool {
	return MouseX > b.X && MouseX < b.X+b.Width &&
		MouseY > b.Y && MouseY < b.Y+b.Height
}
