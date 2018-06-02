package poisson

import (
	"image"
	"image/png"
	"os"
)

type grayScalePngFilter struct {
	*noiseFilter
	density image.Image
}

//NewGrayScalePngFilter returns a PointFilter that uses a gray scale image for filtering
func NewGrayScalePngFilter(width, height int, pngFileName string) PointFilter {
	f, _ := os.Open(pngFileName)
	defer f.Close()
	density, _ := png.Decode(f)

	filter := &grayScalePngFilter{
		noiseFilter: &noiseFilter{
			width:  width,
			height: height,
		},
		density: density,
	}

	filter.noiseFilter.PointFilter = filter.noiseFilter
	filter.noiseValue = filter
	return filter
}

func (f *grayScalePngFilter) GetNoiseValue(x, y int) float64 {
	C := f.density.At(x, y)
	V, _, _, _ := C.RGBA()

	return float64(V>>8) / 255.0
}
