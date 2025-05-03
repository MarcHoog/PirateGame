package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type EbitenSprite interface {
	Draw(screen *ebiten.Image)
	Update()
}

// Sprite
// TODO: Make sprites work with Point and not with  Vector
type Sprite struct {
	Vector        Vector2
	Width, Height int
	Image         *ebiten.Image
}

func NewSprite(image *ebiten.Image, width, height int, vector Vector2) *Sprite {
	return &Sprite{Image: image}
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.Vector.AsFloat64())
	screen.DrawImage(s.Image, op)
}

func (s *Sprite) Update() {}
