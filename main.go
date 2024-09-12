package main

import (
	"fmt"
	"log"
	"os"
)

const fileName = "test.ppm"
const magic = "P3"
const imageWidth = 256
const imageHeight = 256
const maxVal = 255

func main() {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write([]byte(fmt.Sprintf("%s\n%d %d\n%d\n", magic, imageWidth, imageHeight, maxVal)))
	if err != nil {
		log.Fatal("Cannot write header", err)
	}

	for i := range imageWidth {
		log.Println("Scanline remaining: ", imageHeight-i)
		for j := 0; j < imageHeight; j++ {
			// Calculate red and green values based on position
			pixelColor := Vector{float64(j) / float64(imageHeight-1), float64(i) / float64(imageWidth-1), 0}
			_, err := WriteColor(file, &pixelColor, float64(maxVal))
			if err != nil {
				log.Fatal("Cannot write pixel", err)
			}
		}
	}

}
