package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"time"
)

type EditorScene struct {
	level *EditorLevel
	menu  *EditorMenu
}

func NewEditorScene(am *AssetManager, bus *EventBus) *EditorScene {

	selectionIndex := 2

	err := bus.NewTopic("EditorMouseInput")
	if err != nil {
		panic(err)
	}

	mouseInputLevel, err := bus.Subscribe("EditorMouseInput")
	if err != nil {
		panic(err)
	}

	level := &EditorLevel{
		origin:            image.Point{0, 0},
		panActive:         false,
		panOffset:         image.Point{0, 0},
		supportLineScreen: ebiten.NewImage(screenWidth, screenHeight),
		selectionIndex:    &selectionIndex,
		mouseChannel:      mouseInputLevel,

		canvas: NewCanvas(am),
	}

	mouseInputMenu, err := bus.Subscribe(mouseTopic)
	if err != nil {
		panic(err)
	}

	menu := &EditorMenu{
		selectionIndex:      &selectionIndex,
		selectionIndexTimer: time.After(0),
		eventBus:            bus,
		mouseChannel:        mouseInputMenu,
	}

	AddButtonsToMenu(menu, am)
	AddHighLightSpriteToMenu(menu)

	// Quick Fix so we don't need to have the Whole type menu logic
	menu.highLightSprite.pos = menu.buttons[0].s.pos

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
