package cmd

import (
	"bytes"
	"fmt"
	"image"

	// Png needs to be imported to decode png images
	_ "image/png"
	"log"
	"os"

	"github.com/RobCherry/vibrant"
	"github.com/phanirithvij/experiments/goexps/experiments"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the code",
	Long:  `Runs the code and prints colors or somestuff`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := experiments.Asset("assets/square2.png")
		if err != nil {
			log.Fatalln(err)
		}
		// _, _, err = image.Decode(bytes.NewReader(data))
		decodedImage, _, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			log.Fatalln(err)
		}
		palette := vibrant.NewPaletteBuilder(decodedImage).Generate()
		// Iterate over the swatches in the palette...
		// for _, swatch := range palette.Swatches() {
		// 	log.Printf("Swatch has color %v and population %d\n", swatch.RGBAInt(), swatch.Population())
		// }
		for _, target := range palette.Targets() {
			// pall := palette.SwatchForTarget(target)
			log.Println(target)
			// Do something with the swatch for a given target...
		}
	},
}

// Execute executes cobra
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
