package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	windowWidth = 600
	windowHeight = 800
)

func eventLoop() func(screen *ebiten.Image) error {
	return func(screen *ebiten.Image) error {

		// TODO: implement me

		ebitenutil.DebugPrintAt(screen, "Hello, World!", 150, 150)

		return nil
	}
}

func main() {
	ebiten.SetRunnableInBackground(true)
	ebiten.Run(eventLoop(), windowWidth, windowHeight, 1, "Space Invaders")
}