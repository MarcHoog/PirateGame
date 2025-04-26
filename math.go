package main

import "fmt"

type Vector2 struct {
	X, Y float32
}

func PrintVector(v Vector2) {
	fmt.Printf("%.2f,%.2f\n", v.X, v.Y)
}

func NewVector2FromInt(x, y int) Vector2 {
	return Vector2{float32(x), float32(y)}

}

func (v Vector2) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v Vector2) Scale(factor float32) Vector2 {
	return Vector2{
		X: v.X * factor,
		Y: v.Y * factor,
	}
}

func (v Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func (v Vector2) Sub(v2 Vector2) Vector2 {
	return Vector2{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}

func (v Vector2) asFloat64() (float64, float64) {
	return float64(v.X), float64(v.Y)
}
