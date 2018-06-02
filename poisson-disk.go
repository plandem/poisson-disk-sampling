package poisson

import (
	"math"
)

//NewPoissonDisk returns sampling points with provided options
func NewPoissonDisk(points int, opts ...Option) []*Point {
	//make a copy of default settings
	defaultSettings := defaultOptions

	//fill default settings with provided options
	settings := &defaultSettings
	for _, o := range opts {
		o(settings)
	}

	//if there is no minimum distance, then calculate it via number of points
	if settings.minDistance <= 0 {
		settings.minDistance = math.Sqrt(float64(points)) / float64(points)
	}

	cellSize := float64(settings.minDistance) / math.Sqrt2
	width := int(math.Ceil(1.0 / cellSize))
	height := int(math.Ceil(1.0 / cellSize))

	grid := newGrid(width, height, cellSize)
	maxDist := settings.minDistance * settings.minDistance

	queue := make([]*Point, 0)
	samplePoints := make([]*Point, 0)

	var point, newPoint *Point

	//generate the first point randomly
	for {
		if newPoint = (&Point{settings.generator.Float(), settings.generator.Float()}); settings.areaFilter.Filter(newPoint, settings) {
			break
		}
	}

	queue = append(queue, newPoint)
	samplePoints = append(samplePoints, newPoint)
	grid.SetPoint(newPoint)

	//generate other points from points in queue.
	for len(queue) > 0 && len(samplePoints) < points {
		point, queue = queue[0], queue[1:]

		for i := 0; i < settings.tries; i++ {
			newPoint = point.RandomPointAround(settings.minDistance, settings.generator)

			//if point inside of allowed area and there is no any neighbourhood point around, then add it
			if settings.areaFilter.Filter(newPoint, settings) && !grid.IsNeighbourhood(newPoint, maxDist) {
				queue = append(queue, newPoint)
				samplePoints = append(samplePoints, newPoint)
				grid.SetPoint(newPoint)
			}
		}
	}

	if settings.postFilter == nil {
		return samplePoints
	}

	filteredPoints := make([]*Point, 0)
	for _, point := range samplePoints {
		if settings.postFilter.Filter(point, settings) {
			filteredPoints = append(filteredPoints, point)
		}
	}

	return filteredPoints
}
