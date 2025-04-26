package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth    = 1280
	screenHeight   = 720
	tileSize       = 64
	animationSpeed = 6
)

type Game struct {
	editor       *Editor
	editorMenu   *EditorMenu
	assetManager *AssetManager
	cursorAsset  *Asset
}

func NewGame() *Game {
	am, err := NewAssetManager("./graphics")
	if err != nil {
		log.Fatal(fmt.Errorf("error loading new asset manger: %v", err))
	}

	cursor, err := am.Get("graphics/cursors/mouse")
	if err != nil {
		log.Fatal(fmt.Errorf("error getting cursor image: %v", err))
	}

	editor, editorMenu, err := NewEditor(am)
	if err != nil {
		log.Fatal(fmt.Errorf("error creating editor: %v", err))
	}

	g := &Game{
		editor:       editor,
		editorMenu:   editorMenu,
		assetManager: am,
		cursorAsset:  cursor,
	}
	return g
}

func (g *Game) Update() (err error) {
	err = g.editor.Update()
	err = g.editorMenu.Update()

	if err != nil {
		return err
	}
	return
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	g.editor.Draw(screen)
	g.editorMenu.Draw(screen)

	g.DrawCursor(screen)

}

func (g *Game) DrawCursor(screen *ebiten.Image) {
	x, y := ebiten.CursorPosition()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.cursorAsset.Image, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Super Pirate World")
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}

}
