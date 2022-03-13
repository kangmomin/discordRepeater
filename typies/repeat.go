package typies

type RepeatData struct {
	UserID    string `json:"userId"`
	Time      string `json:"time"`
	Count     int    `json:"count"`
	Message   string `json:"message"`
	ChannelId string `json:"channelId"`
}
