package storage

import (
	"errors"
	"sync"
	"task12/internal/model"
	"time"
)

type MapStorage struct{
	data 	map[string] *model.Event
	mutex	sync.Mutex
}

func New() *MapStorage {
	data := make(map[string] *model.Event)
	return &MapStorage{data: data}
}

func (s *MapStorage) Create(event *model.Event) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, ok := s.data[event.ID]; ok{
		return errors.New("event already exists")
	}
	s.data[event.ID] = event
	return nil
}

func (s *MapStorage) Update(event *model.Event) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, ok := s.data[event.ID]; !ok{
		return errors.New("event not exists")
	}
	s.data[event.ID] = event
	return nil
}

func (s *MapStorage) Delete(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, ok := s.data[id]; !ok{
		return errors.New("event not exists")
	}
	delete(s.data, id)
	return nil
}

func (s *MapStorage) GetForDay(date time.Time) []model.Event {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var events []model.Event
	for _, event := range s.data {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() && event.Date.Day() == date.Day() {
			events = append(events, *event)
		}
	}
	return events
}

func (s *MapStorage) GetForWeek(date time.Time) []model.Event {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var events []model.Event
	startOfWeek := date.AddDate(0, 0, -int(date.Weekday()))
	endOfWeek := startOfWeek.AddDate(0, 0, 7)
	for _, event := range s.data {
		if (event.Date.Equal(startOfWeek) || event.Date.After(startOfWeek)) && event.Date.Before(endOfWeek) {
			events = append(events, *event)
		}
	}
	return events
}

func (s *MapStorage) GetForMonth(date time.Time) []model.Event {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var events []model.Event
	startOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0)
	for _, event := range s.data {
		if (event.Date.Equal(startOfMonth) || event.Date.After(startOfMonth)) && event.Date.Before(endOfMonth) {
			events = append(events, *event)
		}
	}
	return events
}