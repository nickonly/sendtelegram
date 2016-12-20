package main

import (
	"net/url"
	"encoding/json"
)



type Response struct {
	Ok bool `json:"ok"`
	ErrorCode int64 `json:"error_code"`
	Description string `json:"description"`
}



func ParseResponse(bytes []byte) (*Response, error) {
	var response Response
	err := json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}



func SendMessage(token string, chat_id string, text string) (*Response, error) {
	values := url.Values{}
	values.Add("chat_id", chat_id)
	values.Add("text", text)
	response, err := Request(token, "sendMessage", values)
	if err != nil {
		return nil, err
	}
	return ParseResponse(response)
}

