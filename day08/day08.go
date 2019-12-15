package main

import (
	"advent-2019/utils"
	"fmt"
	"math"
)

func main() {
	imageData := utils.ReadFile("./day08/input.txt")
	image := NewImage(imageData, 25, 6)

	fmt.Println("----- Part 1 -----")
	layer := image.GetLayerWithFewest(0)
	fmt.Printf("For the layer with the fewest 0s, multiplying the count of 1s by the count of 2s gives us: %d\n\n", layer.counts[1] * layer.counts[2])

	fmt.Println("----- Part 2 -----")
	fmt.Printf("Here's the rendered image:\n\n%s\n\n", image.Render())
}

type Image struct {
	layers []Layer
	width, height int
}

type Layer struct {
	data []int
	counts map[int]int
}

func NewImage(data string, width, height int) *Image {
	image := new(Image)
	image.width = width
	image.height = height

	layerSize := width * height
	totalLayers := len(data) / layerSize

	for layerIndex := 0; layerIndex < totalLayers; layerIndex++ {
		layer := new(Layer)
		layer.data = make([]int, width*height)
		layer.counts = createLayerCount()
		for i, r := range substr(data, layerIndex*layerSize, layerSize) {
			digit := int(r - '0')
			layer.data[i] = digit
			layer.counts[digit]++
		}

		image.layers = append(image.layers, *layer)
	}

	return image
}

func createLayerCount() map[int]int {
	res := make(map[int]int)

	for i := 0; i < 10; i++ {
		res[i] = 0
	}

	return res
}

func (i *Image) GetLayerWithFewest(digit int) *Layer {
	var layer Layer
	digits := math.MaxInt32

	for _, l := range i.layers {
		if l.counts[digit] < digits {
			digits = l.counts[digit]
			layer = l
		}
	}

	return &layer
}

func (i *Image) mergeLayers() *Layer {
	// Create a new layer which is transparent by default
	result := new(Layer)
	result.data = make([]int, i.width*i.height)
	for i := range result.data {
		result.data[i] = 2
	}

	// Iterate through each layer, applying the pixel value only if that pixel was previously transparent
	for _, layer := range i.layers {
		for y := 0; y < i.height; y++ {
			for x := 0; x < i.width; x++ {
				if result.data[y*i.width+x] == 2 {
					result.data[y*i.width+x] = layer.data[y*i.width+x]
				}
			}
		}
	}

	return result
}

func (i *Image) Render() string {
	merged := i.mergeLayers()
	result := ""

	for y := 0; y < i.height; y++ {
		for x := 0; x < i.width; x++ {
			switch merged.data[y*i.width+x] {
			case 0:
				result += " "
			case 1:
				result += "X"
			}
		}
		result += "\n"
	}

	return result
}

func substr(str string, start int, length int) string {
	return str[start:start+length]
}
