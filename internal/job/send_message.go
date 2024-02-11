package job

import (
	"github.com/namhq1989/maid-bots/internal/service"
	"github.com/namhq1989/maid-bots/platform/telegram"
	"github.com/namhq1989/maid-bots/util/appcontext"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func sendMessage(ctx *appcontext.AppContext, ownerID primitive.ObjectID, targetCode, message string) {
	var (
		hcrSvc = service.HealthCheckRecord{}
	)

	// get latest history with status "down" of the target using target code
	isDown, _ := hcrSvc.IsTargetDownRecently(ctx, targetCode)
	if isDown {
		// if found, skip
		return
	}

	// send message
	var (
		userSvc = service.User{}
	)

	// find user with owner id
	user, _ := userSvc.FindByID(ctx, ownerID)
	if user == nil {
		// return if not found
		ctx.Logger.Error("user not found", nil, appcontext.Fields{"ownerID": ownerID.Hex()})
		return
	}

	// Telegram
	if user.Telegram != nil {
		telegram.SendMessage(ctx, user.Telegram.RoomID, message)
	}

	// Discord

	// Slack
}
