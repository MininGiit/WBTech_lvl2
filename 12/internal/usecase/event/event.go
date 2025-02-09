package event

import (
	"task12/internal/model"
	"task12/internal/repository"
	"time"
)

type UseCaseEvent struct{
	repo	repository.EventRepo
}

func New(repo repository.EventRepo) *UseCaseEvent{
	return &UseCaseEvent{repo: repo}
}

func (uc *UseCaseEvent) Create(event *model.Event) error {
	return uc.repo.Create(event)
}

func (uc *UseCaseEvent) Update(event *model.Event) error {
	return uc.repo.Update(event)
}

func (uc *UseCaseEvent) Delete(id string) error {
	return uc.repo.Delete(id)
}

func (uc *UseCaseEvent) GetForDay(date time.Time) []model.Event {
	return uc.repo.GetForDay(date)
}

func (uc *UseCaseEvent) GetForWeek(date time.Time) []model.Event {
	return uc.repo.GetForWeek(date)
}

func (uc *UseCaseEvent) GetForMonth(date time.Time) []model.Event {
	return uc.repo.GetForMonth(date)
}