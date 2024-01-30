package echocontext

import (
	"encoding/base64"
	"time"

	"github.com/goccy/go-json"
)

// PageToken ...
type PageToken struct {
	Page      int64
	Timestamp time.Time
	Order     int
}

func getDefaultPageToken() (response PageToken) {
	response.Page = 0
	response.Timestamp = time.Now()
	return response
}

// PageTokenDecode decode page token from query
func PageTokenDecode(s string) PageToken {
	if s == "" {
		return getDefaultPageToken()
	}

	// Decode string
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return getDefaultPageToken()
	}

	// Parse token
	var pageToken PageToken
	err = json.Unmarshal(decoded, &pageToken)
	if err != nil {
		return getDefaultPageToken()
	}

	return pageToken
}
