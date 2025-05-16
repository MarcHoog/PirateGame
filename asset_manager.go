package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Asset struct {
	Image  *ebiten.Image
	OsPath string
}

type AssetManager struct {
	assets map[string]*Asset
}

func NewAssetManager(rootPath string) (*AssetManager, error) {
	assets := make(map[string]*Asset)
	am := &AssetManager{assets: assets}
	err := am.Scan(rootPath)
	if err != nil {
		return nil, err
	}

	return am, nil

}

func (am *AssetManager) loadImage(osPath string) (*ebiten.Image, error) {
	f, err := os.Open(osPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}

func (am *AssetManager) Get(assetPath string) (*Asset, error) {
	if asset, ok := am.assets[assetPath]; ok {
		if asset.Image == nil {
			im, err := am.loadImage(asset.OsPath)
			if err != nil {
				return nil, err
			}
			asset.Image = im
			return asset, nil
		}

		return asset, nil
	}

	return nil, fmt.Errorf("asset %s not found", assetPath)

}

func (am *AssetManager) Scan(root string) (err error) {

	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			assetPath := strings.ReplaceAll(path[0:len(path)-len(filepath.Ext(path))], "\\", "/")
			am.assets[assetPath] = &Asset{Image: nil, OsPath: "." + "/" + path}

		}
		return nil

	})

	return

}
