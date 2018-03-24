package main

import (
	"encoding/json"
)

type CommandInterface interface {
	Execute() string
	GetUri() string
}

type Command struct {
	Client                          Client
	Uri, HttpMethod, TelegramMethod string
	Payload                         []byte
	Response                        Response
}

func NewCommand(uri, httpMethod, telegramMethod string, payload []byte) Command {
	return Command{NewClient(), uri, httpMethod, telegramMethod, payload, Response{}}
}

func (c *Command) Execute() {
	c.Client.ExecuteCommand(c)
}

func (c *Command) GetUri() string {
	return c.Uri + c.TelegramMethod
}

//getMe command
type GetMe struct {
	Command
}

func NewGetMe(uri string) GetMe {
	return GetMe{NewCommand(uri, "GET", "getme", []byte(""))}
}

func (g *GetMe) ParseResult() User {
	var u User
	json.Unmarshal(g.Response.Result, &u)

	return u
}

//sendMessage command
type SendMessage struct {
	Command
}

func NewSendMessage(uri string, payload []byte) SendMessage {
	return SendMessage{NewCommand(uri, "POST", "sendmessage", payload)}
}

func (s *SendMessage) ParseResult() Message {
	var m Message
	json.Unmarshal(s.Response.Result, &m)

	return m
}

//getUpdates command
type GetUpdates struct {
	Command
}

func NewGetUpdates(uri string) GetUpdates {
	return GetUpdates{NewCommand(uri, "GET", "getupdates", []byte(""))}
}

func (g *GetUpdates) ParseResult() []Update {
	var u []Update
	json.Unmarshal(g.Response.Result, &u)

	return u
}
