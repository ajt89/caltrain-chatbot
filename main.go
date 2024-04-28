package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

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
	fmt.Println(m.Content)

	// Strip username mention out of incoming command
	username := fmt.Sprintf("<@%s> ", s.State.User.ID)
	content_no_username := strings.TrimPrefix(m.Content, username)

	// Generate list of commands and get length
	commands := strings.Split(content_no_username, " ")
	command_len := len(commands)
	log.Printf("commands: %s\n", commands)
	log.Printf("command_len: %d\n", command_len)

	// If the message is "headers" reply with headers timestamp
	if command_len == 1 && commands[0] == "headers" {
		data := caltrain.GetRealTime()
		message := fmt.Sprintf("Header timestamp: %d", data.RealTime.Header.Timestamp)
		s.ChannelMessageSend(m.ChannelID, message)
		log.Printf("Send message: %s\n", message)
	}
	// nta (Next to Arrive)
	if commands[0] == "nta" {
		stopName := strings.Join(commands[1:], " ")
		log.Printf("stop: %s\n", stopName)

		stop := caltrain.GetStopByName(stopName)
		stopNId := stop.NorthboundId
		stopSId := stop.SouthboundId

		data := caltrain.ParseCalTrainStop(stopNId, stopSId)
		trainCountMsg := fmt.Sprintf("%d trains found:\n", len(data.CalTrainVehicles))
		trainInfoMsgs := []string{}
		for _, t := range data.CalTrainVehicles {
			currentTime := time.Now()
			arrivalHumanTime := time.Unix(t.ArrivalTime, 0)
			departureHumanTime := time.Unix(t.DepartureTime, 0)
			timeLeft := arrivalHumanTime.Sub(currentTime)
			trainInfoMsg := fmt.Sprintf(
				"train id: %s\n arrival: %s\n departure: %s\n direction: %s\n stops left: %d\n time left: %v\n current stop: %s\n train type: %s\n",
				t.Id, arrivalHumanTime.Local().Format(time.Kitchen), departureHumanTime.Local().Format(time.Kitchen), t.Direction, t.StopsLeft, timeLeft.Truncate(time.Second), t.CurrentStop, t.TripType)
			trainInfoMsgs = append(trainInfoMsgs, trainInfoMsg)
		}
		trainInfoMsgsJoin := strings.Join(trainInfoMsgs, "\n")
		message := fmt.Sprintf("%s\n%s", trainCountMsg, trainInfoMsgsJoin)
		s.ChannelMessageSend(m.ChannelID, message)
		log.Printf("Send message: %s\n", message)
	}
}
