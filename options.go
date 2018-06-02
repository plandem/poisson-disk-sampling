package poisson

//Option is function to update required setting
type Option func(*options)

type options struct {
	//tries is a number of attempts to generate a new point
	tries int

	//minDistance is a minimum distance between two points
	minDistance float64

	//generator to use for random numbers
	generator RandomGenerator

	//areaFilter filters candidate point during 'generation' phase to get valid points inside of are. Allowed area boundaries - [0.0, 1.1] box
	areaFilter PointFilter

	//postFilter filters already valid point for additional condition
	postFilter PointFilter
}

var defaultOptions = options{
	tries:      30,
	areaFilter: NewRectangleFilter(0, 1),
	generator:  NewBasicGenerator(0),
}

//WithTries set number of tries to generate a new point
func WithTries(tries int) Option {
	return func(o *options) {
		o.tries = tries
	}
}

//WithGenerator set a random number genetator
func WithGenerator(generator RandomGenerator) Option {
	return func(o *options) {
		o.generator = generator
	}
}

//WithMinDistance set a minimum distance between any two points
func WithMinDistance(distance float64) Option {
	return func(o *options) {
		o.minDistance = distance
	}
}

//WithAreaFilter set an area filter for candidate points
func WithAreaFilter(filter PointFilter) Option {
	return func(o *options) {
		o.areaFilter = filter
	}
}

//WithPostFilter set a post filter to drop points that do not meet required condition
func WithPostFilter(filter PointFilter) Option {
	return func(o *options) {
		o.postFilter = filter
	}
}
