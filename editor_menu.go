package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"time"
)

const (
	selectionCoolDown = 100 * time.Millisecond
	ButtonSize        = 80
	ButtonCols        = 2
	ButtonRows        = 2
	BackGroundMargin  = 6
	BackgroundHeight  = ButtonSize * ButtonRows
	BackgroundWidth   = ButtonSize * ButtonCols
	ButtonMargin      = 6
)

type EditorMenu struct {
	state               *EditorState
	assetManager        *AssetManager
	selectionIndexTimer <-chan time.Time
	menuEvents          chan func()
	backgroundImage     *ebiten.Image
	buttonImage         *ebiten.Image
}

func (em *EditorMenu) Init() {

	em.backgroundImage = ebiten.NewImage(BackgroundWidth, BackgroundHeight)
	em.buttonImage = ebiten.NewImage(ButtonSize-ButtonMargin, ButtonSize-ButtonMargin)

}

func (em *EditorMenu) Update() error {
	em.UpdateSelectionIndexInput()
	return nil
}

func (em *EditorMenu) UpdateSelectionIndexInput() {

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
		em.selectionIndexTimer = time.After(selectionCoolDown)
		em.state.selectionIndex = newIndex
		// fmt.Printf("Selection index: %d\n", em.state.selectionIndex)
	default:
		return

	}
}

func (em *EditorMenu) Draw(screen *ebiten.Image) {

	// This is the top left 'Ancor point' of which we will reference all objects
	v := Vector2{
		X: float32(screenWidth - BackgroundWidth - BackGroundMargin),
		Y: float32(screenHeight - BackgroundHeight - BackGroundMargin)}

	em.DrawBackGround(v, screen)
	em.DrawButtons(v, screen)

}

func (em *EditorMenu) DrawButtons(anchor Vector2, screen *ebiten.Image) {
	totalButtonWidth := ButtonCols*ButtonSize - ButtonMargin
	totalButtonHeight := ButtonRows*ButtonSize - ButtonMargin

	centerOffset := Vector2{
		X: float32(BackgroundWidth-totalButtonWidth) / ButtonCols,
		Y: float32(BackgroundHeight-totalButtonHeight) / ButtonRows,
	}

	for row := 0; row < ButtonCols; row++ {
		for col := 0; col < ButtonRows; col++ {
			offset := Vector2{
				X: float32(row)*ButtonSize + centerOffset.X,
				Y: float32(col)*ButtonSize + centerOffset.Y,
			}
			v := anchor.Add(offset)
			em.buttonImage.Fill(color.RGBA{R: 255, G: 255, B: 180, A: 255})
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(v.asFloat64())
			screen.DrawImage(em.buttonImage, op)
		}
	}
}
func (em *EditorMenu) DrawBackGround(v Vector2, screen *ebiten.Image) {

	em.backgroundImage.Fill(color.RGBA{R: 255, G: 105, B: 180, A: 255})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(v.asFloat64())
	screen.DrawImage(em.backgroundImage, op)

}
