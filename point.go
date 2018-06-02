package poisson

import (
	"math"
)

type Point struct {
	X float64
	Y float64
}

func (point *Point) Distance(other *Point) (float64) {
	dx, dy := point.X - other.X, point.Y - other.Y
	return dx * dx + dy * dy
}

//generate random point in annulus [r, 2r]
func (point *Point) RandomPointAround(minDist float64, generator Generator) (*Point) {
	//random radius between minDist and 2 * minDist
	radius := minDist * (generator.RandomFloat() + 1)

	//random angle
	angle := 2 * math.Pi * generator.RandomFloat()

	//the new point is generated around the point (x, y)
	newX := point.X + radius * math.Cos(angle)
	newY := point.Y + radius * math.Sin(angle)
	return &Point{newX, newY }
}


