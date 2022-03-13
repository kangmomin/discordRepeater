package main

import (
	"os"
	"os/signal"
	"repeater/logger"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var log = logger.Logger

func main() {
	err := godotenv.Load("TOKEN")
	if err != nil {
		log.Fatalln(err)
	}

	sc := make(chan os.Signal, 1)

	c, err := discordgo.New("Bot " + os.Getenv("TOKEN"))

	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	c.AddHandler(router.Route)

	c.Identify.Intents = discordgo.IntentGuildMessages

	err = c.Open()
	if err != nil {
		log.Fatalln(err)
	}

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
