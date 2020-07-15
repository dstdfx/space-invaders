package main

import (
	"math"
)

type circle struct {
	x float64
	y float64
	radius float64
}

func collides(c1, c2 circle) bool {
	dist := math.Sqrt(math.Pow(c2.x-c1.x, 2) +
		math.Pow(c2.y-c1.y, 2))

	return dist <= c1.radius+c2.radius
}
