package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

func GetMousePos() image.Point {
	x, y := ebiten.CursorPosition()
	return image.Point{X: x, Y: y}
}

func GetMouseWheel() image.Point {
	xOff, yOff := ebiten.Wheel()
	v := image.Point{
		X: int(xOff),
		Y: int(yOff),
	}
	return v

}

func gridOffsets(rows, cols, cellSize, margin, bgW, bgH int) []image.Point {
	totalW := cols*cellSize + (cols-1)*margin
	totalH := rows*cellSize + (rows-1)*margin
	offsetX := (bgW - totalW) / 2
	offsetY := (bgH - totalH) / 2

	var offs []image.Point
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			x := c*(cellSize+margin) + offsetX
			y := r*(cellSize+margin) + offsetY
			offs = append(offs, image.Point{X: x, Y: y})
		}
	}
	return offs
}

func GetMousePosInGrid(pos image.Point) image.Point {
	if pos.X > 0 {
		pos.X = pos.X / tileSize
	} else {
		pos.X = pos.X/tileSize - 1
	}

	if pos.Y > 0 {
		pos.Y = pos.Y / tileSize
	} else {
		pos.Y = pos.Y/tileSize - 1
	}

	return pos

}
