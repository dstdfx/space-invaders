package main

import (
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	playerBulletSize = 16
	playerBulletSpeed = 15
)

// playerBullet represents player's bullet entity.
type playerBullet struct {
	x float64
	y float64
	image *ebiten.Image
	angle float64
	isActive bool
}

func newPlayerBullet() (*playerBullet, error) {
	bulletImg, _, err := ebitenutil.NewImageFromFile("sprites/player_bullet.png", ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}

	bullet := &playerBullet{
		x: windowWidth / 2,
		y: windowHeight - playerBulletSize,
		image: bulletImg,
	}

	return bullet, nil
}

func (b *playerBullet) draw(dst *ebiten.Image) {
	if !b.isActive {
		return
	}

	w, h := b.image.Size()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Translate(b.x, b.y)
	_ = dst.DrawImage(b.image, op)
}

func (b *playerBullet) update(){
	b.x += playerBulletSpeed * math.Cos(b.angle)
	b.y += playerBulletSpeed * math.Sin(b.angle)

	if b.x > windowWidth || b.x < 0 || b.y > windowHeight || b.y < 0 {
		b.isActive = false
	}
}