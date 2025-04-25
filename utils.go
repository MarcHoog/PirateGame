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
