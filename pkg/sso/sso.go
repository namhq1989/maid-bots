package sso

import (
	"errors"
)

type Providers struct {
	// Google
	GoogleClientID     string
	GoogleClientSecret string

	// GitHub
	GitHubClientID     string
	GitHubClientSecret string
}

// hold providers information
var providers Providers

func Init(p Providers) {
	// check
	if p.GoogleClientID == "" || p.GitHubClientID == "" {
		panic(errors.New("invalid SSO: no provider provided"))
	}

	// assign
	providers = p
}
