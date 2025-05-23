package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"strings"
)

type CanvasTile struct {
	sprite *BasicSprite

	hasTerrain        bool
	terrainNeighbours []string

	hasWater   bool
	waterOnTop bool

	coin  int
	enemy int
	// objects []*interface{}
}

var DirectionAnnotations = []string{
	"A",
	"B",
	"C",
	"D",
	"E",
	"F",
	"G",
	"H",
}

var NeighbourDirections = map[image.Point]int{
	{0, -1}:  0,
	{1, -1}:  1,
	{1, 0}:   2,
	{1, 1}:   3,
	{0, 1}:   4,
	{-1, 1}:  5,
	{-1, 0}:  6,
	{-1, -1}: 7,
}

var TopBottomOffsets = []image.Point{
	{0, 1},
	{0, -1},
}

func (ct *CanvasTile) RefreshTileSprite(am *AssetManager) {
	ct.sprite.Img.Clear()

	if ct.hasWater {

		var a *Asset
		var err error

		if ct.waterOnTop {
			a, err = am.Get("graphics/terrain/water/animation/0")
			if err != nil {
				fmt.Printf("error getting water asset: %v\n", err)
			}

		} else {
			a, err = am.Get("graphics/terrain/water/water_bottom")
			if err != nil {
				fmt.Printf("error getting water asset: %v\n", err)
			}

		}

		ct.sprite.Img.DrawImage(a.Image, nil)
	}

	if ct.hasTerrain {

		fileName := strings.Join(ct.terrainNeighbours, "")
		a, err := am.Get(fmt.Sprintf("graphics/terrain/land/%s", fileName))
		if err != nil {
			a, err = am.Get(fmt.Sprintf("graphics/terrain/land/X"))
			if err != nil {
				fmt.Printf("error getting terrain asset: %v\n", err)
			}

		}
		ct.sprite.Img.DrawImage(a.Image, nil)
	}
	if ct.coin != 0 {
		a, err := am.Get("graphics/items/gold/0")
		if err == nil {
			ct.sprite.Img.DrawImage(a.Image, nil)
		} else {
			fmt.Printf("error getting terrain asset: %v\n", err)
		}
	}

	if ct.enemy != 0 {
		a, err := am.Get("graphics/enemies/tooth/idle/0")
		if err == nil {
			ct.sprite.Img.DrawImage(a.Image, nil)
		} else {
			fmt.Printf("error getting terrain asset: %v\n", err)
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
func (c *Canvas) CheckTerrainNeighbours(tile *CanvasTile, pos image.Point) {
	tile.terrainNeighbours = make([]string, 8)
	for offset, idx := range NeighbourDirections {
		adjPos := pos.Add(offset)
		if adjTile, ok := c.canvas[adjPos]; ok {
			if adjTile.hasTerrain {
				tile.terrainNeighbours[idx] = DirectionAnnotations[idx]
				oppOffset := offset.Mul(-1)
				if oppIdx, ok := NeighbourDirections[oppOffset]; ok {
					adjTile.terrainNeighbours[oppIdx] = DirectionAnnotations[oppIdx]
					adjTile.RefreshTileSprite(c.assetManager)
				}
			}
		}
	}

}

func (c *Canvas) CheckWaterNeighbours(tile *CanvasTile, pos image.Point) {
	for _, offset := range TopBottomOffsets {
		adjPos := pos.Add(offset)
		adjTile, ok := c.canvas[adjPos]

		isTop := offset.Y == -1
		isBottom := offset.Y == 1

		if ok {
			if adjTile.hasWater && isTop {
				tile.waterOnTop = false
			} else if isBottom {
				tile.waterOnTop = true
				adjTile.waterOnTop = false
				adjTile.RefreshTileSprite(c.assetManager)
			}
		} else if isBottom {
			tile.waterOnTop = true
		}
	}
}

func (c *Canvas) NewTile(pos image.Point, selectionIndex int) error {

	dataEntry, ok := c.editorData[selectionIndex]
	if !ok {
		return fmt.Errorf("data entry not found for index: %d", selectionIndex)
	}

	var tile *CanvasTile
	if tile, ok = c.canvas[pos]; ok {
		tile.AddIndex(dataEntry, selectionIndex)

	} else {
		blankImage := ebiten.NewImage(tileSize, tileSize)
		tile = &CanvasTile{
			sprite:            NewBasicSprite(blankImage, pos),
			hasTerrain:        false,
			terrainNeighbours: make([]string, 8),
			hasWater:          false,
			waterOnTop:        false,
		}

		tile.AddIndex(dataEntry, selectionIndex)
		c.canvas[pos] = tile
	}

	if dataEntry.Style == "terrain" {
		c.CheckTerrainNeighbours(tile, pos)

	} else if dataEntry.Style == "water" {
		c.CheckWaterNeighbours(tile, pos)
	}

	tile.RefreshTileSprite(c.assetManager)
	return nil
}
