package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var conf *Conf

func webhook(w http.ResponseWriter, r *http.Request) {
	Bot := NewBot("donttrythisathome")
	body, _ := ioutil.ReadAll(r.Body)

	var update Update
	json.Unmarshal([]byte(body), &update)
	chatId := update.Message.Chat.Id

	Bot.SendMessage(Message{ChatId: chatId, Text: "Привет " + update.Message.From.Username})
}

func main() {
	http.HandleFunc("/webhook", webhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Conf struct {
	Uri  string                     `json:"telegram_uri"`
	Bots map[string]json.RawMessage `json:"bots"`
}

func NewConf() *Conf {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&conf); err != nil {
		fmt.Println("error:", err)
	}

	return conf
}

func GetConf() *Conf {
	if conf == nil {
		conf = NewConf()
	}
	return conf
}
