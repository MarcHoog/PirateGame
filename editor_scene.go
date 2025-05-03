package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

type EditorScene struct {
	level *EditorLevel
	menu  *EditorMenu
}

func NewEditorScene(am *AssetManager) *EditorScene {
	selectionIndex := 2
	level := &EditorLevel{
		origin:            Vector2{0, 0},
		panActive:         false,
		panOffset:         Vector2{0, 0},
		supportLineScreen: ebiten.NewImage(screenWidth, screenHeight),
	}

	menu := &EditorMenu{
		selectionIndex:      &selectionIndex,
		selectionIndexTimer: time.After(0),
		buttons:             nil,
	}

	AddButtonsToMenu(menu, am)
	es := EditorScene{
		level: level,
		menu:  menu,
	}

	return &es
}

func (es *EditorScene) Draw(screen *ebiten.Image) {
	es.level.Draw(screen)
	es.menu.Draw(screen)

}

func (es *EditorScene) Update() (err error) {
	err = es.level.Update()
	err = es.menu.Update()
	return

}
