package image

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

var colors = []color.RGBA{
	{18, 67, 89, 255},
	{58, 165, 229, 255},
	{78, 118, 33, 255},
	{77, 206, 105, 255},
	{104, 6, 9, 255},
	{229, 22, 37, 255},
	{147, 82, 17, 255},
	{225, 205, 31, 255},
	{63, 32, 77, 255},
	{139, 85, 193, 255},
}

// DrawPoint draws a point at (`x`; `y`) in the given image with the given color
func DrawPoint(img image.RGBA, x, y, width int, color color.RGBA) {
	// Draw multiples layers for the width
	for w := 0; w < width; w++ {
		r := float64(1+w)
		iterations := 2000
		t := math.Pi/float64(iterations)

		// Draw the point
		for i := 0; i < int(iterations*2); i++ {
			// Calculate the coordinates
			x1 := r*math.Cos(t)+float64(x)
			y1 := r*math.Sin(t)+float64(y)

			// Draw the point
			img.Set(int(x1), int(y1), color)

			t += math.Pi/float64(iterations)
		}
	}
}

// DrawCircle draws in the given image the middle circle of the visualization with
// the 10 segments in color.
func DrawCircle(img image.RGBA, radius, width, image int) {
	fmt.Println("Drawing main circle")

	// Draw multiples layers for the width
	for w := 0; w < width; w++ {
		iterations := 2000

		// Draw a circle with a parametric equation
		t := math.Pi / float64(iterations)
		for i := 0; i < iterations*2; i++ {
			// Choose a color for each part
			var color color.RGBA
			for n := range colors {
				// Calculate the part of the circle `t` is in
				part := t/(math.Pi/5)
				color = colors[int(part)%(n+1)]
			}

			// Calculate the coordinates for `t` in the center of the image
			x := float64(radius + w) * math.Cos(t) + float64(image/2)
			y := float64(radius + w) * math.Sin(t) + float64(image/2)
			// Draw the point
			img.Set(int(x), int(y), color)

			t += math.Pi/float64(iterations)
		}
	}
}
