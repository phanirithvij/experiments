package config

import (
	"fmt"
	"log"
	"time"
)

// https://stackoverflow.com/a/62733180/8608146

// Version The version of the build
var Version string

// BuildTime Built on
var BuildTime string

// CommitID The commit ID of the build
var CommitID string

// Architecture of the OS 32bit or 64 bit
var Architecture string

// Platform i.e. the OS
var Platform string

// BuildTimeDate returns the date representation of the build time
func BuildTimeDate() (time.Time, error) {
	return time.Parse(time.RFC1123Z, BuildTime)
}

// LogVersionInfo logs the build and version information
func LogVersionInfo() {
	fmt.Println("Platform", Platform)
	fmt.Println("Architecture", Architecture)
	fmt.Println("Version", Version)
	fmt.Println("Build commit", CommitID)
	date, err := BuildTimeDate()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Built on", date)
}
