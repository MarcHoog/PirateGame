package main

type CanvasTile struct {
	hasTerrain        bool
	terrainNeighbours []int

	hasWater   bool
	waterOnTop bool
}
