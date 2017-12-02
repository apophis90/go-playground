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
		epsX                   = (xmax - xmin) / width
		epsY                   = (ymax - ymin) / height
	)
	offx := []float64{-epsX, epsX}
	offy := []float64{-epsY, epsY}

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Iterate over 1024x1024 raster
	for py := 0; py < height; py++ {
		// Transfer pixel (x,y) into 4x4 [-2,2] raster
		y := float64(py)/height*(ymax-ymin) + ymin // [-2, 2]
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin // [-2, 2]
			subPixels := make([]color.Color, 0)

			// See https://github.com/torbiak/gopl/blob/master/ex3.6/main.go
			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					// Compute subpixel values by considering neighbor pixels.
					z := complex(x+offx[i], y+offy[j])
					subPixels = append(subPixels, mandelbrot(z))
				}
			}
			img.Set(px, py, average(subPixels))
		}
	}
	png.Encode(os.Stdout, img) // Errors are ignored.
}

func average(colors []color.Color) color.Color {
	var r, g, b, a uint16
	n := len(colors)

	for _, c := range colors {
		rCol, gCol, bCol, aCol := c.RGBA()
		r += uint16(rCol / uint32(n))
		g += uint16(gCol / uint32(n))
		b += uint16(bCol / uint32(n))
		a += uint16(aCol / uint32(n))
	}
	return color.RGBA64{r, g, b, a}
}

/*
  The Mandelbrot function tests whether repeatedly squaring and adding the
  number the point (px, py) represents "escapes" the circle of radius 2. If this
  is the case, the corresponding point is shaded by the number of iterations
  required to escape.
*/
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
