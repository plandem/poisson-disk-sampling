package poisson

import (
	"os"
	"image"
	"image/png"
)

type GrayscalePngFilter struct {
	*noiseFilter
	density image.Image
}

func NewGrayscalePngFilter(width, height int, pngFileName string) (PointFilter) {
	f,_ := os.Open(pngFileName)
	defer f.Close()
	density, _ := png.Decode(f)

	filter := &GrayscalePngFilter {
		noiseFilter: &noiseFilter {
			width: width,
			height: height,
		},
		density: density,
	}

	filter.noiseFilter.PointFilter = filter.noiseFilter
	filter.noiseValue = filter
	return filter
}

func (f *GrayscalePngFilter) GetNoiseValue(x, y int) (float64) {
	C := f.density.At(x, y)
	V,_,_,_ := C.RGBA()

	return float64(V >> 8) / 255.0
}