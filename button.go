package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

type Button struct {
	s             *BasicSprite
	ImageSurfaces []*ebiten.Image

	imgIndexes    []int
	imgAltIndexes []int
	index         int
	menuType      string
	mainActive    bool
}

func CreateButton(anchor image.Point, offset image.Point, buttonSize int, surfaces []*ebiten.Image, indices []int) Button {
	pos := anchor.Add(offset)
	img := ebiten.NewImage(buttonSize, buttonSize)
	spr := NewBasicSprite(img, pos)
	return Button{
		s:             spr,
		index:         0,
		imgIndexes:    indices,
		imgAltIndexes: nil,
		mainActive:    true,
		ImageSurfaces: surfaces,
	}
}

func (b *Button) MoveIndex() {
	b.index += 1
	if b.index >= len(b.imgIndexes) {
		b.index = 0
	}
}

func (b *Button) Click() int {
	if b.mainActive {
		return b.imgIndexes[b.index]
	}
	return b.imgAltIndexes[b.index]
}

func (b *Button) Update() {
	b.s.Img.Fill(color.RGBA{R: 10, G: 0, B: 0, A: 190})

	var menuIcon *BasicSprite

	if b.mainActive {
		menuIcon = NewBasicSprite(b.ImageSurfaces[b.imgIndexes[b.index]], image.Point{})
	} else {
		menuIcon = NewBasicSprite(b.ImageSurfaces[b.imgAltIndexes[b.index]], image.Point{})

	}
	menuIcon.DrawCentered(b.s.Img)

}
