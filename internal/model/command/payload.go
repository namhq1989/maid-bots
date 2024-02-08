package modelcommand

type Payload struct {
	Platform string
	ChatID   string
	Message  string
	User     User
}
