package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func sendDemoListedEmbed(channel *discordgo.Channel) {
	fields := make([]*discordgo.MessageEmbedField, 0)
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:  "Lorem ipsum dolor sit amet",
		Value: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua.",
	})
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:  "Lorem ipsum dolor sit amet",
		Value: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua.",
	})
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:  "Lorem ipsum dolor sit amet",
		Value: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua.",
	})

	embed := &discordgo.MessageEmbed{
		Title:       "2. Embed mit Liste",
		Description: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
		Color:       0xffff00,
		Fields:      fields,
		Footer:      &discordgo.MessageEmbedFooter{Text: "Footer"},
	}

	_, err := session.ChannelMessageSendEmbed(channel.ID, embed)
	if err != nil {
		log.Println("Error sending message: ", err)
		return
	}
}

func sendDemoEmbed(channel *discordgo.Channel) {
	embed := &discordgo.MessageEmbed{
		Title:       "1. Embed",
		Description: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
		Color:       0xffff00,
		Footer:      &discordgo.MessageEmbedFooter{Text: "Footer"},
	}

	_, err := session.ChannelMessageSendEmbed(channel.ID, embed)
	if err != nil {
		log.Println("Error sending message: ", err)
		return
	}
}
