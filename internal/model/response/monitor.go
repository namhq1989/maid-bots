package modelresponse

type Monitor struct {
	Code      string        `json:"code"`
	Type      string        `json:"type"`
	Target    string        `json:"target"`
	Interval  int           `json:"interval"`
	CreatedAt *TimeResponse `json:"createdAt"`
}
