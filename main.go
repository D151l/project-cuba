package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var session *discordgo.Session = nil

func main() {
	log.SetPrefix("discord-bot: ")
	log.Println("Starting bot...")

	token := os.Getenv("token")

	log.Println("Creating Discord session...")

	newSession, err := discordgo.New("Bot " + token)
	session = newSession

	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
		return
	}

	session.Identify.Intents = discordgo.IntentsAll

	err = session.Open()
	if err != nil {
		log.Fatal("Error opening connection: ", err)
		return
	}

	log.Println("Bot is now running. Press CTRL-C to exit.")

	defer func(session *discordgo.Session) {
		log.Println("Closing connection...")

		err := session.Close()
		if err != nil {
			log.Println("Error closing connection: ", err)
			return
		}
	}(session)

	testCommand()
	setStatus()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
