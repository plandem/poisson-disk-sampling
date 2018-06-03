package poisson

type circleFilter struct {
	center Point
	radius float64
}

//NewCircleFilter returns a PointFilter that returns only points inside of circle
func NewCircleFilter(x, y, r float64) PointFilter {
	return &circleFilter{
		center: Point{
			x,
			y,
		},
		radius: r,
	}
}

//filters points inside of circle
func (f *circleFilter) Filter(point *Point, settings *options) bool {
	cx := point.X - f.center.X
	cy := point.Y - f.center.Y
	return (cx*cx + cy*cy) <= f.radius
}

