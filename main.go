package main

import (
	"./digits"
	digartImage "./image"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

// CreateImage returns the created image of `width` and `height` dimensions.
func CreateImage(width int, height int, background color.RGBA) *image.RGBA {
	rect := image.Rect(0, 0, width, height)
	img := image.NewRGBA(rect)
	draw.Draw(img, img.Bounds(), &image.Uniform{C: background}, image.ZP, draw.Src)

	return img
}

func main() {
	filePath := os.Args[1]
	if filePath == "" {
		fmt.Println("The file path was not found, using resources/pi by default.")
		filePath = "resources/pi.txt"
	}

	// Read and parse π
	d := 2000
	pi := digits.SerializeDigits(filePath, d)
	parsedDigits := digits.ParseDigits(pi)

	// Create the image
	out, err := os.Create("digart.png")
	if err != nil {
		panic(err)
	}

	dim := d*2
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
