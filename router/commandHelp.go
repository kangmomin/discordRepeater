package router

import "github.com/bwmarrin/discordgo"

func CmdHelp(w *discordgo.Session, r *discordgo.MessageCreate) {
	w.ChannelMessageSend(r.ChannelID, `
$반복 [반복할 말]
$종료 [id]: 해당 id값의 반복을 종료한다.
	`)
}
