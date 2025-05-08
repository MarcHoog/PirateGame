package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Sprite interface {
	Draw(screen *ebiten.Image)
	DrawCentered(screen *ebiten.Image)
	Position() image.Point
	Size() image.Point
}

type BasicSprite struct {
	Img  *ebiten.Image
	pos  image.Point // top-left in screen coords
	size image.Point // width,height
}

// NewBasicSprite creates a sprite at pos with given image.
func NewBasicSprite(img *ebiten.Image, pos image.Point) *BasicSprite {
	return &BasicSprite{
		Img:  img,
		pos:  pos,
		size: img.Bounds().Size(),
	}
}

func (s *BasicSprite) Position() image.Point { return s.pos }
func (s *BasicSprite) Size() image.Point     { return s.size }

// Draw draws the sprite’s image at pos.
func (s *BasicSprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.pos.X), float64(s.pos.Y))
	screen.DrawImage(s.Img, op)
}

func (s *BasicSprite) DrawCentered(screen *ebiten.Image) {
	screenSize := screen.Bounds().Size()
	spriteSize := s.Size()

	x := float64((screenSize.X - spriteSize.X) / 2)
	y := float64((screenSize.Y - spriteSize.Y) / 2)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(s.Img, op)
}

// --- Helper functions operating on the Sprite interface ---

// DrawSprite draws any Sprite at its stored position.
func DrawSprite(s Sprite, screen *ebiten.Image) {
	s.Draw(screen)
}

// IsColliding checks if a Point is colliding with a sprite.
func IsColliding(p image.Point, s Sprite) bool {
	size := s.Size()
	pos := s.Position()
	return p.X > pos.X && p.X < pos.X+size.X && p.Y > pos.Y && p.Y < pos.Y+size.Y
}

// DrawSpriteCentered draws any Sprite centered on its Position().
func DrawSpriteCentered(s Sprite, screen *ebiten.Image) {
	s.DrawCentered(screen)
}

// GetSize returns width,height of any Sprite.
func GetSize(s Sprite) image.Point {
	return s.Size()
}

// GetPosition returns the top‑left of any Sprite.
func GetPosition(s Sprite) image.Point {
	return s.Position()
}
