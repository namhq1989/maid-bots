package pagetoken

import (
	"encoding/base64"
	"time"

	"github.com/goccy/go-json"
)

// PageToken ...
type PageToken struct {
	Page      int
	Timestamp time.Time
}

func getDefaultPageToken() (response PageToken) {
	response.Page = 0
	response.Timestamp = time.Now()
	return response
}

// Encode encode next page token for api response
func Encode(page int, timestamp time.Time) string {
	tokenData := PageToken{
		Page:      page,
		Timestamp: timestamp,
	}
	tokenString, _ := json.Marshal(tokenData)
	encodedString := base64.StdEncoding.EncodeToString(tokenString)
	return encodedString
}

// Decode decode page token from query
func Decode(encodedString string) PageToken {
	if encodedString == "" {
		return getDefaultPageToken()
	}

	// Decode string
	decoded, err := base64.StdEncoding.DecodeString(encodedString)

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

// WithPage generate next page token using "page"
func WithPage(page int) string {
	return Encode(page, time.Now())
}

// WithTimestamp generate next page token using "timestamp"
func WithTimestamp(timestamp time.Time) string {
	return Encode(0, timestamp)
}
