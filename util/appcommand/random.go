package appcommand

import "slices"

var RandomTargets = struct {
	Number ArgumentTarget
	String ArgumentTarget
}{
	Number: ArgumentTarget{
		Name:        "number",
		Description: "Random a number",
	},
	String: ArgumentTarget{
		Name:        "string",
		Description: "Random a string",
	},
}

var RandomNumberParameters = struct {
	Type   string
	Min    string
	Max    string
	Count  string
	Unique string
}{
	Type:   "type",
	Min:    "min",
	Max:    "max",
	Count:  "count",
	Unique: "unique",
}

var RandomStringTargets = struct {
	Person        string
	Email         string
	Phone         string
	Username      string
	Address       string
	LatLon        string
	Sentence      string
	Paragraph     string
	Quote         string
	UUID          string
	HexColor      string
	RGBColor      string
	URL           string
	ImageURL      string
	Domain        string
	IPv4          string
	IPv6          string
	UserAgent     string
	Date          string
	Timezone      string
	CreditCard    string
	WalletAddress string
	Pet           string
	Emoji         string
}{
	Person:        "person",
	Email:         "email",
	Phone:         "phone",
	Username:      "username",
	Address:       "address",
	LatLon:        "latlon",
	Sentence:      "sentence",
	Paragraph:     "paragraph",
	Quote:         "quote",
	UUID:          "uuid",
	HexColor:      "hexcolor",
	RGBColor:      "rgbcolor",
	URL:           "url",
	ImageURL:      "imageurl",
	Domain:        "domain",
	IPv4:          "ipv4",
	IPv6:          "ipv6",
	UserAgent:     "ua",
	Date:          "date",
	Timezone:      "timezone",
	CreditCard:    "creditcard",
	WalletAddress: "walletaddress",
	Pet:           "pet",
	Emoji:         "emoji",
}

var RandomStringTargetsArray = []string{
	RandomStringTargets.Person,
	RandomStringTargets.Email,
	RandomStringTargets.Phone,
	RandomStringTargets.Username,
	RandomStringTargets.Address,
	RandomStringTargets.LatLon,
	RandomStringTargets.Sentence,
	RandomStringTargets.Paragraph,
	RandomStringTargets.Quote,
	RandomStringTargets.UUID,
	RandomStringTargets.HexColor,
	RandomStringTargets.RGBColor,
	RandomStringTargets.URL,
	RandomStringTargets.ImageURL,
	RandomStringTargets.Domain,
	RandomStringTargets.IPv4,
	RandomStringTargets.IPv6,
	RandomStringTargets.UserAgent,
	RandomStringTargets.Date,
	RandomStringTargets.Timezone,
	RandomStringTargets.CreditCard,
	RandomStringTargets.WalletAddress,
	RandomStringTargets.Pet,
	RandomStringTargets.Emoji,
}

func IsRandomStringTargetValid(t string) bool {
	return slices.Contains(RandomStringTargetsArray, t)
}
