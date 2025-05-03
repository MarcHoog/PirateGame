package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Button struct {
	*Sprite
	ImageSurfaces []*ebiten.Image

	indexes    []int
	altIndexes []int
	index      int
	mainActive bool
}

func CreateButton(anchor Vector2, offset Vector2, buttonSize int, surfaces []*ebiten.Image, indices []int) Button {
	pos := anchor.Add(offset)
	img := ebiten.NewImage(buttonSize, buttonSize)
	spr := &Sprite{
		Vector: pos,
		Width:  buttonSize,
		Height: buttonSize,
		Image:  img,
	}
	return Button{
		Sprite:        spr,
		ImageSurfaces: surfaces,
		indexes:       indices,
		altIndexes:    nil,
		index:         0,
		mainActive:    true,
	}
}

// Contains Returns if the mouse is over the Sprite
// TODO: Should just be it's own function that takes in a ebitSprite
func (b *Button) Contains() bool {
	pos := GetMousePos()
	vx, vy := b.Sprite.Vector.X, b.Sprite.Vector.Y
	w := float32(b.Sprite.Width)
	h := float32(b.Sprite.Height)

	return pos.X > vx &&
		pos.X < vx+w &&
		pos.Y > vy &&
		pos.Y < vy+h
}

func (b *Button) Activate() {
	b.mainActive = true
}

func (b *Button) Update() {
	b.Sprite.Image.Fill(color.RGBA{R: 126, G: 124, B: 124, A: 255})

	surf := b.ImageSurfaces[b.indexes[0]]
	surfSize := surf.Bounds().Size()
	surfW, surfH := float32(surfSize.X), float32(surfSize.Y)

	// Center point of the button
	center := Vector2{float32(b.Width) / 2, float32(b.Height) / 2}

	// Top-left point to draw the image centered
	pos := center.Add(Vector2{-surfW / 2, -surfH / 2})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(pos.AsFloat64())
	b.Sprite.Image.DrawImage(surf, op)
}
