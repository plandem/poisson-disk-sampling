package poisson_test

import (
	"os"
	"image"
	"image/draw"
	"image/color"
	"image/png"
	"github.com/plandem/poisson-disk-sampling"
	"github.com/ojrac/opensimplex-go"
)

const WIDTH = 1024
const HEIGHT = 1024
const NUM_POINTS = 50000

func ExampleBasic() {
	//generate poisson disk samplings
	points := poisson.NewPoissonDisk(NUM_POINTS)

	//draw result
	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)
	for _, point := range points {
		x := int(point.X * float64(WIDTH))
		y := int(point.Y * float64(HEIGHT))
		img.Set(x, y, color.White)
	}

	//save result png
	f, _ := os.OpenFile("example-output-basic.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

	// Output:
}

func ExampleMinDistance() {
	//generate poisson disk samplings
	points := poisson.NewPoissonDisk(
		NUM_POINTS,
		poisson.WithMinDistance(0.01),
	)

	//draw result
	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)
	for _, point := range points {
		x := int(point.X * float64(WIDTH))
		y := int(point.Y * float64(HEIGHT))
		img.Set(x, y, color.White)
	}

	//save result png
	f, _ := os.OpenFile("example-output-min-distance.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

	// Output:
}

func ExampleAreaCircleFilter() {
	//generate poisson disk samplings
	points := poisson.NewPoissonDisk(
		NUM_POINTS,
		poisson.WithAreaFilter(poisson.NewCircleFilter(0.5, 0.5, 0.25)),
	)

	//draw result
	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)
	for _, point := range points {
		x := int(point.X * float64(WIDTH))
		y := int(point.Y * float64(HEIGHT))
		img.Set(x, y, color.White)
	}

	//save result png
	f, _ := os.OpenFile("example-output-area-circle-filter.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

	// Output:
}

func ExampleAreaRectangleFilter() {
	//generate poisson disk samplings
	points := poisson.NewPoissonDisk(
		NUM_POINTS,
		poisson.WithAreaFilter(poisson.NewRectangleFilter(0.2, 0.8)),
	)

	//draw result
	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)
	for _, point := range points {
		x := int(point.X * float64(WIDTH))
		y := int(point.Y * float64(HEIGHT))
		img.Set(x, y, color.White)
	}

	//save result png
	f, _ := os.OpenFile("example-output-area-rectangle-filter.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

	// Output:
}

func ExampleWithSimplexNoisePostFilter() {
	//generate poisson disk samplings
	points := poisson.NewPoissonDisk(
		NUM_POINTS,
		poisson.WithPostFilter(poisson.NewSimplexNoiseFilter(WIDTH, HEIGHT, WIDTH / 8, 0)),
	)

	//draw result
	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)

	for _, point := range points {
		x := int(point.X * float64(WIDTH))
		y := int(point.Y * float64(HEIGHT))
		img.Set(x, y, color.White)
	}

	//save result png
	f, _ := os.OpenFile("example-output-simplex-post-filter.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

	// Output:
}

func ExampleWithPngNoisePostFilter() {
	//generate png file with simple noise
	noise := opensimplex.NewWithSeed(0)
	noisePng := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	featureSize := WIDTH / 4

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			value := noise.Eval2(float64(x) / float64(featureSize), float64(y) / float64(featureSize))
			c := 0x01 * uint8((value + 1) * 127)
			noisePng.Set(x, y, color.RGBA{c,c,c,255})
		}
	}

	f, _ := os.OpenFile("example-noise.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, noisePng)

	//generate poisson disk samplings
	points := poisson.NewPoissonDisk(
		NUM_POINTS,
		poisson.WithPostFilter(poisson.NewGrayscalePngFilter(WIDTH, HEIGHT, "example-noise.png")),
	)

	//draw result
	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)

	for _, point := range points {
		x := int(point.X * float64(WIDTH))
		y := int(point.Y * float64(HEIGHT))
		img.Set(x, y, color.White)
	}

	//save result png
	f, _ = os.OpenFile("example-output-png-post-filter.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

	// Output:
}

func ExampleFullFeatured() {
	//generate poisson disk samplings
	points := poisson.NewPoissonDisk(
		NUM_POINTS,
		poisson.WithTries(10000),
		poisson.WithMinDistance(0.01),
		poisson.WithGenerator(poisson.NewBasicGenerator(100)),
		poisson.WithAreaFilter(poisson.NewCircleFilter(0.5, 0.5, 0.25)),
		poisson.WithPostFilter(poisson.NewSimplexNoiseFilter(WIDTH, HEIGHT, WIDTH / 16, 0)),
	)

	//draw result
	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)

	for _, point := range points {
		x := int(point.X * float64(WIDTH))
		y := int(point.Y * float64(HEIGHT))
		img.Set(x, y, color.White)
	}

	//save result png
	f, _ := os.OpenFile("example-output-full-featured.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

	// Output:
}

