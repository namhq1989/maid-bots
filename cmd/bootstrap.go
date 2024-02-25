package main

import (
	"crypto/subtle"
	"github.com/namhq1989/maid-bots/internal/job"

	"github.com/labstack/echo/v4/middleware"
	"github.com/namhq1989/maid-bots/content"
	"github.com/namhq1989/maid-bots/pkg/queue"
	"github.com/namhq1989/maid-bots/platform/telegram"
	"github.com/namhq1989/maid-bots/platform/web/route"

	"github.com/labstack/echo/v4"
	"github.com/namhq1989/maid-bots/config"
	"github.com/namhq1989/maid-bots/pkg/logger"
	"github.com/namhq1989/maid-bots/pkg/mongodb"
	"github.com/namhq1989/maid-bots/pkg/redis"
	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/pkg/sso"
	"github.com/namhq1989/maid-bots/util/jwt"
)

func bootstrap(e *echo.Echo) {
	// config
	config.Init()
	cfg := config.GetENV()

	// logger
	logger.Init(cfg.Environment)

	// jwt
	jwt.Init(cfg.AuthSecret)

	// database
	mongodb.Connect(cfg.MongoURL, cfg.MongoDBName)
	redis.Connect(cfg.RedisURL)

	// sentry
	sentryio.Init(e, cfg.SentryDSN, cfg.SentryMachine, cfg.Environment)

	// sso
	sso.Init(sso.Providers{
		GoogleClientID:     cfg.SSOGoogleClientID,
		GoogleClientSecret: cfg.SSOGoogleClientSecret,
		GitHubClientID:     cfg.SSOGitHubClientID,
		GitHubClientSecret: cfg.SSOGitHubClientSecret,
	})

	// queue
	queue.Init(cfg.RedisURL, cfg.QueueConcurrency)
	e.Any("/q/*", echo.WrapHandler(queue.Dashboard(cfg.RedisURL)), middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if !config.IsRelease() {
			return true, nil
		}

		return subtle.ConstantTimeCompare([]byte(username), []byte(cfg.QueueUsername)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(cfg.QueuePassword)) == 1, nil
	}))

	// job
	job.Init()

	// load content
	content.Load()

	// platforms
	// web
	route.Init(e)

	// telegram
	telegram.Init(cfg.TelegramEnabled, cfg.TelegramBotToken)
}
