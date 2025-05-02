package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"strconv"
	"time"
)

const (
	SelectionCoolDown = 100 * time.Millisecond
	ButtonSize        = 80
	ButtonCols        = 2
	ButtonRows        = 2
	BackGroundMargin  = 6
	BackgroundHeight  = ButtonSize*ButtonRows + 4
	BackgroundWidth   = ButtonSize*ButtonCols + 4
	ButtonMargin      = 6
)

type EditorMenu struct {
	state               *EditorState
	assetManager        *AssetManager
	selectionIndexTimer <-chan time.Time
	backgroundImage     *ebiten.Image
	buttons             []*Button
	prevMouseDown       bool
}

func (em *EditorMenu) Init() {

	em.backgroundImage = ebiten.NewImage(BackgroundWidth, BackgroundHeight)
	em.backgroundImage.Fill(color.RGBA{R: 255, G: 105, B: 180, A: 255})
	buttonImage := ebiten.NewImage(ButtonSize-ButtonMargin, ButtonSize-ButtonMargin)
	buttonImage.Fill(color.RGBA{R: 255, G: 255, B: 255, A: 255})

	anchor := Vector2{
		X: float32(screenWidth - BackgroundWidth - BackGroundMargin),
		Y: float32(screenHeight - BackgroundHeight - BackGroundMargin)}
	totalButtonWidth := ButtonCols*ButtonSize - ButtonMargin
	totalButtonHeight := ButtonRows*ButtonSize - ButtonMargin

	centerOffset := Vector2{
		X: float32(BackgroundWidth-totalButtonWidth) / ButtonCols,
		Y: float32(BackgroundHeight-totalButtonHeight) / ButtonRows,
	}

	i := 0
	for row := 0; row < ButtonCols; row++ {
		for col := 0; col < ButtonRows; col++ {
			i++

			offset := Vector2{
				X: float32(row)*ButtonSize + centerOffset.X,
				Y: float32(col)*ButtonSize + centerOffset.Y,
			}
			v := anchor.Add(offset)
			em.buttons = append(em.buttons, &Button{
				ID:     strconv.Itoa(i),
				V2:     v,
				Width:  ButtonSize,
				Height: ButtonSize,
				image:  buttonImage,
			})

		}
	}
}

func (em *EditorMenu) Update() error {
	em.updateSelectionIndex()
	em.updateMouseInput()
	return nil
}

func (em *EditorMenu) updateSelectionIndex() {
	newIndex := em.state.selectionIndex
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		newIndex = em.state.selectionIndex - 1
		if newIndex < 2 {
			newIndex = len(em.state.editorData) - 1
		}

	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		newIndex = em.state.selectionIndex + 1
		if newIndex >= len(em.state.editorData) {
			newIndex = 2
		}

	}

	select {
	case <-em.selectionIndexTimer:
		if em.state.selectionIndex != newIndex {
			em.state.selectionIndex = newIndex
			fmt.Printf("Selection index: %d\n", em.state.selectionIndex)
		}
		em.selectionIndexTimer = time.After(SelectionCoolDown)
	default:
	}

}

// TODO: Maybe make a seperate function called
// PressButton that contains most of this logic
// Also there isn't really a system
func (em *EditorMenu) updateMouseInput() {
	mouseX, mouseY := ebiten.CursorPosition()
	mouseDown := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	// Check for click (transition from up to down)
	clicked := mouseDown && !em.prevMouseDown

	for _, button := range em.buttons {
		if clicked && button.IsMouseOver(mouseX, mouseY) {
			fmt.Printf("Button pressed: %s\n", button.ID)
		}
	}

	em.prevMouseDown = mouseDown
}

func (em *EditorMenu) Draw(screen *ebiten.Image) {
	anchor := Vector2{
		X: float32(screenWidth - BackgroundWidth - BackGroundMargin),
		Y: float32(screenHeight - BackgroundHeight - BackGroundMargin)}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(anchor.AsFloat64())
	screen.DrawImage(em.backgroundImage, op)

	for i := 0; i < len(em.buttons); i++ {
		em.buttons[i].Draw(screen)
	}

}
