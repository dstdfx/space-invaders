package main

import (
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	playerSize = 87
	playerSpeed = 7
	shootCooldown = time.Millisecond * 250
)

// player represents player entity.
type player struct {
	x float64
	y float64
	image *ebiten.Image
	lastShoot time.Time
	bulletPool []*playerBullet
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

	// init player's bullet pool
	if err := player.initBulletPool(); err != nil {
		return nil, err
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

	// Draw bullets
	for _, b := range p.bulletPool {
		b.draw(dst)
	}
}

func (p *player) handleControl() {
	// Handle bullets updating
	for _, b := range p.bulletPool {
		b.update()
	}

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

	// Shooting control
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if time.Since(p.lastShoot) < shootCooldown {
			return
		}

		// Make sure that bullet comes from the right position
		p.shoot(p.x, p.y - playerBulletSize)
		p.lastShoot = time.Now().UTC()
	}
}

func (p *player) shoot(x, y float64) {
	if b, ok := p.getBulletFromPool(); ok {
		b.isActive = true
		b.x = x
		b.y = y
		b.angle = 270 * (math.Pi / 180)
	} else {
		log.Println("bullet pool is empty")
	}
}

func (p *player) initBulletPool() error {
	for i:=0;i<10;i++{
		pb, err := newPlayerBullet()
		if err != nil {
			return err
		}
		p.bulletPool = append(p.bulletPool, pb)
	}

	return nil
}

func (p *player) getBulletFromPool() (*playerBullet, bool) {
	for _, b := range p.bulletPool {
		if !b.isActive {
			return b, true
		}
	}

	return nil, false
}
