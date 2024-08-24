package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"sync"
)

func main() {

	LoadConfig()

	token := GetDiscordToken()
	if token == "" {
		log.Fatalf("No bot token provided. Please set the DISCORD_TOKEN environment variable.")
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	dg.AddHandler(messageCreate)

	// Open the websocket connection to Discord and begin listening
	if err := dg.Open(); err != nil {
		log.Fatalf("Error opening connection: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go runServer(&wg)

	log.Println("Bot is now running. Press CTRL+C to exit.")

	select {}
	wg.Wait()
}
