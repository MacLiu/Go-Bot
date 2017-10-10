package main

type Messenger struct {
	id string `json:"id"`;
	name string `json:"name"`;
}

type Conversation struct {
	id string `json:"id"`;
	name string `json:"name"`;
}

type Recipient struct {
	id string `json:"id"`;
	name string `json:"name"`;
}

type Activity struct {
	message Message `json:"message"`;
}

type Message struct {
	activityType string `json:"type"`;
	id string `json:"id"`;
	timestamp string `json:"timestamp"`;
	serviceUrl string `json:"serviceUrl"`;
	channelId string `json:"channelId"`;
	from Messenger `json:"from"`;
	conversation Conversation `json:"conversation"`;
	recipient Recipient `json:"recipient"`;
	text string `json:"text"`;
}