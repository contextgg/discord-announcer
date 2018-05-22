package main

import (
	"log"
	"os"

	"github.com/nordicgaming/discord-announcer/cmd/discord-announcer/announcer"
	"github.com/nordicgaming/discord-announcer/cmd/discord-announcer/config"
)

func main() {
	cfg := new(config.Config)
	if err := config.ReadConfig(cfg); err != nil {
		log.Fatalf("ReadConfig error, %v", err)
		return
	}

	noucer, err := announcer.NewAnnouncer(cfg)
	if err != nil {
		log.Fatalf("Could not create announcer: %s", err)
		return
	}

	filePaths := os.Args[1:]

	announcements, err := announcer.ParseFiles(filePaths)
	if err != nil {
		log.Fatalf("LoadFiles error: %v", err)
		return
	}

	if len(announcements) == 0 {
		log.Println("No announcement was supplied")
		return
	}

	if err := noucer.SendAnnouncements(announcements); err != nil {
		log.Fatalf("Could not send message: %s", err)
		return
	}
}
