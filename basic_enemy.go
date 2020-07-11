package main

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	basicEnemySize = 49
	basicEnemySpeed = 3
)

// basicEnemy represents basic enemy entity.
type basicEnemy struct {
	x float64
	y float64
	image *ebiten.Image
	lastMove time.Time
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
		lastMove: time.Now().UTC(),
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


func initBasicEnemiesSquad() ([]*basicEnemy, error){
	basicEnemies := make([]*basicEnemy, 0)

	// Init enemies
	for i:=0;i<8; i++ {
		for j:=0;j<5;j++ {
			x := (float64(i)/8) * windowWidth + basicEnemySize / 2
			y := float64(j) * basicEnemySize + basicEnemySize / 2
			be, err := newBasicEnemy(x, y)
			if err != nil {
				return nil, err
			}
			basicEnemies = append(basicEnemies, be)
			fmt.Println("X, Y=", x, y)
		}
	}

	return basicEnemies, nil
}

