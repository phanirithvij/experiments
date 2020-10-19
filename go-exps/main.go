package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/phanirithvij/experiments/go-exps/config"
)

//go:generate go-bindata assets/

func main() {
	log.SetFlags(0)
	log.Println("Version", config.Version)
	log.Println("Build commit", config.CommitID)
	date, err := config.BuildTimeDate()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Built on", date)
	data, err := Asset("assets/square2.png")
	if err != nil {
		log.Fatalln(err)
	}
	_, _, err = image.Decode(bytes.NewReader(data))
	// decodedImage, _, _ := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatalln(err)
	}
	// palette := vibrant.NewPaletteBuilder(decodedImage).Generate()
	// Iterate over the swatches in the palette...
	// for _, swatch := range palette.Swatches() {
	// 	// log.Printf("Swatch has color %v and population %d\n", swatch.RGBAInt(), swatch.Population())
	// }
	// for _, target := range palette.Targets() {
	// 	// pall := palette.SwatchForTarget(target)
	// 	// log.Println(target)
	// 	// Do something with the swatch for a given target...
	// }
}
