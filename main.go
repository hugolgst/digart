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
	// Read and parse π
	pi := digits.SerializeDigits("resources/pi.txt", 2300)
	parsedDigits := digits.ParseDigits(pi)

	// Create the image
	out, err := os.Create("digart.png")
	if err != nil {
		panic(err)
	}

	width, height := 4500, 4500
	background := color.RGBA{A: 255}
	img := CreateImage(width, height, background)

	// Draw the elements
	digartImage.DrawCircle(*img, 350, 40, width, height)
	digartImage.DrawData(*img, parsedDigits, width, height)

	err = png.Encode(out, img)
	if err != nil {
		panic(err)
	}

	fmt.Println("Image saved.")
}