package handlers

import (
	"encoding/json"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	// decode struct from json body
	var msg models.Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// set current time and save in database
	msg.SentAt = time.Now()
	models.Index++
	models.MessageStore.Data[models.Index] = msg
	// parse struct to json and send it back in response
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	// get params and body from the request
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// msg is the updated message
	var msg models.Message
	err = json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// m is the original message
	// update message if found
	var response models.CustomResponse
	w.Header().Set("Content-Type", "application/json")
	if m, ok := models.MessageStore.Data[id]; ok {
		msg.SentAt = m.SentAt
		delete(models.MessageStore.Data, id)
		models.MessageStore.Data[id] = msg
		response.Status = "Success"
		response.Description = "Message has been updated"
		w.WriteHeader(http.StatusOK)
	} else {
		response.Status = "Failed"
		response.Description = "Message has not been found"
		w.WriteHeader(http.StatusNotFound)
	}
	// parse struct to json and send it back in response
	j, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Write(j)
	}
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	// get params from the request
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// delete the message if found
	var response models.CustomResponse
	w.Header().Set("Content-Type", "application/json")
	if _, ok := models.MessageStore.Data[id]; ok {
		delete(models.MessageStore.Data, id)
		response.Status = "Success"
		response.Description = "Message has been deleted"
		w.WriteHeader(http.StatusOK)
	} else {
		response.Status = "Failed"
		response.Description = "Message has not been found"
		w.WriteHeader(http.StatusNotFound)
	}
	// parse struct to json and send it back in response
	j, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Write(j)
	}
}
