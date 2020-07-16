package main

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	basicEnemySize          = 49
	basicEnemyRadius        = 18
	basicEnemySpeed         = 3
	basicEnemyMoveDownSpeed = 3
)

const (
	basicEnemyIdleSequence    = "basic_enemy_idle"
	basicEnemyDestroySequence = "basic_enemy_destroy"
)

// TODO: generalize enemies to be able to have squads of different enemies

// basicEnemy represents basic enemy entity.
type basicEnemy struct {
	x                  float64
	y                  float64
	radius             float64
	sequences          map[string]*sequence
	currentSequence    string
	lastFrameChangedAt time.Time
	animationFinished  bool
	isActive           bool
}

func newBasicEnemy(x, y float64) (*basicEnemy, error) {
	basicEnemy := &basicEnemy{
		x:                  x,
		y:                  y,
		isActive:           true,
		radius:             basicEnemyRadius,
		currentSequence:    basicEnemyIdleSequence,
		lastFrameChangedAt: time.Now().UTC(),
	}

	idleSequence, err := newSequence("sprites/basic_enemy/idle", 2, true)
	if err != nil {
		panic(fmt.Errorf("creating idle sequence: %v", err))
	}
	destroySequence, err := newSequence("sprites/basic_enemy/destroy", 4, false)
	if err != nil {
		panic(fmt.Errorf("creating destroy sequence: %v", err))
	}
	basicEnemy.sequences = map[string]*sequence{
		basicEnemyIdleSequence:    idleSequence,
		basicEnemyDestroySequence: destroySequence,
	}

	return basicEnemy, nil
}

func (be *basicEnemy) draw(dst *ebiten.Image) {
	if !be.isActive && be.currentSequence == basicEnemyDestroySequence && be.animationFinished {
		return
	}

	img := be.sequences[be.currentSequence].image()
	w, h := img.Size()
	op := &ebiten.DrawImageOptions{}
	// Calculate the center of the object
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Translate(be.x, be.y)
	_ = dst.DrawImage(img, op)
}

func (be *basicEnemy) update() {
	sequence := be.sequences[be.currentSequence]

	frameInterval := float64(time.Second) / sequence.sampleRate

	if time.Since(be.lastFrameChangedAt) >= time.Duration(frameInterval) {
		be.animationFinished = sequence.nextFrame()
		be.lastFrameChangedAt = time.Now().UTC()
	}
}
