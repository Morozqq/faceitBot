package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"sync"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize configuration
	LoadConfig()

	// Get the bot token from environment variables
	token := GetDiscordToken()
	if token == "" {
		log.Fatalf("No bot token provided. Please set the DISCORD_TOKEN environment variable.")
	}

	// Create a new Discord session
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	// Register the message handler
	dg.AddHandler(messageCreate)

	// Open the websocket connection to Discord and begin listening.
	if err := dg.Open(); err != nil {
		log.Fatalf("Error opening connection: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go runServer(&wg)

	log.Println("Bot is now running. Press CTRL+C to exit.")
	select {} // Block until CTRL+C is pressed
	wg.Wait()
}
