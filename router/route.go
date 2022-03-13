package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"repeater/typies"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Route(w *discordgo.Session, r *discordgo.MessageCreate) {
	go func() {
		for range time.Tick(time.Second * 1) {
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
				w.ChannelMessageSend(data[i].ChannelId, data[i].Message+"["+strconv.Itoa(data[i].Id)+"]")
			}
			fmt.Println("dasdw")
		}
	}()

	msg := strings.Split(r.Content, "")

	if msg[0] != "$" {
		return
	}
	cmd := strings.Split(strings.Join(msg[1:], ""), " ")

	if cmd[0] == "반복" {
		StartRepeat(w, r)
	}
	if cmd[0] == "종료" {
		StopRepeat(w, r)
	}
}
