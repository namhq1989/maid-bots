package appcommand

var RandomTypes = struct {
	Number ArgumentType
	String ArgumentType
}{
	Number: ArgumentType{
		Name:        "number",
		Description: "Random a number",
	},
	String: ArgumentType{
		Name:        "string",
		Description: "Random a string",
	},
}

var RandomNumberParameters = struct {
	Type   string
	Format string
	Min    string
	Max    string
	Count  string
	Unique string
}{
	Type:   "type",
	Format: "format",
	Min:    "min",
	Max:    "max",
	Count:  "count",
	Unique: "unique",
}

var RandomStringParameters = struct {
	Type  string
	Value string
}{
	Type:  "string",
	Value: "value",
}

var RandomStringValues = struct {
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
