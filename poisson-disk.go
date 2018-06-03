package poisson

import (
	"math"
)

//NewPoissonDisk returns sampling points with provided options
func NewPoissonDisk(numPoints int, opts ...Option) []*Point {
	//make a copy of default settings
	defaultSettings := defaultOptions

	//fill default settings with provided options
	settings := &defaultSettings
	for _, o := range opts {
		o(settings)
	}

	//if there is no minimum distance, then calculate it via number of points
	if settings.minDistance <= 0 {
		settings.minDistance = math.Sqrt(float64(numPoints)) / float64(numPoints)
	}

	maxDist := settings.minDistance * settings.minDistance

	cellSize := float64(settings.minDistance) / math.Sqrt2
	width := int(math.Ceil(1.0 / cellSize))
	height := int(math.Ceil(1.0 / cellSize))
	grid := NewGrid(width, height, cellSize)

	var point, newPoint *Point

	queue := make([]*Point, 0)
	samplePoints := make([]*Point, 0)

	//populate grid with predefined points
	//because of dimension of the grid - its possible that few points will be related to same cell, so last one will be used
	if len(settings.points) > 0 {
		for _, point = range settings.points {
			grid.SetPoint(point)
		}
	}

	//resolve first point to start from
	if newPoint = settings.startPoint; newPoint != nil {
		//accept only valid startPoint
		if *newPoint == (Point{}) || !settings.areaFilter.Filter(newPoint, settings) {
			panic("You must provide a valid start point with a valid coordinates that fit area")
		}
	} else {
		//generate the first point randomly
		for {
			if newPoint = NewRandomPoint(settings.generator); settings.areaFilter.Filter(newPoint, settings) {
				break
			}
		}
	}

	queue = append(queue, newPoint)
	samplePoints = append(samplePoints, newPoint)
	grid.SetPoint(newPoint)

	//generate other points from points in queue.
	for len(queue) > 0 && len(samplePoints) < numPoints {
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
