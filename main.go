package main

import (
	"./digits"
	digartImage "./image"
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"syscall/js"
)

// CreateImage returns the created image of `width` and `height` dimensions.
func CreateImage(width int, height int, background color.RGBA) *image.RGBA {
	rect := image.Rect(0, 0, width, height)
	img := image.NewRGBA(rect)
	draw.Draw(img, img.Bounds(), &image.Uniform{C: background}, image.ZP, draw.Src)

	return img
}

func BuildImage(number string) string {
	// Read and parse the number
	parsedDigits := digits.ParseDigits(number)

	out := bytes.NewBuffer(nil)

	fmt.Println(len(number))

	dim := len(number) * 2
	background := color.RGBA{A: 255}
	img := CreateImage(dim, dim, background)

	// Draw the elements
	digartImage.DrawCircle(*img, 350, 40, dim)
	digartImage.DrawData(*img, parsedDigits, 350, dim)

	if err := png.Encode(out, img); err != nil {
		panic(err)
	}

	b64 := base64.StdEncoding.EncodeToString(out.Bytes())
	return fmt.Sprint(`data:image/png;base64,`, b64)
}

func main() {
	var number string

	c := make(chan struct{}, 0)

	runAction := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		numberInput := js.Global().Get("document").
			Call("getElementById", "number").
			Get("value")

		number = numberInput.String()

		imageSrc := BuildImage(number)
		js.Global().Get("document").Call("getElementById", "image").Set("src", imageSrc)

		return nil
	})

	js.Global().Get("document").
		Call("getElementById", "submit").
		Call("addEventListener", "click", runAction)

	<-c
}
