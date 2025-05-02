package main

import "github.com/hajimehoshi/ebiten/v2"

type EbitSprite interface {
	Draw(screen *ebiten.Image)
	Update()
}

type Sprite struct {
	Vector        Vector2
	Width, Height int
	image         *ebiten.Image
}

func NewSprite(image *ebiten.Image, width, height int, vector2 Vector2) *Sprite {
	return &Sprite{image: image}
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.Vector.AsFloat64())
	screen.DrawImage(s.image, op)
}

func (s *Sprite) Update() {}

type SpriteGroup struct {
	Sprites []EbitSprite
}

func NewSpriteGroup() *SpriteGroup {
	return &SpriteGroup{Sprites: make([]EbitSprite, 0)}
}

func (s *SpriteGroup) Draw(screen *ebiten.Image) {
	for _, sprite := range s.Sprites {
		sprite.Draw(screen)
	}
}

func (s *SpriteGroup) Update() {
	for _, sprite := range s.Sprites {
		sprite.Update()
	}
}
