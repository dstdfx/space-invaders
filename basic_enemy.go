package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	basicEnemySize = 49
	basicEnemySpeed = 3
	basicEnemyMoveDownSpeed = 3
)

// TODO: generalize enemies to be able to have squads of different enemies

// basicEnemy represents basic enemy entity.
type basicEnemy struct {
	x float64
	y float64
	image *ebiten.Image
}

func newBasicEnemy(x, y float64) (*basicEnemy, error){
	basicEnemyImg, _, err := ebitenutil.NewImageFromFile("sprites/basic_enemy_f0.png", ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}

	basicEnemy := &basicEnemy{
		x: x,
		y: y,
		image: basicEnemyImg,
	}

	return basicEnemy, nil
}

func (be *basicEnemy) draw(dst *ebiten.Image) {
	w, h := be.image.Size()
	op := &ebiten.DrawImageOptions{}
	// Calculate the center of the object
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	//op.GeoM.Rotate(math.Pi)
	op.GeoM.Translate(be.x, be.y)
	_ = dst.DrawImage(be.image, op)
}
