package poisson

import (
	"os"
	"image"
	"image/png"
)

type grayScalePngFilter struct {
	density image.Image
	width   float64
	height  float64
}

//NewGrayScalePngFilter returns a PointFilter that uses a gray scale image for filtering
func NewGrayScalePngFilter(pngFileName string) PointFilter {
	f, _ := os.Open(pngFileName)
	defer f.Close()
	density, _ := png.Decode(f)

	bounds := density.Bounds()
	filter := &grayScalePngFilter{
		density: density,
		width:   float64(bounds.Dx()),
		height:  float64(bounds.Dy()),
	}

	return filter
}

//filters points based on noise value
func (f *grayScalePngFilter) Filter(point *Point, settings *options) bool {
	R := settings.generator.Float()

	//get color from map
	C := f.density.At(int(point.X*f.width), int(point.Y*f.height))

	//extract red value (in gray scale image each part equals)
	V, _, _, _ := C.RGBA()

	//reduce value to uint8 and normalize to [0, 1] from [0, 255]
	P := float64(V>>8) / 255.0

	return R < P
}
