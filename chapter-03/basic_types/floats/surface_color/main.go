package main

import (
	"fmt"
	"image/color"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	//zmin, zmax := zminmax()
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			//color := computeColor(i, j, zmin, zmax)
			r, g, b, _ := computeColor(i, j).RGBA()
			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' style='fill: #%.2x%.2x%.2x'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, uint8(r), uint8(g), uint8(b))
		}
	}
	fmt.Println("</svg>")
}

/*
	This implementation of computing color based on cell height originally comes
	from https://github.com/ladrift/gopl-exercises/blob/master/ch3/ex3.3/main.go.
*/
func computeColor(i, j int) color.Color {
	x := xyrange * (float64(i)/cells - 0.5) // -15.0..+15.0
	y := xyrange * (float64(j)/cells - 0.5) // -15.0..+15.0

	z := f(x, y)

	r := 255 * (z + 1) / 2
	b := 255 * (1 - z) / 2

	return color.RGBA{uint8(r), 0x00, uint8(b), 0x00}
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5) // -15.0..+15.0
	y := xyrange * (float64(j)/cells - 0.5) // -15.0..+15.0

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0, 0)
	return math.Sin(r) / r
}
