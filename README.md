# Poisson Disk Sampling
[![Build Status](https://travis-ci.org/plandem/poisson-disk-sampling.svg?branch=master)](https://travis-ci.org/plandem/poisson-disk-sampling) 
[![Go Report Card](https://goreportcard.com/badge/github.com/plandem/poisson-disk-sampling)](https://goreportcard.com/report/github.com/plandem/poisson-disk-sampling)
[![GoDoc](https://godoc.org/github.com/plandem/poisson-disk-sampling?status.svg)](https://godoc.org/github.com/plandem/poisson-disk-sampling)
[![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/plandem/poisson-disk-sampling/master/LICENSE)
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.me/plandem)
<!-- [![Code Coverage](https://codecov.io/gh/plandem/poisson-disk-sampling/branch/master/graph/badge.svg)](https://codecov.io/gh/plandem/poisson-disk-sampling) -->

Based on article: http://devmag.org.za/2009/05/03/poisson-disk-sampling/

```go
package main

import (
	"os"
	"image"
	"image/draw"
	"image/color"
	"image/png"
	"github.com/plandem/poisson-disk-sampling"
)

func main() {
	width, height, numPoints := 1024, 1024, 1000
	
	//generate poisson disk samplings
	points := poisson.NewPoissonDisk(numPoints)

	//draw result
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)
	for _, point := range points {
		x := int(point.X * float64(width))
		y := int(point.Y * float64(height))
		img.Set(x, y, color.White)
	}

	//save result png
	f, _ := os.OpenFile("output.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
```

For for examples at file **example_test.go**

### basic example
![Basic Example](https://raw.github.com/plandem/poisson-disk-sampling/master/example-output-basic.png)

### with minimum distance
![Minimum Distance Example](https://raw.github.com/plandem/poisson-disk-sampling/master/example-output-min-distance.png)

### with circle area filter
![Circle Area Filter Example](https://raw.github.com/plandem/poisson-disk-sampling/master/example-output-area-circle-filter.png)

### with rectangle area filter
![Rectangle Area Filter Example](https://raw.github.com/plandem/poisson-disk-sampling/master/example-output-area-rectangle-filter.png)

### with simplex noise post filter
![Simplex Noise Post Filter Example](https://raw.github.com/plandem/poisson-disk-sampling/master/example-output-simplex-post-filter.png)

### with PNG map post filter
![PNG Map](https://raw.github.com/plandem/poisson-disk-sampling/master/example-noise.png)
![PNG Post Filter Example](https://raw.github.com/plandem/poisson-disk-sampling/master/example-output-png-post-filter.png)

### full featured example
![Full Featured Example](https://raw.github.com/plandem/poisson-disk-sampling/master/example-output-full-featured.png)


