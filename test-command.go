package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func testCommand() {
	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "ping" {
			_, err := s.ChannelMessageSend(m.ChannelID, "pong")
			if err != nil {
				log.Println("Error sending message: ", err)
				return
			}
		}
	})
}
