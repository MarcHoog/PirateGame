package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
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

func AddHighLightSpriteToMenu(em *EditorMenu) {

	border := 6
	img := ebiten.NewImage(buttonSize, buttonSize)

	// 1x1 pixel to draw with
	pixel := ebiten.NewImage(1, 1)
	pixel.Fill(color.RGBA{R: 184, G: 134, B: 11, A: 255})

	op := &ebiten.DrawImageOptions{}

	// Top
	op.GeoM.Reset()
	op.GeoM.Scale(float64(buttonSize), float64(border))
	img.DrawImage(pixel, op)

	// Bottom
	op.GeoM.Reset()
	op.GeoM.Scale(float64(buttonSize), float64(border))
	op.GeoM.Translate(0, float64(buttonSize-border))
	img.DrawImage(pixel, op)

	// Left
	op.GeoM.Reset()
	op.GeoM.Scale(float64(border), float64(buttonSize))
	img.DrawImage(pixel, op)

	// Right
	op.GeoM.Reset()
	op.GeoM.Scale(float64(border), float64(buttonSize))
	op.GeoM.Translate(float64(buttonSize-border), 0)

	img.DrawImage(pixel, op)

	s := NewBasicSprite(img, image.Point{})
	em.highLightSprite = s
}

func AddButtonsToMenu(em *EditorMenu, am *AssetManager) {
	dm, err := NewData("./data/editor_data.json")
	if err != nil {
		log.Fatalf("loading data map: %v", err)
	}

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

	anchor := image.Point{
		X: screenWidth - backgroundWidth - backGroundMargin,
		Y: screenHeight - backgroundHeight - backGroundMargin,
	}

	offsets := gridOffsets(
		2, 2,
		buttonSize,
		buttonMargin,
		backgroundWidth,
		backgroundHeight,
	)

	// This is a little bit Custom
	keys := []string{"terrain", "coin", "palm fg", "enemy"}

	em.buttons = nil
	for i, off := range offsets {
		btn := CreateButton(anchor, off, buttonSize, buttonImages, menuMap[keys[i]])
		if keys[i] == "palm fg" {
			btn.imgAltIndexes = menuMap["palm bg"]
		}
		em.buttons = append(em.buttons, &btn)
	}
}

type EditorMenu struct {
	buttons []*Button

	highLightSprite     *BasicSprite
	selectionIndex      *int
	selectionIndexTimer <-chan time.Time

	menuMap    map[string][]int
	menuImages []*ebiten.Image

	eventBus     *EventBus
	mouseChannel chan Event
}

func (em *EditorMenu) Update() error {
	em.updateSelectionIndex()
	em.updateMouseInput()
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
		button.s.Draw(screen)
	}
	em.highLightSprite.Draw(screen)
}

func (em *EditorMenu) updateMouseInput() {
	select {
	case val := <-em.mouseChannel:
		btn := val.Data["btn"].(ebiten.MouseButton)
		pos := val.Data["pos"].(image.Point)
		for _, b := range em.buttons {
			collide := IsColliding(pos, b.s)
			if collide {
				if btn == ebiten.MouseButtonMiddle {
					if b.imgAltIndexes != nil {
						b.mainActive = !b.mainActive
					}
				} else if btn == ebiten.MouseButtonRight {
					b.MoveIndex()
				}

				*em.selectionIndex = b.Click()
				em.highLightSprite.pos = b.s.pos

				return
			}
		}

		// The menu has priority over the clicks if the menu doesn't colide with anything the click goes to the level
		em.eventBus.Publish("EditorMouseInput", val)

	default:
		// Optional: do nothing if no value is available
	}

	return
}
