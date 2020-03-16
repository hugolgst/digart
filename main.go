package main

import (
	"fmt"
	"github.com/hugolgst/digart/digits"
	digartImage "github.com/hugolgst/digart/image"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"strconv"
)

// CreateImage returns the created image of `width` and `height` dimensions.
func CreateImage(width int, height int, background color.RGBA) *image.RGBA {
	rect := image.Rect(0, 0, width, height)
	img := image.NewRGBA(rect)
	draw.Draw(img, img.Bounds(), &image.Uniform{C: background}, image.ZP, draw.Src)

	return img
}

func main() {
	// File to use
	var filePath string
	if len(os.Args) < 2 {
		fmt.Println("The file path was not found, using resources/pi by default.")
		filePath = "resources/pi.txt"
	} else {
		filePath = os.Args[1]
	}

	// Number of digits
	var d string
	if len(os.Args) < 3 {
		fmt.Println("The number of digits was not found, using 2000 by default.")
		d = "2000"
	} else {
		d = os.Args[2]
	}

	// Parse the number of digits into an integer
	parsedD, err := strconv.Atoi(d)
	if err != nil || parsedD < 1500 {
		fmt.Println("The number of digits is not a valid number, using 2000. (must not be lower than 1500)")
		parsedD = 2000
	}

	// Read and parse π
	pi := digits.SerializeDigits(filePath, parsedD)
	parsedDigits := digits.ParseDigits(pi)

	// Create the image
	out, err := os.Create("digart.png")
	if err != nil {
		panic(err)
	}

	dim := parsedD * 2
	background := color.RGBA{A: 255}
	img := CreateImage(dim, dim, background)

	// Draw the elements
	digartImage.DrawCircle(*img, 350, 40, dim)
	digartImage.DrawData(*img, parsedDigits, 350, dim)

	err = png.Encode(out, img)
	if err != nil {
		panic(err)
	}

	fmt.Println("Image saved.")
}
