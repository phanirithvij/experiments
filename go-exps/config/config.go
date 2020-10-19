package config

import "time"

// https://stackoverflow.com/a/62733180/8608146

// Version The version of the build
var Version string

// BuildTime Built on
var BuildTime string

// CommitID The commit ID of the build
var CommitID string

// BuildTimeDate returns the date representation of the build time
func BuildTimeDate() (time.Time, error) {
	return time.Parse(time.RFC1123Z, BuildTime)
}
