package poisson

import (
	"math"
)

//Point holds coordinate x,y
type Point struct {
	X float64
	Y float64
}

//Distance returns a distance between point and other points
func (point *Point) Distance(other *Point) float64 {
	dx, dy := point.X-other.X, point.Y-other.Y
	return dx*dx + dy*dy
}

//RandomPointAround generates random point in annulus [minDist, 2 * minDist] using generator to get random radius and angle
func (point *Point) RandomPointAround(minDist float64, generator RandomGenerator) *Point {
	//random radius between minDist and 2 * minDist
	radius := minDist * (generator.Float() + 1)

	//random angle
	angle := 2 * math.Pi * generator.Float()

	//the new point is generated around the point (x, y)
	newX := point.X + radius*math.Cos(angle)
	newY := point.Y + radius*math.Sin(angle)
	return &Point{newX, newY}
}
