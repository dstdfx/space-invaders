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
	basicEnemiesSquad *enemySquad
}

func newWorld() (*world, error) {
	w := &world{}
	var err error

	// Init player
	w.player, err = newPlayer()
	if err != nil {
		return nil, err
	}

	// Init basic enemies squad
	w.basicEnemiesSquad, err = newEnemySquad(5)
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

		// Draw enemies squad
		w.basicEnemiesSquad.draw(screen)

		// Update enemies squad
		w.basicEnemiesSquad.update(w.player.bulletPool)

		return nil
	}
}

func main() {
	w, err := newWorld()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetRunnableInBackground(true)
	err = ebiten.Run(w.eventLoop(), windowWidth, windowHeight, 1, "Space Invaders")
	if err != nil {
		log.Fatal(err)
	}
}
