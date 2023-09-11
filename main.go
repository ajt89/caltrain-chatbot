package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/ajt89/caltrain-chatbot/caltrain"
	"github.com/bwmarrin/discordgo"
)

var s *discordgo.Session

func main() {
	discord_token := os.Getenv("DISCORD_TOKEN")
	dg, err := discordgo.New("Bot " + discord_token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func isUserIdInMention(users []*discordgo.User, userId string) bool {
	for _, value := range users {
		if value.ID == userId {
			return true
		}
	}
	return false
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// Only respond if bot is mentioned
	if !isUserIdInMention(m.Mentions, s.State.User.ID) {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if strings.Contains(m.Content, "headers") {
		data := caltrain.GetRealTime()
		message := fmt.Sprintf("Header timestamp: %d", data.RealTime.Header.Timestamp)
		s.ChannelMessageSend(m.ChannelID, message)
		log.Printf("Send message: %s\n", message)
	}
}
