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
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img) // Errors are ignored.
}

var iterationStep = func(z complex128) complex128 {
	f := func(z complex128) complex128 {
		return z*z*z*z - 1
	}
	firstDerivative := func(z complex128) complex128 {
		return 4 * z * z * z
	}
	return z - f(z)/firstDerivative(z)
}

func newton(z complex128) color.Color {
	const iterations = 40
	const contrast = 5

	for i := 0; i < iterations; i++ {
		z = iterationStep(z)
		if cmplx.Abs(z) < 1e-6 {
			// Choose color based on root
		}
	}
	return color.Black
}
