package genetics

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func (d *DNA) readImage(pathToImage string) {
	// TODO: should be configurable by user
	// Register image format
	image.RegisterFormat("jpg", "jpg", jpeg.Decode, jpeg.DecodeConfig)

	file, err := os.Open(pathToImage)

	if err != nil {
		log.Fatal("Can't read image. Check if given path to file is correct.")
	}

	// Automaticly closes file descriptor when it's needed anymore
	defer file.Close()

	imageData, _, err := image.Decode(file)

	if err != nil {
		log.Fatal("Can't decode image. Check if given file has correct format.")
	}

	d.originalImage = imageData
}

func (d *DNA) saveImage(imageToSave image.RGBA, imgName string) {
	file, err := os.Create(imgName)

	if err != nil {
		log.Fatal("Problem with saving image")
	}

	defer file.Close()

	png.Encode(file, &imageToSave)
}

func readImage(pathToImage string) image.Image {
	// TODO: should be configurable by user
	// Register image format
	image.RegisterFormat("jpg", "jpg", jpeg.Decode, jpeg.DecodeConfig)

	file, err := os.Open(pathToImage)

	if err != nil {
		log.Fatal("Can't read image. Check if given path to file is correct.")
	}

	// Automaticly closes file descriptor when it's needed anymore
	defer file.Close()

	imageData, _, err := image.Decode(file)

	if err != nil {
		log.Fatal("Can't decode image. Check if given file has correct format.")
	}

	return imageData
}
