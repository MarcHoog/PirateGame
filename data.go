package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// TODO: Marc Hoogendoorn Rewrite this for buttons to just use the Normal full structure

type DataEntry struct {
	Index    int
	Style    string `json:"style"`
	Type     string `json:"type"`
	Menu     string `json:"menu"`
	MenuSurf string `json:"menu_surf"`
	Preview  string `json:"preview"`
	Graphics string `json:"graphics"`
}

func NewData(osPath string) ([]DataEntry, error) {
	f, err := os.Open(osPath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	raw := make(map[string]DataEntry)
	err = json.NewDecoder(f).Decode(&raw)
	if err != nil {
		return nil, err
	}

	data := make([]DataEntry, 0)
	for k, v := range raw {

		i, err := strconv.Atoi(k)

		if err != nil {
			return nil, fmt.Errorf("invalid key, key should be convertable to integer: %s", k)
		}
		v.Index = i
		data = append(data, v)
	}
	return data, nil
}

type DataEntry2 struct {
	Style    string `json:"style"`
	Type     string `json:"type"`
	Menu     string `json:"menu"`
	MenuSurf string `json:"menu_surf"`
	Preview  string `json:"preview"`
	Graphics string `json:"graphics"`
}

func NewData2(osPath string) (map[int]DataEntry2, error) {
	f, err := os.Open(osPath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	raw := make(map[string]DataEntry2)
	err = json.NewDecoder(f).Decode(&raw)
	if err != nil {
		return nil, err
	}

	data := make(map[int]DataEntry2, len(raw))
	for k, v := range raw {

		i, err := strconv.Atoi(k)

		if err != nil {
			return nil, fmt.Errorf("invalid key, key should be convertable to integer: %s", k)
		}
		data[i] = v
	}
	return data, nil
}
