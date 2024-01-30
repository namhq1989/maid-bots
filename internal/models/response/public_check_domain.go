package modelresponse

type PublicCheckDomain struct {
	IsHTTPS          bool          `json:"isHttps"`
	IsUp             bool          `json:"isUp"`
	ResponseTimeInMS int64         `json:"responseTime"`
	DomainName       string        `json:"domainName"`
	Scheme           string        `json:"scheme"`
	ExpireAt         *TimeResponse `json:"expireAt"`
	Issuer           string        `json:"issuer"`
	IPResolves       []string      `json:"ipResolves"`
}
