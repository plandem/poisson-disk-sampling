package poisson

import (
	"github.com/ojrac/opensimplex-go"
)

type simplexNoiseFilter struct {
	featureSize float64
	noise       *opensimplex.Noise
}

//NewSimplexNoiseFilter returns a PointFilter that uses the simplex noise algorithm
func NewSimplexNoiseFilter(featureSize int, noiseSeed int64) PointFilter {
	filter := &simplexNoiseFilter{
		featureSize: float64(featureSize),
		noise:       opensimplex.NewWithSeed(noiseSeed),
	}

	return filter
}

//filters points based on noise value
func (f *simplexNoiseFilter) Filter(point *Point, settings *options) bool {
	R := settings.generator.Float()

	//get noise value for point
	N := f.noise.Eval2(point.X*f.featureSize, point.Y*f.featureSize)

	//normalize it to [0, 1] from [-1, 1]
	P := (N + 1) / 2
	return R < P
}
