package poisson

type rectangleFilter struct {
	min float64
	max float64
}

//NewRectangleFilter returns a PointFilter that returns only points inside of rectangle
func NewRectangleFilter(min, max float64) PointFilter {
	return &rectangleFilter{
		min: min,
		max: max,
	}
}

//filters points inside of rectangle
func (f *rectangleFilter) Filter(point *Point, settings *options) bool {
	return (point.X >= f.min && point.Y >= f.min) && (point.X <= f.max && point.Y <= f.max)
}
