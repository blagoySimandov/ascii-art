package main

// Importing the needed packages
import (
	"fmt"
	"image"
	_ "image/jpeg"
	"io"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

func main() {
	//Creating a slice out of all the ascii characters sorted by character density
	ascii := strings.Split("`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$", "")
	file, err := os.Open("images/ascii-pineapple.jpg")
	if err != nil {
		fmt.Println("Error: Image could not be opened")
		os.Exit(1)
	}

	//Create a 2D array out of the pixels in the image
	pixels, err := getPixels(file, 280)
	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}
	//creating a scale for how much of each gray scale color (0-256) should corespond to one ascii character
	scale := 256 / len(ascii)
	var picture [][]string
	var arr []string
	//Iterating through the 2d array pixels and mapping  the position of each ascii character coresponding to the grayscale value of the pixel in the original image
	for _, x := range pixels {
		for _, y := range x {
			if (y / scale) > len(ascii)-1 {
				arr = append(arr, string(ascii[len(ascii)-1]))
			} else {
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

//Function that creates a 2d array out of the grayscale values of the pixels in the original image
func getPixels(file io.Reader, size uint) ([][]int, error) {
	imageDecoded, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}
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

//function that takes in RGBA values and outputs the average value (aka the grayscale value)
func rgbaToAverageMapping(r uint32, g uint32, b uint32, a uint32) int {
	return (int(r>>8) + int(g>>8) + int(b>>8)) / 3
}
