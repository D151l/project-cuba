package main

import (
	"log"
	"os"
)

func setStatus() {
	statusType := os.Getenv("statusType")
	statusContent := os.Getenv("statusContent")

	if statusType == "PLAYING" {
		err := session.UpdateGameStatus(0, statusContent)
		if err != nil {
			log.Println("Error setting status: ", err)
			return
		}
	}

	if statusType == "STREAMING" {
		streamingURL := os.Getenv("streamingURL")
		err := session.UpdateStreamingStatus(0, statusContent, streamingURL)
		if err != nil {
			log.Println("Error setting status: ", err)
			return
		}
	}

	if statusType == "LISTENING" {
		err := session.UpdateListeningStatus(statusContent)
		if err != nil {
			log.Println("Error setting status: ", err)
			return
		}
	}

	if statusType == "WATCHING" {
		err := session.UpdateListeningStatus(statusContent)
		if err != nil {
			log.Println("Error setting status: ", err)
			return
		}
	}
}
