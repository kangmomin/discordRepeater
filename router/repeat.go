package router

import (
	"encoding/json"
	"io/ioutil"
	"repeater/logger"
	"repeater/typies"
	"repeater/util"
	"strconv"
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

	msg := strings.Split(r.Content, "")
	if len(msg) < 5 {
		w.ChannelMessageSend(r.ChannelID, "반복을 실행하지 못하였습니다.")
		return
	}

	cmd := strings.Join(msg[1:], " ")
	data.ChannelId = r.ChannelID
	data.Message = cmd
	data.Time = time.Now().Format("2006-01-02 15:04:05")
	data.Id = util.GenerateId(repeater)
	data.UserID = r.Author.ID

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

	cmd := strings.Split(r.Content, " ")
	id, err := strconv.Atoi(cmd[1])

	if err != nil {
		w.ChannelMessageSend(r.ChannelID, "id값이 잘못되었습니다.")
		return
	}

	comment := ""
	for idx, val := range repeater {
		if id != val.Id {
			continue
		}

		repeater = append(repeater[:idx], repeater[idx+1:]...)
		comment = "id: " + strconv.Itoa(val.Id) + " / started-user <@" + val.UserID + "> / start-time: " + val.Time
	}

	byteData, err = json.Marshal(repeater)
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile("./data/user.json", byteData, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	w.ChannelMessageSend(r.ChannelID, comment)
}
