package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func GetMousePos() Vector2 {
	return NewVector2FromInt(ebiten.CursorPosition())
}

func GetMouseWheel() Vector2 {
	xOff, yOff := ebiten.Wheel()
	v := Vector2{
		X: float32(xOff),
		Y: float32(yOff),
	}
	return v

}

func gridOffsets(rows, cols, cellSize, margin, bgW, bgH int) []Vector2 {
	totalW := cols*cellSize + (cols-1)*margin
	totalH := rows*cellSize + (rows-1)*margin
	offsetX := float32(bgW-totalW) / 2
	offsetY := float32(bgH-totalH) / 2

	var offs []Vector2
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			x := float32(c*(cellSize+margin)) + offsetX
			y := float32(r*(cellSize+margin)) + offsetY
			offs = append(offs, Vector2{X: x, Y: y})
		}
	}
	return offs
}
