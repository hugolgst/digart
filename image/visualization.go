package image

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

// DrawData draws in the given image all the points for the parsed decimals
func DrawData(img image.RGBA, parsedDigits [][][]int, radius, image int) {
	fmt.Println("Drawing the data")

	// Iterate through segments
	for segmentIndex, segment := range parsedDigits {
		// Browse all the lines
		for line := 0; line < len(segment); line++ {
			// Calculate the radius for the actual line
			r := float64(radius*5/4 + line * radius/15)

			// Arbitrary number of iterations
			iterations := 55
			t := math.Pi/float64(iterations)

			// Draw the points
			for i := 0; i < int(iterations*2); i++ {
				// Calculate the coordinates
				x1 := r*math.Cos(t)+float64(image/2)
				y1 := r*math.Sin(t)+float64(image/2)

				t += math.Pi/float64(iterations)
				size := iterations/5

				// Only iterate through the wanted part of the circle
				if i < segmentIndex*size || i >= size*(segmentIndex+1) + len(segment[line]) - size {
					continue
				}

				// Choose the correct color
				colorIndex := segment[line][i%size]
				var c color.RGBA
				if colorIndex == -1 {
					c = color.RGBA{A: 255}
				} else {
					c = colors[colorIndex]
				}

				// Then draw the point
				DrawPoint(img, int(x1), int(y1), 9, c)
			}
		}
	}
}
