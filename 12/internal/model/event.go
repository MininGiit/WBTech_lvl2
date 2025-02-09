package model

import "time"

type Event struct {
	ID        string	`json:"id"`
	UserID    string    `json:"user_id"`
	Date      time.Time `json:"date"`
}

func NewEvent(id, userId string, date time.Time) *Event{
	return &Event{
		ID: id,
		UserID: userId,
		Date: date,
	}	
}