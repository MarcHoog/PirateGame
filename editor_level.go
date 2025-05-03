package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type EditorLevel struct {
	origin            Vector2
	panActive         bool
	panOffset         Vector2
	supportLineScreen *ebiten.Image

	selectionIndex *int
}

func (l *EditorLevel) Update() (err error) {
	l.UpdatePanInput()
	return

}

func (l *EditorLevel) UpdatePanInput() {
	middleMousePressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle)
	if middleMousePressed && !l.panActive {
		l.panActive = true
		l.panOffset = GetMousePos().Sub(l.origin)
	}
	if !middleMousePressed && l.panActive {
		l.panActive = false
	}

	// Left to right panning with Mouse wheel
	middleMouseWheel := GetMouseWheel()
	if !middleMouseWheel.IsZero() {
		if ebiten.IsKeyPressed(ebiten.KeyControlLeft) {
			l.origin.Y -= middleMouseWheel.Y * 25
		} else {
			l.origin.X += middleMouseWheel.Y * 25
		}
	}

	if l.panActive {
		l.origin = GetMousePos().Sub(l.panOffset)
	}

	return
}

func (l *EditorLevel) Draw(screen *ebiten.Image) {

	l.DrawTileLines(screen)
	vector.DrawFilledCircle(screen, l.origin.X, l.origin.Y, 10, color.RGBA{R: 255, B: 0, G: 0, A: 255}, true)

}

func (l *EditorLevel) DrawTileLines(screen *ebiten.Image) {
	l.supportLineScreen.Clear()

	cols := float32(screenWidth / tileSize)
	rows := float32(screenHeight / tileSize)

	offsetVector := Vector2{
		l.origin.X - float32(int(l.origin.X/tileSize)*tileSize),
		l.origin.Y - float32(int(l.origin.Y/tileSize)*tileSize)}

	for i := float32(0); i < cols+1; i++ {
		x := offsetVector.X + i*tileSize
		vector.StrokeLine(l.supportLineScreen, x, 0, x, screenHeight, 2, color.Black, true)
	}

	for i := float32(0); i < rows+1; i++ {
		y := offsetVector.Y + i*tileSize
		vector.StrokeLine(l.supportLineScreen, 0, y, screenWidth, y, 2, color.Black, true)
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	op.ColorScale.ScaleAlpha(0.1)
	screen.DrawImage(l.supportLineScreen, op)
}
