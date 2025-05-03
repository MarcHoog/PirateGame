package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"time"
)

const (
	selectionCoolDown = 100 * time.Millisecond
	selectionIndexMax = 18
	selectionIndexMin = 2
	buttonSize        = 90
	backGroundMargin  = 6
	backgroundHeight  = buttonSize*2 + 4
	backgroundWidth   = buttonSize*2 + 4
	buttonMargin      = 4
)

// Move a big part of this logic like which components it can select into the
// button itself

func AddButtonsToMenu(em *EditorMenu, am *AssetManager) {
	// 1) load your data & build menuMap
	dm, err := NewData("./data/editor_data.json")
	if err != nil {
		log.Fatalf("loading data map: %v", err)
	}

	// 2) Get all menu Menu Images
	buttonImages := make([]*ebiten.Image, len(dm))
	menuMap := map[string][]int{}
	for _, entry := range dm {
		if entry.Menu == "" {
			continue
		}
		asset, err := am.Get(entry.MenuSurf)
		if err != nil {
			log.Printf("warning: can't load image %q: %v", entry.MenuSurf, err)
			continue
		}
		buttonImages[entry.Index] = asset.Image
		menuMap[entry.Menu] = append(menuMap[entry.Menu], entry.Index)
	}

	// 3) compute the absolute anchor for your whole panel
	anchor := Vector2{
		X: float32(screenWidth - backgroundWidth - backGroundMargin),
		Y: float32(screenHeight - backgroundHeight - backGroundMargin),
	}

	// 4) get your grid offsets
	offsets := gridOffsets(
		2, 2, // 2 rows, 2 cols
		buttonSize,       // each cell is buttonSize
		buttonMargin,     // margin between cells
		backgroundWidth,  // background area width
		backgroundHeight, // background area height
	)

	// 5) Build your buttons
	keys := []string{"terrain", "coin", "enemy", "palm fg"}

	em.buttons = nil
	for i, off := range offsets {
		btn := CreateButton(anchor, off, buttonSize, buttonImages, menuMap[keys[i]])
		if keys[i] == "palm fg" {
			btn.altIndexes = menuMap["palm bg"]
		}
		em.buttons = append(em.buttons, btn)
	}
}

type EditorMenu struct {
	buttons             []Button
	selectionIndex      *int
	selectionIndexTimer <-chan time.Time
	menuMap             map[string][]int
	menuImages          []*ebiten.Image
}

func (em *EditorMenu) Update() error {
	em.updateSelectionIndex()
	for _, button := range em.buttons {
		button.Update()
	}
	return nil
}

func (em *EditorMenu) updateSelectionIndex() {
	newIndex := *em.selectionIndex
	currentIndex := *em.selectionIndex

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		newIndex -= 1
		if newIndex < selectionIndexMin {
			newIndex = selectionIndexMax
		}

	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		newIndex += 1
		if newIndex >= selectionIndexMax {
			newIndex = selectionIndexMin
		}

	}

	select {
	case <-em.selectionIndexTimer:
		if currentIndex != newIndex {
			em.selectionIndex = &newIndex
			fmt.Println("New Selection index: ", newIndex)
		}
		em.selectionIndexTimer = time.After(selectionCoolDown)
	default:
	}

}

func (em *EditorMenu) Draw(screen *ebiten.Image) {
	for _, button := range em.buttons {
		button.Draw(screen)
	}

}
