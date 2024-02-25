package config

import "os"

const (
	envRelease     = "release"
	envDevelopment = "development"
)

func IsRelease() bool {
	return os.Getenv("ENVIRONMENT") == envRelease
}

func IsDevelopment() bool {
	return os.Getenv("ENVIRONMENT") == envDevelopment
}
