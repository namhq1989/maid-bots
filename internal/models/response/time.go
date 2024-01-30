package modelresponse

import (
	"time"

	"github.com/goccy/go-json"
)

const (
	formatLayoutFull = "2006-01-02T15:04:05.000Z"
)

type TimeResponse struct {
	Time time.Time
}

func (t *TimeResponse) UnmarshalJSON(b []byte) error {
	if string(b) == "" || string(b) == "\"\"" {
		return nil
	}
	return json.Unmarshal(b, &t.Time)
}

func (t *TimeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.FormatISODate())
}

func (t *TimeResponse) FormatISODate() string {
	if t == nil || t.Time.IsZero() {
		return ""
	}
	return t.Time.Format(formatLayoutFull)
}

func (t *TimeResponse) FormatYYYYMMDD() string {
	if t == nil || t.Time.IsZero() {
		return ""
	}
	return t.Time.Format("2006-01-02")
}

func NewTimeResponse(t time.Time) *TimeResponse {
	return &TimeResponse{Time: t}
}
