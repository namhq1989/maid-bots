package mongodb

// // RunTransaction ...
// func RunTransaction(ctx appcontext.AppContext, fn func(sessionContext mongo.SessionContext) error) error {
// 	ctx.Logger.Text("start new MongoDB transaction")
//
// 	client := db.Client()
// 	session, err := client.StartSession()
// 	if err != nil {
// 		return err
// 	}
//
// 	wc := writeconcern.Majority()
// 	rc := readconcern.Snapshot()
// 	rf := readpref.Primary()
//
// 	maxCommitTime := time.Minute
// 	txnOpts := options.Transaction().SetReadPreference(rf).
// 		SetWriteConcern(wc).SetReadConcern(rc).SetMaxCommitTime(&maxCommitTime)
//
// 	defer session.EndSession(ctx.Context)
// 	err = mongo.WithSession(ctx.Context, session, func(sessionContext mongo.SessionContext) error {
// 		var errTransaction error
// 		if errTransaction = session.StartTransaction(txnOpts); errTransaction != nil {
// 			return errTransaction
// 		}
//
// 		// Handle func
// 		if errTransaction = fn(sessionContext); errTransaction != nil {
// 			return errTransaction
// 		}
//
// 		// Commit
// 		return session.CommitTransaction(sessionContext)
// 	})
//
// 	if err != nil {
// 		ctx.Logger.Error("processed transaction failure, rollback", err, appcontext.Fields{})
//
// 		// Rollback
// 		if abortErr := session.AbortTransaction(ctx.Context); abortErr != nil {
// 			return abortErr
// 		}
// 		return err
// 	}
// 	return nil
// }
