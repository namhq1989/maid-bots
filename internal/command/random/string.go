package random

import "github.com/namhq1989/maid-bots/util/appcontext"

type String struct {
	Message string
}

func (c String) Process(_ *appcontext.AppContext) string {
	return ""
}
