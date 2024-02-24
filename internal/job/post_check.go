package job

import (
	"fmt"
	"math"
	"time"

	"github.com/namhq1989/maid-bots/internal/service"
	"github.com/namhq1989/maid-bots/pkg/mongodb"
	"github.com/namhq1989/maid-bots/platform/telegram"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

const (
	last15Minutes = time.Minute * 15

	successMessage = "[resolved]  ✅✅✅  %s %s is UP  ✅✅✅  "
	failedMessage  = "[incident]  🆘🆘🆘  %s %s is DOWN - %s  🆘🆘🆘  "
)

func sendMessage(ctx *appcontext.AppContext, doc mongodb.HealthCheckRecord) {
	var (
		hcrSvc  = service.HealthCheckRecord{}
		message string
	)

	// get latest record
	recentRecord, err := hcrSvc.GetRecentRecordOfMonitor(ctx, doc.Code)
	if err != nil {
		ctx.Logger.Error("error when get recent record", err, appcontext.Fields{"targetCode": doc.Code})
		return
	}

	// if not found
	if recentRecord == nil {
		// if up then do nothing
		if doc.Status == mongodb.HealthCheckRecordStatusUp {
			return
		}

		// if failed, set message
		message = fmt.Sprintf(failedMessage, doc.Type, doc.Target, doc.Description)
	} else {
		// compare the status of recent and current record
		if doc.Status == mongodb.HealthCheckRecordStatusUp && recentRecord.Status == mongodb.HealthCheckRecordStatusDown {
			message = fmt.Sprintf(successMessage, doc.Type, doc.Target)
		} else if doc.Status == mongodb.HealthCheckRecordStatusDown && doc.CreatedAt.After(recentRecord.CreatedAt.Add(last15Minutes)) {
			message = fmt.Sprintf(failedMessage, doc.Type, doc.Target, doc.Description)
		} else {
			return
		}
	}

	// send message
	var (
		userSvc = service.User{}
	)

	// find user with owner id
	user, _ := userSvc.FindByID(ctx, doc.Owner)
	if user == nil {
		// return if not found
		ctx.Logger.Error("user not found", nil, appcontext.Fields{"ownerID": doc.Owner.Hex()})
		return
	}

	// Telegram
	if user.Telegram != nil {
		telegram.SendMessage(ctx, user.Telegram.RoomID, message)
	}

	// Discord

	// Slack
}

const (
	sslExpiryRemainingDaysCheck = 70
	sslExpirationMessage        = "[attention]  🆘🆘🆘  the SSL certificate for %s will expire in the next %d days, please take necessary actions to renew it  🆘🆘🆘  "
)

func checkAndSendSSLExpirationMessage(ctx *appcontext.AppContext, doc mongodb.HealthCheckRecord, expireAt time.Time) {
	var (
		now                = time.Now()
		daysUntilSSLExpiry = expireAt.Sub(now).Hours() / 24
	)

	if daysUntilSSLExpiry > sslExpiryRemainingDaysCheck {
		return
	}

	// count if there is already a heath check record for today, return
	var (
		hcrSvc = service.HealthCheckRecord{}
	)

	if isChecked := hcrSvc.IsTargetCheckedToday(ctx, doc.Code); isChecked {
		return
	}

	// send message
	var (
		userSvc = service.User{}
	)

	// find user with owner id
	user, _ := userSvc.FindByID(ctx, doc.Owner)
	if user == nil {
		// return if not found
		ctx.Logger.Error("user not found", nil, appcontext.Fields{"ownerID": doc.Owner.Hex()})
		return
	}

	message := fmt.Sprintf(sslExpirationMessage, doc.Target, int(math.Round(daysUntilSSLExpiry)))

	// Telegram
	if user.Telegram != nil {
		telegram.SendMessage(ctx, user.Telegram.RoomID, message)
	}

	// Discord

	// Slack
}
