package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	windowWidth = 600
	windowHeight = 800
)

// world represents global object that holds the state of the game.
type world struct {
	player *player
}

func newWorld() (*world, error) {
	w := &world{}
	var err error
	w.player, err = newPlayer()
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (w *world) eventLoop() func(screen *ebiten.Image) error {
	return func(screen *ebiten.Image) error {
		screen.Set(windowWidth, windowHeight, color.Black)

		// Draw player
		w.player.draw(screen)

		// Handle player's control
		w.player.handleControl()

		return nil
	}
}

func main() {
	w, err := newWorld()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetRunnableInBackground(true)
	ebiten.Run(w.eventLoop(), windowWidth, windowHeight, 1, "Space Invaders")
}