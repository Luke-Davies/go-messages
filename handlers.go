package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var nextMsgID = 1000

// GetMessages returns all messages
func GetMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(StoredMessages)
}

// CreateMessage adds a new message
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	// check if valid supported content type
	set := map[string]bool{"text/plain": true, "application/x-www-form-urlencoded": true}
	if _, v := set[r.Header.Get("Content-Type")]; !v {
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(400)
		fmt.Fprintln(w, "Content-Type type not supported.")
		return
	}
	var body, _ = ioutil.ReadAll(r.Body)
	messageText := string(body)

	// check a message is passed
	if len(messageText) == 0 {
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(400)
		fmt.Fprintln(w, "No data provided.")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	messageID := nextMsgID
	nextMsgID++
	message := Message{ID: messageID, Text: messageText}
	StoredMessages = append(StoredMessages, message)
	response := struct {
		ID int `json:"id"`
	}{
		ID: messageID,
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(response)
}

// DeleteMessages returns a message for a given msgId
func DeleteMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	StoredMessages = Messages{}
	w.WriteHeader(204)
	fmt.Fprintln(w, "All messages cleared")
}

// GetMessage returns a message for a given msgId
func GetMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	vars := mux.Vars(r)
	msgID, err := strconv.Atoi(vars["msgId"])
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "Invalid ID provided. Expected integer, got ", vars["msgId"])
		return
	}

	for _, v := range StoredMessages {
		if v.ID == msgID {
			w.WriteHeader(200)
			fmt.Fprintln(w, v.Text)
			return
		}
	}

	w.WriteHeader(404)
	fmt.Fprintln(w, "No message found for ID", msgID)
}
