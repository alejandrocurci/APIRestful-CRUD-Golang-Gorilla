package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"api-crud-gorilla/models"

	"github.com/gorilla/mux"
)

func GetMessages(w http.ResponseWriter, r *http.Request) {
	// fetch all structs in database
	var messages []models.Message
	for _, msg := range models.MessageStore.Data {
		messages = append(messages, msg)
	}
	// parse list into json and send it in response
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(messages)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	// decode struct from json body
	var msg models.Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		panic(err)
	}
	// set current time and save in database
	msg.SentAt = time.Now()
	models.Index++
	models.MessageStore.Data[models.Index] = msg
	// parse struct to json and send it back in response
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	// get params and body from the request
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}
	// msg is the updated message
	var msg models.Message
	err = json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		panic(err)
	}
	// m is the original message
	// update message if found
	if m, ok := models.MessageStore.Data[id]; ok {
		msg.SentAt = m.SentAt
		delete(models.MessageStore.Data, id)
		models.MessageStore.Data[id] = msg
		log.Print("successful update")
	} else {
		log.Print("id not found")
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	// get params from the request
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}
	// delete the message if found
	if _, ok := models.MessageStore.Data[id]; ok {
		delete(models.MessageStore.Data, id)
		log.Print("successful delete")
	} else {
		log.Print("id not found")
	}
	w.WriteHeader(http.StatusNoContent)
}
