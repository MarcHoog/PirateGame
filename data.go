package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type DataEntry struct {
	Style    string  `json:"style"`
	Type     string  `json:"type"`
	Menu     *string `json:"menu"`      // null → nil
	MenuSurf *string `json:"menu_surf"` // null → nil
	Preview  *string `json:"preview"`   // null → nil
	Graphics *string `json:"graphics"`  // null → nil
}

func NewDataMap(osPath string) (*map[int]DataEntry, error) {
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

	data := make(map[int]DataEntry)
	for k, v := range raw {

		i, err := strconv.Atoi(k)
		if err != nil {
			return nil, fmt.Errorf("invalid key, key should be convertable to integer: %s", k)
		}
		data[i] = v

	}

	return &data, nil
}
