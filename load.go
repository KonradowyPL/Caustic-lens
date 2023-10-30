package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)



func getImageFromFilePath(filePath string) image.Image {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return img
}

func spawn() {
	path, _ :=filepath.Abs("example.png")
	img := getImageFromFilePath(path)
	size := img.Bounds().Size()
	screenX = size.X
	screenY = size.Y

	fmt.Println(size)

	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			pixel := img.At(x,y)
			R,_,_,_ := pixel.RGBA()
			
			if (rand.Float64() < float64(R)/65536 ) {
				particles = append(particles, Particle{float64(x),float64(y),0,0})
			}
		}
	}
}
