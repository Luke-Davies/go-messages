package main

// Message has text and an ID
type Message struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

// Messages is an array of messages
type Messages []Message