package poisson_test

import (
	"github.com/ojrac/opensimplex-go"
	"github.com/plandem/poisson-disk-sampling"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

const WIDTH = 1024
const HEIGHT = 1024
const NUM_POINTS = 50000

func Example_basic() {
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
	f, _ := os.OpenFile("./example-screenshots/output-basic.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

	// Output:
}

func Example_minDistance() {
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
	f, _ := os.OpenFile("./example-screenshots/output-min-distance.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

	// Output:
}

func Example_areaCircleFilter() {
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
	f, _ := os.OpenFile("./example-screenshots/output-area-circle-filter.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

	// Output:
}

func Example_areaRectangleFilter() {
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
	f, _ := os.OpenFile("./example-screenshots/output-area-rectangle-filter.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

	// Output:
}

func Example_withSimplexNoisePostFilter() {
	//generate poisson disk samplings
	points := poisson.NewPoissonDisk(
		NUM_POINTS,
		poisson.WithPostFilter(poisson.NewSimplexNoiseFilter(WIDTH, HEIGHT, WIDTH/8, 0)),
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
	f, _ := os.OpenFile("./example-screenshots/output-simplex-post-filter.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

	// Output:
}

func Example_withPngNoisePostFilter() {
	//generate png file with simple noise
	noise := opensimplex.NewWithSeed(0)
	noisePng := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	featureSize := WIDTH / 4

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			value := noise.Eval2(float64(x)/float64(featureSize), float64(y)/float64(featureSize))
			c := 0x01 * uint8((value+1)*127)
			noisePng.Set(x, y, color.RGBA{c, c, c, 255})
		}
	}

	f, _ := os.OpenFile("./example-screenshots/noise.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, noisePng)

	//generate poisson disk samplings
	points := poisson.NewPoissonDisk(
		NUM_POINTS,
		poisson.WithPostFilter(poisson.NewGrayScalePngFilter(WIDTH, HEIGHT, "./example-screenshots/noise.png")),
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
	f, _ = os.OpenFile("./example-screenshots/output-png-post-filter.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

	// Output:
}

func Example_fullFeatured() {
	//generate poisson disk samplings
	points := poisson.NewPoissonDisk(
		NUM_POINTS,
		poisson.WithTries(10000),
		poisson.WithMinDistance(0.01),
		poisson.WithGenerator(poisson.NewBasicGenerator(100)),
		poisson.WithAreaFilter(poisson.NewCircleFilter(0.5, 0.5, 0.25)),
		poisson.WithPostFilter(poisson.NewSimplexNoiseFilter(WIDTH, HEIGHT, WIDTH/16, 0)),
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
	f, _ := os.OpenFile("./example-screenshots/output-full-featured.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

	// Output:
}
