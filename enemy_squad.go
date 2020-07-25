package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	rightDirection = iota
	leftDirection
)

const enemyShootCooldown = time.Millisecond * 500

// enemySquad represents a common entity that describes a bunch of enemies.
type enemySquad struct {
	basicEnemies []*basicEnemy
	bulletPool   []*basicEnemyBullet
	direction    int
	lastShoot    time.Time
}

func newEnemySquad(rows int) (*enemySquad, error) {
	es := &enemySquad{direction: rightDirection}
	es.basicEnemies = make([]*basicEnemy, 0, rows+7)

	// Init enemies
	for i := 0; i < 8; i++ {
		for j := 0; j < rows; j++ {
			x := (float64(i)/10)*windowWidth + basicEnemySize/2
			y := float64(j)*basicEnemySize + basicEnemySize/2
			be, err := newBasicEnemy(x, y)
			if err != nil {
				return nil, err
			}
			es.basicEnemies = append(es.basicEnemies, be)
		}
	}

	// init basic's enemy bullet pool
	if err := es.initBulletPool(); err != nil {
		return nil, err
	}

	return es, nil
}

func (es *enemySquad) draw(dst *ebiten.Image) {
	// Draw enemies
	for _, be := range es.basicEnemies {
		be.draw(dst)
	}

	// Draw bullets
	for _, b := range es.bulletPool {
		b.draw(dst)
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
				be.currentSequence = basicEnemyDestroySequence
				b.isActive = false
			}
		}

		// Update enemies' frames
		be.update()
	}

	// Check if we need to go down
	if es.hasChangedDirection() {
		es.moveEnemiesDown()
	}

	// Make bottom enemies shoot
	es.randomEnemyShoot(es.getAllBottomEnemies())

	// Update bullets
	for _, b := range es.bulletPool {
		b.update()
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

func (es *enemySquad) getAllBottomEnemies() []*basicEnemy {
	allXPositions := es.getAllXPositions()
	bottomEnemies := make([]*basicEnemy, 0)

	for _, x := range allXPositions {
		var (
			bestY       float64
			lowestEnemy *basicEnemy
		)
		for _, e := range es.basicEnemies {
			if e.x == x && e.y > bestY && e.isActive {
				bestY = e.y
				lowestEnemy = e
			}
		}

		if lowestEnemy != nil {
			bottomEnemies = append(bottomEnemies, lowestEnemy)

		}
	}

	return bottomEnemies
}

func (es *enemySquad) getAllXPositions() []float64 {
	xs := make([]float64, 0, len(es.basicEnemies))
	for _, e := range es.basicEnemies {
		xs = append(xs, e.x)
	}

	return xs
}

func (es *enemySquad) initBulletPool() error {
	for i := 0; i < 10; i++ {
		pb, err := newBasicEnemyBullet()
		if err != nil {
			return err
		}
		es.bulletPool = append(es.bulletPool, pb)
	}

	return nil
}

func (es *enemySquad) getBulletFromPool() (*basicEnemyBullet, bool) {
	for _, b := range es.bulletPool {
		if !b.isActive {
			return b, true
		}
	}

	return nil, false
}

func (es *enemySquad) randomEnemyShoot(enemies []*basicEnemy) {
	if time.Since(es.lastShoot) < enemyShootCooldown {
		return
	}
	source := rand.New(rand.NewSource(time.Now().Unix()))
	if len(enemies) == 0 {
		return
	}
	shooter := enemies[source.Intn(len(enemies))]

	b, ok := es.getBulletFromPool()
	if ok {
		b.x = shooter.x
		b.y = shooter.y
		b.angle = 270 * (math.Pi / 180)
		b.isActive = true
		es.lastShoot = time.Now().UTC()
	}
}
