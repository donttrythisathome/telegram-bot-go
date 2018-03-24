package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct{}

func NewClient() Client {
	return Client{}
}

func (c *Client) ExecuteCommand(cmd *Command) {

	client := &http.Client{}
	req, _ := http.NewRequest(cmd.HttpMethod, cmd.GetUri(), bytes.NewBuffer(cmd.Payload))
	req.Header.Set("Content-Type", "application/json")
	response, _ := client.Do(req)
	defer response.Body.Close()
	contents, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal([]byte(contents), &cmd.Response)
	if !cmd.Response.Ok {
		panic(string(contents))
	}
}

type Response struct {
	Ok     bool `json:"ok"`
	Result json.RawMessage
}
