package modelresponse

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func formatReadableInt(v int64) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%d", v)
}

func formatReadableFloat64(v float64, precision int) string {
	p := message.NewPrinter(language.English)
	f := fmt.Sprintf("%%.%df", precision)
	return p.Sprintf(f, v)
}
