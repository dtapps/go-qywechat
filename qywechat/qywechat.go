package qywechat

import (
	"encoding/json"
	"fmt"
	"github.com/dtapps/go-qywechat/qywechat/message"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const api = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send"

type QyBot struct {
	Key string
}

type response struct {
	Errcode   int64  `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

func (bot *QyBot) Send(msg message.Message) (response, error) {
	var response response
	qyUrl := fmt.Sprintf("%s?key=%s", api, bot.Key)
	j, e := json.Marshal(msg)
	if e != nil {
		return response, e
	}
	log.Printf("qyUrl：%v\n", qyUrl)
	log.Printf("msg：%v\n", msg)
	resp, e := http.Post(qyUrl, "application/json", strings.NewReader(string(j)))
	if e != nil {
		return response, e
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	e = json.Unmarshal(body, &response)
	if e != nil {
		return response, e
	}
	return response, nil
}
