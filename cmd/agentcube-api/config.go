package main

import (
	"os"
	"strings"
)

// Config - contains the Agent Cube API configuration.
type Config struct {
	Name                    string
	Version                 string
	Environment             string
	Port                    string
	MongoDBURI              string
	MongoDBName             string
	WebsocketAllowedOrigins []string
	OpenRouterAPIKey        string
	OpenRouterBaseURL       string
	OpenRouterHTTPReferrer  string
	OpenRouterTitle         string
}

// NewConfig - loads the Agent Cube API configuration from environment variables.
func NewConfig() Config {
	return Config{
		Name:                    "Agent Cube API",
		Version:                 "v1",
		Environment:             os.Getenv("ENVIRONMENT"),
		Port:                    os.Getenv("PORT"),
		MongoDBURI:              os.Getenv("MONGODB_URI"),
		MongoDBName:             os.Getenv("MONGODB_DB"),
		WebsocketAllowedOrigins: strings.Split(os.Getenv("WEBSOCKET_ALLOWED_ORIGINS"), ","),
		OpenRouterAPIKey:        os.Getenv("OPENROUTER_API_KEY"),
		OpenRouterBaseURL:       os.Getenv("OPENROUTER_BASE_URL"),
		OpenRouterHTTPReferrer:  os.Getenv("OPENROUTER_HTTP_REFERRER"),
		OpenRouterTitle:         os.Getenv("OPENROUTER_X_TITLE"),
	}
}
