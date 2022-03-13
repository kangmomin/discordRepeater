package router

import (
	"encoding/json"
	"io/ioutil"
	"repeater/logger"
	"repeater/typies"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var log = logger.Logger

func StartRepeat(w *discordgo.Session, r *discordgo.MessageCreate) {
	var repeater []typies.RepeatData
	var data typies.RepeatData
	byteData, err := ioutil.ReadFile("./data/user.json")
	if err != nil {
		log.Println(err)
		w.ChannelMessageSend(r.ChannelID, "반복을 실행하지 못하였습니다.")
		return
	}

	err = json.Unmarshal(byteData, &repeater)

	if err != nil {
		log.Println(err)
		w.ChannelMessageSend(r.ChannelID, "반복을 실행하지 못하였습니다.")
		return
	}

	msg := strings.Split(r.Content, " ")
	cmd := strings.Join(msg[1:], " ")
	data.ChannelId = r.ChannelID
	data.Message = cmd
	data.Time = time.Now().Format("2008-01-08")

	repeater = append(repeater, data)

	byteData, err = json.Marshal(repeater)
	if err != nil {
		log.Println(err)
		w.ChannelMessageSend(r.ChannelID, "반복을 실행하지 못하였습니다.")
		return
	}

	ioutil.WriteFile("./data/user.json", byteData, 0644)
}

func StopRepeat(w *discordgo.Session, r *discordgo.MessageCreate) {
	var repeater []typies.RepeatData
	var data typies.RepeatData
	byteData, err := ioutil.ReadFile("./data/user.json")
	if err != nil {
		log.Println(err)
		w.ChannelMessageSend(r.ChannelID, "반복을 실행하지 못하였습니다.")
		return
	}

	err = json.Unmarshal(byteData, &repeater)

	if err != nil {
		log.Println(err)
		w.ChannelMessageSend(r.ChannelID, "반복을 실행하지 못하였습니다.")
		return
	}
}
