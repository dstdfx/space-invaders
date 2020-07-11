package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	playerSize = 87
	playerSpeed = 7
)

// player represents player entity.
type player struct {
	x float64
	y float64
	image *ebiten.Image
}

func newPlayer() (*player, error) {
	playerImg, _, err := ebitenutil.NewImageFromFile("sprites/player.png", ebiten.FilterDefault)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Start point of the player should be on the bottom line
	// of the screen in the center.
	player := &player{
		x: windowWidth / 2,
		y: windowHeight - playerSize,
		image: playerImg,
	}
	return player, nil
}

func (p *player) draw(dst *ebiten.Image) {
	w, h := p.image.Size()
	op := &ebiten.DrawImageOptions{}
	// Calculate the center of the object
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Translate(p.x, p.y)
	_ = dst.DrawImage(p.image, op)
}

func (p *player) handleControl() {
	// Right side control
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		// Respect right border
		if p.x + playerSpeed < windowWidth - playerSize / 2 {
			p.x += playerSpeed
			return
		}
	}

	// Left side control
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		// Respect left border
		if p.x - playerSpeed - playerSize / 2 > 0 {
			p.x -= playerSpeed
			return
		}
	}
}



