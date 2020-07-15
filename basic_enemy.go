package main

import (
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	basicEnemySize                = 49
	basicEnemyRadius              = 18
	basicEnemySpeed               = 3
	basicEnemyMoveDownSpeed       = 3
	basicEnemyFrameChangeCooldown = 500 * time.Millisecond
)

// TODO: generalize enemies to be able to have squads of different enemies

// basicEnemy represents basic enemy entity.
type basicEnemy struct {
	x                  float64
	y                  float64
	frames             []*ebiten.Image
	currentFrame       int
	lastFrameChangedAt time.Time
	isActive           bool
	radius             float64
}

func newBasicEnemy(x, y float64) (*basicEnemy, error) {
	// TODO: load frames once and reuse them
	frame0, _, err := ebitenutil.NewImageFromFile("sprites/basic_enemy_f0.png", ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}

	frame1, _, err := ebitenutil.NewImageFromFile("sprites/basic_enemy_f1.png", ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}

	basicEnemy := &basicEnemy{
		x:        x,
		y:        y,
		frames:   []*ebiten.Image{frame0, frame1},
		isActive: true,
		radius:   basicEnemyRadius,
	}

	return basicEnemy, nil
}

func (be *basicEnemy) draw(dst *ebiten.Image) {
	if !be.isActive {
		return
	}

	w, h := be.frames[be.currentFrame].Size()
	op := &ebiten.DrawImageOptions{}
	// Calculate the center of the object
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Translate(be.x, be.y)
	_ = dst.DrawImage(be.frames[be.currentFrame], op)
}

func (be *basicEnemy) update() {
	if time.Since(be.lastFrameChangedAt) < basicEnemyFrameChangeCooldown {
		return
	}

	if be.currentFrame == 0 {
		be.currentFrame++
	} else {
		be.currentFrame--
	}

	be.lastFrameChangedAt = time.Now().UTC()
}
