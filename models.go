package main

type User struct {
	Id           int32  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type Message struct {
	MessageId int32  `json:"message_id"`
	Text      string `json:"text"`
	ChatId    int32  `json:"chat_id"`
	Chat      Chat   `json:"chat"`
	From      User   `json:"from"`
	Date      int32  `json:"date"`
}

type Update struct {
	UpdateId int32   `json:"update_id"`
	Message  Message `json:"message"`
}

type Chat struct {
	Id        int32  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}
