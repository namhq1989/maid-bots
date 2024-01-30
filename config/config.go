package config

import "fmt"

type ENV struct {
	AppName     string
	Environment string
	Debug       bool

	// Single Sign On
	SSOGoogleClientID     string
	SSOGoogleClientSecret string
	SSOGoogleCallbackURL  string
	SSOGitHubClientID     string
	SSOGitHubClientSecret string
	SSOGitHubCallbackURL  string

	// Authentication
	AuthSecret string

	// MongoDB
	MongoURL    string
	MongoDBName string

	// Redis
	RedisURL string

	// Queue
	QueueUsername    string
	QueuePassword    string
	QueueConcurrency int

	// Sentry
	SentryDSN     string
	SentryMachine string

	// Telegram
	TelegramEnabled  bool
	TelegramBotToken string
}

var env ENV

func GetENV() ENV {
	return env
}

func Init() {
	env = ENV{
		AppName:     getEnvStr("APP_NAME"),
		Environment: getEnvStr("ENVIRONMENT"),
		Debug:       getEnvBool("DEBUG"),

		SSOGoogleClientID:     getEnvStr("SSO_GOOGLE_CLIENT_ID"),
		SSOGoogleClientSecret: getEnvStr("SSO_GOOGLE_CLIENT_SECRET"),
		SSOGitHubClientID:     getEnvStr("SSO_GITHUB_CLIENT_ID"),
		SSOGitHubClientSecret: getEnvStr("SSO_GITHUB_CLIENT_SECRET"),

		AuthSecret: getEnvStr("AUTH_SECRET"),

		MongoURL:    getEnvStr("MONGO_URL"),
		MongoDBName: getEnvStr("MONGO_DB_NAME"),

		RedisURL: getEnvStr("REDIS_URL"),

		QueueUsername:    getEnvStr("QUEUE_USERNAME"),
		QueuePassword:    getEnvStr("QUEUE_PASSWORD"),
		QueueConcurrency: getEnvInt("QUEUE_CONCURRENCY"),

		SentryDSN:     getEnvStr("SENTRY_DSN"),
		SentryMachine: getEnvStr("SENTRY_MACHINE"),

		TelegramEnabled:  getEnvBool("TELEGRAM_ENABLED"),
		TelegramBotToken: getEnvStr("TELEGRAM_BOT_TOKEN"),
	}

	fmt.Printf("⚡️ [config]: loaded \n")
}
