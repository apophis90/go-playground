// Mandelbrot emits a PNG image of a Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	// Each point of the 1024 x 1024 raster is projected to the Gaussian plane
	// (-2..+2) as a complex number.
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin // [-2, 2]
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin // [-2, 2]
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.SetRGBA(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // Errors are ignored.
}

/*
  The Mandelbrot function tests whether repeatedly squaring and adding the
  number the point (px, py) represents "escapes" the circle of radius 2. If this
  is the case, the corresponding point is shaded by the number of iterations
  required to escape.
*/
func mandelbrot(z complex128) color.RGBA {
	const iterations = 200
	const contrastR = 10
	const contrastG = 15
	const contrastB = 5

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255 - contrastR*n, 255 - contrastG*n, 255 - contrastB*n, 255}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}
