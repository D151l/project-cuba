package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strconv"
	"strings"
)

type EmbedSetup struct {
	channel     string
	embedType   int
	title       string
	colorRed    int
	colorGreen  int
	colorBlue   int
	description string
	list        map[string]string
	footer      string
	preview     bool
}

var currentEmbedsSetup map[string]EmbedSetup = make(map[string]EmbedSetup)

func setupEmbedHandler() {
	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.GuildID != "" {
			return
		}

		if m.Author.ID == s.State.User.ID {
			return
		}

		channel, err := session.UserChannelCreate(m.Author.ID)
		if err != nil {
			log.Println("Error creating user channel: ", err)
			return
		}

		if _, ok := currentEmbedsSetup[m.Author.ID]; !ok {
			embed := &discordgo.MessageEmbed{
				Title:       "Fehler!",
				Description: "Du hast noch keinen Embed Setup gestartet. Bitte starte einen Embed Setup mit dem Command `/message` auf dem Server.",
				Color:       0xff0000,
			}

			_, err := session.ChannelMessageSendEmbed(channel.ID, embed)
			if err != nil {
				log.Println("Error sending message: ", err)
				return
			}

			return
		}

		embedSetup := currentEmbedsSetup[m.Author.ID]

		if embedSetup.embedType == 0 {
			content := m.Message.Content

			if content != "1" && content != "2" {
				embed := &discordgo.MessageEmbed{
					Title:       "Fehler!",
					Description: "Bitte wähle eine Zahl zwischen 1 und 2 aus.",
					Color:       0xff0000,
				}

				_, err := session.ChannelMessageSendEmbed(channel.ID, embed)
				if err != nil {
					log.Println("Error sending message: ", err)
					return
				}

				return
			}

			embedSetup.embedType = int(content[0]) - 48
			currentEmbedsSetup[m.Author.ID] = embedSetup

			embed := &discordgo.MessageEmbed{
				Title:       "Erfolgreich!",
				Description: "Bitte gib nun den Titel der Nachricht ein.",
				Color:       0xffff00,
			}

			_, err := session.ChannelMessageSendEmbed(channel.ID, embed)
			if err != nil {
				log.Println("Error sending message: ", err)
				return
			}

			return
		}

		if embedSetup.title == "" {
			embedSetup.title = m.Message.Content
			currentEmbedsSetup[m.Author.ID] = embedSetup

			// 255, 255, 255
			embed := &discordgo.MessageEmbed{
				Title:       "Erfolgreich!",
				Description: "Bitte gib nun die Farbe der Nachricht ein. \n\n Bitte gib die Farbe in Dezimalzahlen an. \n\n Beispiel: 255, 255, 255",
				Color:       0xffff00,
			}

			_, err := session.ChannelMessageSendEmbed(channel.ID, embed)
			if err != nil {
				log.Println("Error sending message: ", err)
				return
			}

			return
		}

		if embedSetup.colorRed == -1 {
			content := m.Message.Content

			split := strings.Split(content, ", ")
			if len(split) != 3 {
				embed := &discordgo.MessageEmbed{
					Title:       "Fehler!",
					Description: "Bitte gib eine gültige Farbe ein. \n\n Die Farbe muss aus 3 Zahlen bestehen, die durch ein Komma getrennt sind. \n\n Beispiel: 255, 255, 255",
					Color:       0xff0000,
				}

				_, err := session.ChannelMessageSendEmbed(channel.ID, embed)
				if err != nil {
					log.Println("Error sending message: ", err)
					return
				}

				return
			}

			red := split[0]
			green := split[1]
			blue := split[2]

			if len(red) > 3 || len(green) > 3 || len(blue) > 3 {
				embed := &discordgo.MessageEmbed{
					Title:       "Fehler!",
					Description: "Bitte gib eine gültige Farbe ein. \n\n Die Farbe darf maximal 3 Zahlen lang sein.",
					Color:       0xff0000,
				}

				_, err := session.ChannelMessageSendEmbed(channel.ID, embed)
				if err != nil {
					log.Println("Error sending message: ", err)
					return
				}

				return
			}

			redInt, err := strconv.Atoi(red)
			if err != nil {
				embed := &discordgo.MessageEmbed{
					Title:       "Fehler!",
					Description: "Bitte gib eine gültige Farbe ein. \n\n Dein Rot-Wert ist ungültig.",
					Color:       0xff0000,
				}

				_, err := session.ChannelMessageSendEmbed(channel.ID, embed)
				if err != nil {
					log.Println("Error sending message: ", err)
					return
				}
				return
			}

			greenInt, err := strconv.Atoi(green)
			if err != nil {
				embed := &discordgo.MessageEmbed{
					Title:       "Fehler!",
					Description: "Bitte gib eine gültige Farbe ein. \n\n Dein Grün-Wert ist ungültig.",
					Color:       0xff0000,
				}

				_, err := session.ChannelMessageSendEmbed(channel.ID, embed)
				if err != nil {
					log.Println("Error sending message: ", err)
					return
				}
				return
			}

			blueInt, err := strconv.Atoi(blue)
			if err != nil {
				embed := &discordgo.MessageEmbed{
					Title:       "Fehler!",
					Description: "Bitte gib eine gültige Farbe ein. \n\n Dein Blau-Wert ist ungültig.",
					Color:       0xff0000,
				}

				_, err := session.ChannelMessageSendEmbed(channel.ID, embed)
				if err != nil {
					log.Println("Error sending message: ", err)
					return
				}
				return
			}

			if redInt < 0 || redInt > 255 || greenInt < 0 || greenInt > 255 || blueInt < 0 || blueInt > 255 {
				embed := &discordgo.MessageEmbed{
					Title:       "Fehler!",
					Description: "Bitte gib eine gültige Farbe ein. \n\n Die Farbe muss zwischen 0 und 255 liegen.",
					Color:       0xff0000,
				}

				_, err := session.ChannelMessageSendEmbed(channel.ID, embed)
				if err != nil {
					log.Println("Error sending message: ", err)
					return
				}

				return
			}

			embedSetup.colorRed = redInt
			embedSetup.colorGreen = greenInt
			embedSetup.colorBlue = blueInt

			currentEmbedsSetup[m.Author.ID] = embedSetup

			embed := &discordgo.MessageEmbed{
				Title:       "Erfolgreich!",
				Description: "Bitte gib nun die Beschreibung der Nachricht ein.",
				Color:       0xffff00,
			}

			_, err = session.ChannelMessageSendEmbed(channel.ID, embed)
			if err != nil {
				log.Println("Error sending message: ", err)
				return
			}

			return
		}

		if embedSetup.description == "" {
			embedSetup.description = m.Message.Content
			currentEmbedsSetup[m.Author.ID] = embedSetup

			embed := &discordgo.MessageEmbed{
				Title:       "Erfolgreich!",
				Description: "Bitte gib nun den Footer der Nachricht ein. \n\n Wenn du keinen Footer haben möchtest, schreibe `none`.",
				Color:       0xffff00,
			}

			_, err := session.ChannelMessageSendEmbed(channel.ID, embed)
			if err != nil {
				log.Println("Error sending message: ", err)
				return
			}

			return
		}

		if embedSetup.footer == "" {

			if m.Message.Content != "none" {
				embedSetup.footer = m.Message.Content
				currentEmbedsSetup[m.Author.ID] = embedSetup
			} else {
				embedSetup.footer = "none"
				currentEmbedsSetup[m.Author.ID] = embedSetup
			}

			embed := &discordgo.MessageEmbed{
				Title:       "Erfolgreich!",
				Description: "Dir wird nun eine Vorschau der Nachricht gesendet.",
				Color:       0xffff00,
			}

			_, err := session.ChannelMessageSendEmbed(channel.ID, embed)
			if err != nil {
				log.Println("Error sending message: ", err)
				return
			}

			if embedSetup.embedType == 1 {
				embedSetup.preview = true
				currentEmbedsSetup[m.Author.ID] = embedSetup

				previewEmbed(channel, embedSetup)

				embed = &discordgo.MessageEmbed{
					Title:       "Der letzte Schritt!",
					Description: "Soll die Nachricht gesendet werden? \n\n 1. Ja \n 2. Nein, ich möchte etwas ändern \n\n Bitte wähle eine Zahl aus.",
					Color:       0xffff00,
				}

				_, err = session.ChannelMessageSendEmbed(channel.ID, embed)
				if err != nil {
					log.Println("Error sending message: ", err)
					return
				}
			}
			return
		}

		if embedSetup.preview {
			content := m.Message.Content

			if content != "1" && content != "2" {
				embed := &discordgo.MessageEmbed{
					Title:       "Fehler!",
					Description: "Bitte wähle eine Zahl zwischen 1 und 2 aus.",
					Color:       0xff0000,
				}

				_, err := session.ChannelMessageSendEmbed(channel.ID, embed)
				if err != nil {
					log.Println("Error sending message: ", err)
					return
				}

				return
			}

			if content == "1" {
				sendEmbed(embedSetup)
				delete(currentEmbedsSetup, m.Author.ID)

				embed := &discordgo.MessageEmbed{
					Title:       "Erfolgreich!",
					Description: "Die Nachricht wurde erfolgreich gesendet!",
					Color:       0x00ff00,
				}

				_, err := session.ChannelMessageSendEmbed(channel.ID, embed)
				if err != nil {
					log.Println("Error sending message: ", err)
					return
				}
			}

			if content == "2" {

			}
			return
		}
	})
}

func previewEmbed(channel *discordgo.Channel, embedSetup EmbedSetup) {
	footer := embedSetup.footer
	if footer == "none" {
		footer = ""
	}

	embed := &discordgo.MessageEmbed{
		Title:       embedSetup.title,
		Description: embedSetup.description,
		Color:       (embedSetup.colorRed * 256 * 256) + (embedSetup.colorGreen * 256) + embedSetup.colorBlue,
		Footer:      &discordgo.MessageEmbedFooter{Text: footer},
	}

	_, err := session.ChannelMessageSendEmbed(channel.ID, embed)
	if err != nil {
		log.Println("Error sending message: ", err)
		return
	}
}

func sendEmbed(embedSetup EmbedSetup) {
	footer := embedSetup.footer
	if footer == "none" {
		footer = ""
	}

	embed := &discordgo.MessageEmbed{
		Title:       embedSetup.title,
		Description: embedSetup.description,
		Color:       (embedSetup.colorRed * 256 * 256) + (embedSetup.colorGreen * 256) + embedSetup.colorBlue,
		Footer:      &discordgo.MessageEmbedFooter{Text: footer},
	}

	_, err := session.ChannelMessageSendEmbed(embedSetup.channel, embed)
	if err != nil {
		log.Println("Error sending message: ", err)
		return
	}
}
