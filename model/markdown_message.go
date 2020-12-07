package model

type markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

//markdown消息
type markdownMessage struct {
	message
	markdown `json:"markdown"`
}

func NewMarkDownMessage(title, content string) markdownMessage {
	return markdownMessage{message{MsgType: "markdown"}, markdown{title, content}}
}
