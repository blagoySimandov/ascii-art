package main

import (
	"fmt"
	"image"

	"github.com/nfnt/resize"

	_ "image/jpeg"
	"io"
	"os"
	"strings"
)

/*
type Pixel struct {
	R int
	G int
	B int
}
*/
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	ascii := strings.Split("`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$", "")
	var arr []string
	file, err := os.Open("images/ascii-pineapple.jpg")

	if err != nil {
		fmt.Println("Error: Image could not be opened")
		os.Exit(1)
	}
	pixels, err := getPixels(file, 280)
	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}
	scale := 256 / len(ascii)
	var picture [][]string
	fmt.Println(len(pixels))
	fmt.Println(len(pixels[0]))
	fmt.Println(len(pixels[1]))
	for _, x := range pixels {
		for _, y := range x {
			if (y / scale) > len(ascii)-1 {
				arr = append(arr, string(ascii[len(ascii)-1]))

			} else {
				fmt.Println(y / scale)
				arr = append(arr, string(ascii[y/scale]))
			}
		}
		picture = append(picture, arr)
		arr = nil
	}
	for _, row := range picture {
		fmt.Println(row)
	}
}
func getPixels(file io.Reader, size uint) ([][]int, error) {
	imageDecoded, _, err := image.Decode(file)
	check(err)
	img := resize.Resize(size, 0, imageDecoded, resize.Lanczos2)
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	var pixels [][]int
	for y := 0; y < h; y++ {
		var row []int
		for x := 0; x < w; x++ {
			row = append(row, rgbaToAverageMapping(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}
	return pixels, nil
}
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) int {
	//Pixel{int(r / 257), int(g / 257), int(b / 257)}
	return (int(r/257) + int(g/257) + int(b/257)) / 3
}

func rgbaToAverageMapping(r uint32, g uint32, b uint32, a uint32) int {
	//Pixel{int(r / 257), int(g / 257), int(b / 257)}
	return (int(r/257) + int(g/257) + int(b/257)) / 3
}
