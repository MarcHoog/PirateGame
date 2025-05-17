package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type CanvasTile struct {
	sprite *BasicSprite

	hasTerrain        bool
	terrainNeighbours []int

	hasWater   bool
	waterOnTop bool

	coin  int
	enemy int
	// objects []*interface{}
}

func (ct *CanvasTile) RefreshTileSprite(de DataEntry2, am *AssetManager) {
	ct.sprite.Img.Clear()
	if ct.hasTerrain {
		a, err := am.Get("graphics/terrain/land/A")
		if err == nil {
			ct.sprite.Img.DrawImage(a.Image, nil)
		} else {
			fmt.Printf("error getting asset: %v\n", err)
		}
	}
}

func (ct *CanvasTile) AddIndex(de DataEntry2, selectionIndex int) {

	switch de.Style {
	case "terrain":
		ct.hasTerrain = true
	case "water":
		ct.hasWater = true
	case "coin":
		ct.coin = selectionIndex
	case "enemy":
		ct.enemy = selectionIndex
	}
}

type Canvas struct {
	editorData   map[int]DataEntry2
	assetManager *AssetManager
	canvas       map[image.Point]*CanvasTile
}

func NewCanvas(am *AssetManager) *Canvas {
	editorData, err := NewData2("./data/editor_data.json")
	if err != nil {
		panic(err)
	}

	return &Canvas{
		editorData:   editorData,
		assetManager: am,
		canvas:       make(map[image.Point]*CanvasTile),
	}

}

func (c *Canvas) Draw(origin image.Point, screen *ebiten.Image) {
	for pos, tile := range c.canvas {
		tile.sprite.pos = origin.Add(pos.Mul(tileSize))
		tile.sprite.Draw(screen)
	}
}

func (c *Canvas) NewTile(pos image.Point, selectionIndex int) error {

	dataEntry, ok := c.editorData[selectionIndex]
	if !ok {
		return fmt.Errorf("data entry not found for index: %d", selectionIndex)
	}

	if existingTile, ok := c.canvas[pos]; ok {
		existingTile.AddIndex(dataEntry, selectionIndex)
		existingTile.RefreshTileSprite(dataEntry, c.assetManager)

	} else {
		blankImage := ebiten.NewImage(tileSize, tileSize)
		tile := &CanvasTile{
			sprite:            NewBasicSprite(blankImage, pos),
			hasTerrain:        false,
			terrainNeighbours: []int{},
			hasWater:          false,
			waterOnTop:        false,
		}

		tile.AddIndex(dataEntry, selectionIndex)
		tile.RefreshTileSprite(dataEntry, c.assetManager)
		c.canvas[pos] = tile
	}

	fmt.Printf("%v\n", c.canvas)
	return nil
}
