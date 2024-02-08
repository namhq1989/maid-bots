package random

import (
	"fmt"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-telegram/bot"
	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/util/appcommand"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

type String struct {
	Message string
	Target  string
}

func (c String) Process(ctx *appcontext.AppContext) string {
	var result = "invalid target"

	switch c.Target {
	case appcommand.RandomStringTargets.Person:
		result = c.person(ctx)
	case appcommand.RandomStringTargets.Email:
		result = c.email(ctx)
	case appcommand.RandomStringTargets.Phone:
		result = c.phone(ctx)
	case appcommand.RandomStringTargets.Username:
		result = c.username(ctx)
	case appcommand.RandomStringTargets.Address:
		result = c.address(ctx)
	case appcommand.RandomStringTargets.LatLon:
		result = c.latlon(ctx)
	case appcommand.RandomStringTargets.Sentence:
		result = c.sentence(ctx)
	case appcommand.RandomStringTargets.Paragraph:
		result = c.paragraph(ctx)
	case appcommand.RandomStringTargets.Quote:
		result = c.quote(ctx)
	case appcommand.RandomStringTargets.UUID:
		result = c.uuid(ctx)
	case appcommand.RandomStringTargets.HexColor:
		result = c.hexColor(ctx)
	case appcommand.RandomStringTargets.RGBColor:
		result = c.rgbColor(ctx)
	case appcommand.RandomStringTargets.URL:
		result = c.url(ctx)
	case appcommand.RandomStringTargets.ImageURL:
		result = c.imageURL(ctx)
	case appcommand.RandomStringTargets.Domain:
		result = c.domain(ctx)
	case appcommand.RandomStringTargets.IPv4:
		result = c.ipv4(ctx)
	case appcommand.RandomStringTargets.IPv6:
		result = c.ipv6(ctx)
	case appcommand.RandomStringTargets.UserAgent:
		result = c.userAgent(ctx)
	case appcommand.RandomStringTargets.Date:
		result = c.date(ctx)
	case appcommand.RandomStringTargets.Timezone:
		result = c.timezone(ctx)
	case appcommand.RandomStringTargets.CreditCard:
		result = c.creditCard(ctx)
	case appcommand.RandomStringTargets.WalletAddress:
		result = c.walletAddress(ctx)
	case appcommand.RandomStringTargets.Pet:
		result = c.pet(ctx)
	case appcommand.RandomStringTargets.Emoji:
		result = c.emoji(ctx)
	}

	return bot.EscapeMarkdown(result)
}

func (String) person(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "person")
	defer span.Finish()

	v := gofakeit.Person()
	return fmt.Sprintf(
		"Name: %s %s\nAge: %d\nGender: %s\nEmail: %s\nPhone: %s\nAddress: %s",
		v.LastName,
		v.FirstName,
		gofakeit.IntRange(15, 90),
		v.Gender,
		v.Contact.Email,
		gofakeit.PhoneFormatted(),
		v.Address.Address,
	)
}

func (String) email(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "email")
	defer span.Finish()

	return gofakeit.Email()
}

func (String) phone(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "phone")
	defer span.Finish()

	return gofakeit.PhoneFormatted()
}

func (String) username(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "username")
	defer span.Finish()

	return strings.ToLower(gofakeit.Username())
}

func (String) address(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "address")
	defer span.Finish()

	return gofakeit.Address().Address
}

func (String) latlon(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "latlon")
	defer span.Finish()

	return fmt.Sprintf("%f,%f", gofakeit.Latitude(), gofakeit.Longitude())
}

func (String) sentence(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "sentence")
	defer span.Finish()

	return gofakeit.SentenceSimple()
}

func (String) paragraph(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "paragraph")
	defer span.Finish()

	return gofakeit.Paragraph(1, gofakeit.IntRange(2, 5), gofakeit.IntRange(8, 20), "\n")
}

func (String) quote(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "quote")
	defer span.Finish()

	return gofakeit.Quote()
}

func (String) uuid(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "uuid")
	defer span.Finish()

	return gofakeit.UUID()
}

func (String) hexColor(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "hexColor")
	defer span.Finish()

	return gofakeit.HexColor()
}

func (String) rgbColor(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "rgbColor")
	defer span.Finish()

	v := gofakeit.RGBColor()
	return fmt.Sprintf("rgb(%d, %d, %d)", v[0], v[1], v[2])
}

func (String) url(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "url")
	defer span.Finish()

	return gofakeit.URL()
}

func (String) imageURL(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "imageURL")
	defer span.Finish()

	return gofakeit.ImageURL(gofakeit.IntRange(100, 300), gofakeit.IntRange(100, 300))
}

func (String) domain(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "domain")
	defer span.Finish()

	return gofakeit.DomainName()
}

func (String) ipv4(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "ipv4")
	defer span.Finish()

	return gofakeit.IPv4Address()
}

func (String) ipv6(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "ipv6")
	defer span.Finish()

	return gofakeit.IPv6Address()
}

func (String) userAgent(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "userAgent")
	defer span.Finish()

	return gofakeit.UserAgent()
}

func (String) date(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "date")
	defer span.Finish()

	return gofakeit.Date().String()
}

func (String) timezone(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "timezone")
	defer span.Finish()

	return gofakeit.TimeZoneFull()
}

func (c String) creditCard(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "creditCard")
	defer span.Finish()

	v := gofakeit.CreditCard()
	return fmt.Sprintf(
		"Card type: %s \nCard holder name: %s\nCard number: %s\nExpiry date: %s\nCVV: %s",
		v.Type,
		gofakeit.Name(),
		c.formatCardNumber(v.Number),
		v.Exp,
		v.Cvv,
	)
}

func (String) formatCardNumber(cardNumber string) string {
	var builder strings.Builder
	for i := 0; i < len(cardNumber); i += 4 {
		endIndex := i + 4
		if endIndex > len(cardNumber) {
			endIndex = len(cardNumber)
		}
		builder.WriteString(cardNumber[i:endIndex])
		if endIndex < len(cardNumber) {
			builder.WriteString(" ")
		}
	}
	return builder.String()
}

func (String) walletAddress(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "walletAddress")
	defer span.Finish()

	return gofakeit.BitcoinAddress()
}

func (String) pet(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "pet")
	defer span.Finish()

	return gofakeit.PetName()
}

func (String) emoji(ctx *appcontext.AppContext) string {
	span := sentryio.NewSpan(ctx.Context, "emoji")
	defer span.Finish()

	return gofakeit.Emoji()
}
