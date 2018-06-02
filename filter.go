package poisson

//PointFilter interface is for filters to return valid points. In general, any point must be inside of [0.0, 1.1] box
type PointFilter interface {
	Filter(point *Point, settings *options) (bool)
}

type circleFilter struct {
	center Point
	radius float64
}

type rectangleFilter struct {
	min float64
	max float64
}

type noiseValue interface {
	GetNoiseValue(x, y int) (float64)
}

type noiseFilter struct {
	PointFilter
	noiseValue
	width int
	height int
}

func NewCircleFilter(x, y, r float64) (PointFilter) {
	return &circleFilter {
		center: Point{
			x,
			y,
		},
		radius: r,
	}
}

func NewRectangleFilter(min, max float64) (PointFilter) {
	return &rectangleFilter{
		min: min,
		max: max,
	}
}

//filters points inside of circle
func (f *circleFilter) Filter(point *Point, settings *options) (bool) {
	cx := point.X - f.center.X
	cy := point.Y - f.center.Y
	return (cx * cx + cy * cy) <= f.radius
}

//filters points inside of rectangle
func (f *rectangleFilter) Filter(point *Point, settings *options) (bool) {
	return (point.X >= f.min && point.Y >= f.min) && (point.X <= f.max && point.Y <= f.max)
}

//filters points based on noise value
func (f *noiseFilter) Filter(point *Point, settings *options) (bool) {
	x := int(point.X * float64(f.width))
	y := int(point.Y * float64(f.height))
	R := settings.generator.RandomFloat()
	P := f.GetNoiseValue(x, y)

	return R > P
}
