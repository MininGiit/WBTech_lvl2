package http

import (
	"errors"
	"task12/internal/model"
	"time"
)

type Response struct{
	Result interface{} `json:"result"`
}

type BadResponse struct{
	Error string `json:"error"`
}

type EventRequest struct {
	ID        string	`json:"id"`
	UserID    string    `json:"user_id"`
	Date      string	`json:"date"`
}

func (e *EventRequest) ToModel() (*model.Event, error) {
	if e.ID == "" || e.UserID == "" || e.Date == "" {
		return nil, errors.New("bad request")
	}
	date, err := time.Parse("2006-01-02", e.Date)
	if err != nil {
		return nil, errors.New("invalid date format")
	}
	return &model.Event{ID: e.ID, UserID: e.UserID, Date: date}, nil
}