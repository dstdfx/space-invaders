package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

const (
	rightDirection = iota
	leftDirection
)

// enemySquad represents a common entity that describes a bunch of enemies.
type enemySquad struct {
	basicEnemies []*basicEnemy
	direction int
}

func newEnemySquad(rows int) (*enemySquad, error){
	sq := &enemySquad{direction: rightDirection}
	sq.basicEnemies = make([]*basicEnemy, 0, rows + 7)

	// Init enemies
	for i:=0;i<8; i++ {
		for j:=0;j<rows;j++ {
			x := (float64(i)/10) * windowWidth + basicEnemySize / 2
			y := float64(j) * basicEnemySize + basicEnemySize / 2
			be, err := newBasicEnemy(x, y)
			if err != nil {
				return nil, err
			}
			sq.basicEnemies = append(sq.basicEnemies, be)
		}
	}

	return sq, nil
}

func (es *enemySquad) draw(dst *ebiten.Image) {
	for _, be := range es.basicEnemies {
		be.draw(dst)
	}
}

func (es *enemySquad) update(bullets []*playerBullet) {
	for _, be := range es.basicEnemies {
		// Make a move depend on the current direction
		if es.direction == rightDirection {
			be.x += basicEnemySpeed
		} else if es.direction == leftDirection {
			be.x -= basicEnemySpeed
		}

		enemyCircle := circle{
			x:      be.x,
			y:      be.y,
			radius: be.radius,
		}

		// Check for collision
		for _, b := range bullets {
			bulletCircle := circle{
				x:      b.x,
				y:      b.y,
				radius: b.radius,
			}

			if collides(enemyCircle, bulletCircle) && be.isActive && b.isActive {
				be.isActive = false
				b.isActive = false
				fmt.Println("COLLISION IS FOUND")
			}
		}

		// Update enemies' frames
		be.update()
	}

	// Check if we need to go down
	if es.hasChangedDirection() {
		es.moveEnemiesDown()
	}
}

func (es *enemySquad) hasChangedDirection() bool {
	for _, be := range es.basicEnemies {
		if be.x >= windowWidth-basicEnemySize/2 {
			es.direction = leftDirection
			return true
		} else if be.x <= basicEnemySize/2 {
			es.direction = rightDirection
			return true
		}
	}

	return false
}

func (es *enemySquad) moveEnemiesDown() {
	for _, be := range es.basicEnemies {
		be.y += basicEnemyMoveDownSpeed
	}
}

