package qywechat

import (
	"github.com/dtapps/go-qywechat/qywechat"
	"github.com/dtapps/go-qywechat/qywechat/message"
	"log"
)

func main() {
	bot := qywechat.QyBot{
		Key: "",
	}
	msg := message.Message{
		MsgType: message.TextStr,
		Text: message.Text_{
			Content: "测试",
		},
	}
	send, err := bot.Send(msg)
	if err != nil {
		log.Printf("err：%v\n", err)
		return
	}
	log.Printf("send：%v\n", send)
}
