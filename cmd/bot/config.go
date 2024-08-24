package main

import (
	"os"
)

var (
	NameUrl       string
	APIKey        string
	BaseURL       string
	StatsEndPoint string
)

func LoadConfig() {
	NameUrl = os.Getenv("NAME_URL")
	APIKey = os.Getenv("API_KEY")
	BaseURL = os.Getenv("BASE_URL")
	StatsEndPoint = os.Getenv("STATS_ENDPOINT")
}

func GetDiscordToken() string {
	return os.Getenv("DISCORD_TOKEN")
}
