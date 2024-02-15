package mongodb

import (
	"time"

	"github.com/namhq1989/maid-bots/util/appcontext"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

// RunTransaction ...
func RunTransaction(ctx *appcontext.AppContext, fn func(sessionContext mongo.SessionContext) error) error {
	ctx.Logger.Text("start a new MongoDB transaction")

	client := db.Client()
	session, err := client.StartSession()
	if err != nil {
		return err
	}

	var (
		wc = writeconcern.Majority()
		rc = readconcern.Snapshot()
		rf = readpref.Primary()

		maxCommitTime = time.Minute
		txnOpts       = options.Transaction().SetReadPreference(rf).
				SetWriteConcern(wc).SetReadConcern(rc).SetMaxCommitTime(&maxCommitTime)
	)
	defer session.EndSession(ctx.Context)

	if err = mongo.WithSession(ctx.Context, session, func(sessionContext mongo.SessionContext) (errTransaction error) {
		if errTransaction = session.StartTransaction(txnOpts); errTransaction != nil {
			return errTransaction
		}

		// handler
		if errTransaction = fn(sessionContext); errTransaction != nil {
			return errTransaction
		}

		// commit
		return session.CommitTransaction(sessionContext)
	}); err != nil {
		ctx.Logger.Error("transaction processed unsuccessfully, rollback", err, appcontext.Fields{})

		// rollback
		if abortErr := session.AbortTransaction(ctx.Context); abortErr != nil {
			return abortErr
		}
		return err
	}

	ctx.Logger.Text("transaction processed successfully")
	return nil
}
