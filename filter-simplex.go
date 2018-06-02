package poisson

import (
	"github.com/ojrac/opensimplex-go"
)

type SimplexNoiseFilter struct {
	*noiseFilter
	featureSize float64
	noise *opensimplex.Noise
}

func NewSimplexNoiseFilter(width, height int, featureSize int, noiseSeed int64) (PointFilter) {
	filter := &SimplexNoiseFilter {
		noiseFilter: &noiseFilter {
			width: width,
			height: height,
		},
		featureSize: float64(featureSize),
		noise: opensimplex.NewWithSeed(noiseSeed),
	}

	filter.noiseFilter.PointFilter = filter.noiseFilter
	filter.noiseValue = filter

	return filter
}

func (f *SimplexNoiseFilter) GetNoiseValue(x, y int) (float64) {
	N := f.noise.Eval2(float64(x) / float64(f.featureSize), float64(y) / float64(f.featureSize))

	//simplex noise returns value in range [-1, 1], so we return normalized to point's allowed boundaries [0, 1]
	return (N - (-1) ) / (1 - (-1))
}
