package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/signal"
	"repeater/logger"
	"repeater/router"
	"repeater/typies"
	"strconv"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/lemon-mint/godotenv"
)

var log = logger.Logger

func main() {
	godotenv.Load()

	sc := make(chan os.Signal, 1)

	c, err := discordgo.New("Bot " + os.Getenv("TOKEN"))

	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	c.AddHandler(router.Route)

	go func() {
		for range time.Tick(time.Second * 3) {
			byteData, err := ioutil.ReadFile("./data/user.json")
			if err != nil {
				log.Println(err)
				return
			}

			var data []typies.RepeatData
			err = json.Unmarshal(byteData, &data)
			if err != nil {
				log.Println(err)
				return
			}

			for i := 0; i < len(data); i++ {
				c.ChannelMessageSend(data[i].ChannelId, data[i].Message+"["+strconv.Itoa(data[i].Id)+"]")
			}
		}
	}()

	c.Identify.Intents = discordgo.IntentGuildMessages

	err = c.Open()
	if err != nil {
		log.Fatalln(err)
	}

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
