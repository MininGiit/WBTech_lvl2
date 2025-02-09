package repository

import (
	"task12/internal/model"
	"time"
)

type EventRepo interface{
	Create(event *model.Event) error
	Update(event *model.Event) error
	Delete(id string) error
	GetForDay(date time.Time) []model.Event
	GetForWeek(date time.Time) []model.Event
	GetForMonth(date time.Time) []model.Event
}