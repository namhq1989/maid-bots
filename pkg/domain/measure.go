package domain

import (
	"net/http"
	"time"
)

func MeasureResponseTime(url string) (int64, error) {
	startTime := time.Now()

	r, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer func() { _ = r.Body.Close() }()

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	return duration.Milliseconds(), nil
}
