package models

import "time"

type CustomResponse struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}

type Message struct {
	Author  string    `json:"author"`
	Content string    `json:"content"`
	SentAt  time.Time `json:"sent_at"`
}

type Database struct {
	Data map[int]Message
}

var MessageStore Database = Database{
	Data: make(map[int]Message),
}

var Index int
