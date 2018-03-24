package main

import (
	"encoding/json"
	"os"
	"strings"
)

type Bot struct {
	Token string `json:"token"`
}

func NewBot(name string) Bot {
	c := GetConf()
	var b Bot
	json.Unmarshal(c.Bots[name], &b)

	return b
}

func (b *Bot) GetMe() User {
	cmd := NewGetMe(b.GetUri())
	cmd.Execute()
	res := cmd.ParseResult()

	return res
}

func (b *Bot) SendMessage(msg Message) Message {
	payload, _ := json.Marshal(msg)
	cmd := NewSendMessage(b.GetUri(), payload)
	cmd.Execute()
	res := cmd.ParseResult()

	return res
}

func (b *Bot) GetUpdates() []Update {
	cmd := NewGetUpdates(b.GetUri())
	cmd.Execute()
	res := cmd.ParseResult()

	return res
}

func (b *Bot) GetUri() string {

	c := GetConf()
	t := os.Getenv(b.Token)
	return strings.Replace(c.Uri, "<token>", t, 1)
}
