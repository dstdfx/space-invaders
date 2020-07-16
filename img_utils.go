package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func loadImage(filename string) (*ebiten.Image, error) {
	img, _, err := ebitenutil.NewImageFromFile(filename, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	return img, nil
}
