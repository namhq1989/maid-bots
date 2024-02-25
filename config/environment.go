package config

import "os"

const (
	envRelease     = "release"
	envDevelopment = "development"
)

var actualEnvironment = os.Getenv("ENVIRONMENT")

func IsRelease() bool {
	return actualEnvironment == envRelease
}

func IsDevelopment() bool {
	return actualEnvironment == envDevelopment
}
