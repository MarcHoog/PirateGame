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
		pixelPos := origin.Add(pos.Mul(tileSize))
		if tile.hasTerrain {
			a, err := c.assetManager.Get("graphics/terrain/land/A")
			if err == nil {
				sprite := NewBasicSprite(a.Image, pixelPos)
				sprite.Draw(screen)
			} else {
				fmt.Printf("error getting asset: %v\n", err)
			}

		}
	}
}

func (c *Canvas) NewTile(pos image.Point, selectionIndex int) error {

	if existingTile, ok := c.canvas[pos]; ok {
		err := c.AddIndex(existingTile, selectionIndex)
		if err != nil {
			fmt.Printf("error updating tile: %v\n", err)
		}

	} else {
		tile := &CanvasTile{
			hasTerrain:        false,
			terrainNeighbours: []int{},
			hasWater:          false,
			waterOnTop:        false,
		}

		err := c.AddIndex(tile, selectionIndex)

		if err != nil {
			fmt.Printf("error creating new tile: %v\n", err)
		} else {
			c.canvas[pos] = tile
		}

	}

	fmt.Printf("%v\n", c.canvas)
	return nil
}

func (c *Canvas) RefreshTileSprite(pos image.Point) {
	tile, ok := c.canvas[pos]
	if !ok {
		return
	}
}

func (c *Canvas) AddIndex(tile *CanvasTile, selectionIndex int) error {

	e, ok := c.editorData[selectionIndex]
	if !ok {
		return fmt.Errorf("no entry found for index: %d", selectionIndex)
	}

	switch e.Style {
	case "terrain":
		tile.hasTerrain = true
	case "water":
		tile.hasWater = true
	case "coin":
		tile.coin = selectionIndex
	case "enemy":
		tile.enemy = selectionIndex
	}

	return nil
}
