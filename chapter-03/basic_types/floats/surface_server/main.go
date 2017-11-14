package main

import (
	"fmt"
	"image/color"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height    = 600, 320            // canvas size in pixels
	cells            = 100                 // number of grid cells
	xyrange          = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale          = width / 2 / xyrange // pixels per x or y unit
	zscale           = height * 0.4        // pixels per z unit
	angle            = math.Pi / 6         // angle of x, y axes (=30°)
	red, green, blue = "red", "green", "blue"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "image/svg+xml")

		var maxColor string
		var minColor string

		minColorParam, ok := req.URL.Query()["minColor"]
		if !ok || len(minColorParam) < 1 {
			minColor = "blue"
		} else {
			minColor = minColorParam[0]
		}
		log.Printf("Min color is: %s\n", minColor)

		maxColorParam, ok := req.URL.Query()["maxColor"]
		if !ok || len(maxColorParam) < 1 {
			maxColor = "red"
		} else {
			maxColor = maxColorParam[0]
		}
		log.Printf("Max color is: %s\n", maxColor)

		svg(rw, minColor, maxColor)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

func svg(w io.Writer, minCol, maxCol string) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			//color := computeColor(i, j, zmin, zmax)
			r, g, b, _ := computeColor(i, j, minCol, maxCol).RGBA()
			fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' style='fill: #%.2x%.2x%.2x'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, uint8(r), uint8(g), uint8(b))
		}
	}
	fmt.Fprintln(w, "</svg>")
}

/*
	This implementation of computing color based on cell height originally comes
	from https://github.com/ladrift/gopl-exercises/blob/master/ch3/ex3.3/main.go.
*/
func computeColor(i, j int, minCol, maxCol string) color.Color {
	x := xyrange * (float64(i)/cells - 0.5) // -15.0..+15.0
	y := xyrange * (float64(j)/cells - 0.5) // -15.0..+15.0

	z := f(x, y)

	r, g, b := processInput(minCol, maxCol)

	//r := 255 * (z + 1) / 2
	//b := 255 * (1 - z) / 2

	return color.RGBA{uint8(r(z)), uint8(g(z)), uint8(b(z)), 0x00}
}

func processInput(minCol, maxCol string) (func(float64) float64, func(float64) float64, func(float64) float64) {
	var r, g, b func(float64) float64

	if minCol == red && maxCol == red {
		r = func(z float64) float64 {
			if z > 0 {
				return 255 * (z + 1) / 2
			}
			return 255 * (1 - z) / 2
		}
	} else if minCol == red {
		r = func(z float64) float64 {
			return 255 * (1 - z) / 2
		}
	} else if maxCol == red {
		r = func(z float64) float64 {
			return 255 * (z + 1) / 2
		}
	} else {
		r = func(z float64) float64 {
			return z * 0
		}
	}

	if minCol == green && maxCol == green {
		g = func(z float64) float64 {
			if z > 0 {
				return 255 * (z + 1) / 2
			}
			return 255 * (1 - z) / 2
		}
	} else if minCol == green {
		g = func(z float64) float64 {
			return 255 * (1 - z) / 2
		}
	} else if maxCol == green {
		g = func(z float64) float64 {
			return 255 * (z + 1) / 2
		}
	} else {
		g = func(z float64) float64 {
			return z * 0
		}
	}

	if minCol == blue && maxCol == blue {
		b = func(z float64) float64 {
			if z > 0 {
				return 255 * (z + 1) / 2
			}
			return 255 * (1 - z) / 2
		}
	} else if minCol == blue {
		b = func(z float64) float64 {
			return 255 * (1 - z) / 2
		}
	} else if maxCol == blue {
		b = func(z float64) float64 {
			return 255 * (z + 1) / 2
		}
	} else {
		b = func(z float64) float64 {
			return z * 0
		}
	}
	return r, g, b
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
