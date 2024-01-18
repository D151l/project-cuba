package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func setupMessageCommand() {
	_, err := session.ApplicationCommandBulkOverwrite(session.State.User.ID, guildID, []*discordgo.ApplicationCommand{
		{
			Name:        "message",
			Description: "A command that sends a embed message.",
		},
	})
	if err != nil {
		log.Println("Error overwriting commands: ", err)
		return
	}

	session.AddHandler(messageCommandHandler)
}

func messageCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name != "message" {
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:       "Erfolgreich!",
		Description: "Dir wurde eine Privatnachricht gesendet! :mailbox_with_mail: \nDies kann einige Sekunden dauern.",
		Color:       0x00ff00,
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})

	if err != nil {
		log.Println("Error responding to interaction: ", err)
		return
	}

	channel, err := session.UserChannelCreate(i.Member.User.ID)
	if err != nil {
		log.Println("Error creating user channel: ", err)
		return
	}

	embedPrivet := &discordgo.MessageEmbed{
		Title:       "Willkommen im setup!",
		Description: "Bitte folge den Anweisungen des Bots. \n\n Welches format soll die Nachricht haben? \n\n 1. Embed \n 2. Embed mit Liste \n\n Bitte wähle eine Zahl aus.",
		Color:       0x00ff00,
	}

	_, err = session.ChannelMessageSendEmbed(channel.ID, embedPrivet)
	if err != nil {
		log.Println("Error sending message: ", err)
		return
	}

	sendDemoEmbed(channel)
	sendDemoListedEmbed(channel)
}
