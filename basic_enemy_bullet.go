package main

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

const (
	basicEnemyBulletSize   = 15
	basicEnemyBulletRadius = 3
	basicEnemyBulletSpeed  = 7
)

// basicEnemy represents basic enemy's bullet entity.
type basicEnemyBullet struct {
	x        float64
	y        float64
	image    *ebiten.Image
	angle    float64
	isActive bool
	radius   float64
}

func newBasicEnemyBullet() (*basicEnemyBullet, error) {
	bulletImg, err := loadImage("sprites/basic_enemy_bullet.png")
	if err != nil {
		return nil, err
	}

	bullet := &basicEnemyBullet{
		x:      windowWidth / 2,
		y:      windowHeight - basicEnemyBulletSize,
		image:  bulletImg,
		radius: basicEnemyBulletRadius,
	}

	return bullet, nil
}

func (b *basicEnemyBullet) draw(dst *ebiten.Image) {
	if !b.isActive {
		return
	}

	w, h := b.image.Size()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Translate(b.x, b.y)
	_ = dst.DrawImage(b.image, op)
}

func (b *basicEnemyBullet) update() {
	b.x -= basicEnemyBulletSpeed * math.Cos(b.angle)
	b.y -= basicEnemyBulletSpeed * math.Sin(b.angle)

	if b.x > windowWidth || b.x < 0 || b.y > windowHeight || b.y < 0 {
		b.isActive = false
	}
}
