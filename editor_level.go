package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	"image/color"
)

type EditorLevel struct {
	origin            image.Point
	panActive         bool
	panOffset         image.Point
	supportLineScreen *ebiten.Image
	mouseChannel      chan Event
	selectionIndex    *int

	lastSelectedCell image.Point

	canvas *Canvas
}

func (l *EditorLevel) Update() (err error) {
	l.UpdatePanInput()
	l.UpdateMouseInput()
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
	if middleMouseWheel.Y != 0 {
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

func (l *EditorLevel) UpdateMouseInput() {
	select {
	case val := <-l.mouseChannel:
		if val.Data["btn"] == ebiten.MouseButtonLeft {
			pos := val.Data["pos"].(image.Point)
			l.canvasAdd(pos)

		}
	default:
	}

}

func (l *EditorLevel) canvasAdd(pos image.Point) {
	pos = GetMousePosInGrid(pos.Sub(l.origin))
	if l.lastSelectedCell != pos {
		// Maybe not ignore this error but for now who really cares am i right?
		_ = l.canvas.NewTile(pos, *l.selectionIndex)

	}
	l.lastSelectedCell = pos
}

func (l *EditorLevel) Draw(screen *ebiten.Image) {

	l.DrawTileLines(screen)
	vector.DrawFilledCircle(screen, float32(l.origin.X), float32(l.origin.Y), 10, color.RGBA{R: 255, B: 0, G: 0, A: 255}, true)
	l.canvas.Draw(l.origin, screen)
}

func (l *EditorLevel) DrawTileLines(screen *ebiten.Image) {
	l.supportLineScreen.Clear()

	cols := screenWidth / tileSize
	rows := screenHeight / tileSize

	offsetVector := image.Point{
		X: l.origin.X - l.origin.X/tileSize*tileSize,
		Y: l.origin.Y - l.origin.Y/tileSize*tileSize}

	for i := 0; i < cols+1; i++ {
		x := offsetVector.X + i*tileSize
		vector.StrokeLine(l.supportLineScreen, float32(x), 0, float32(x), screenHeight, 2, color.Black, true)
	}

	for i := 0; i < rows+1; i++ {
		y := offsetVector.Y + i*tileSize
		vector.StrokeLine(l.supportLineScreen, 0, float32(y), screenWidth, float32(y), 2, color.Black, true)
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	op.ColorScale.ScaleAlpha(0.1)
	screen.DrawImage(l.supportLineScreen, op)
}
