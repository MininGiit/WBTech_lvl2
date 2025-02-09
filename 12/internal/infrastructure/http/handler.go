package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task12/internal/usecase"
	"time"
)

type EventHandler struct {
	ucEvent	usecase.IEvent
}

func NewHandler(ucEvent	usecase.IEvent) *EventHandler{
	return &EventHandler{ucEvent: ucEvent,}
}

func (h *EventHandler)InitRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("POST /create_event", LoggerMiddleware(h.CreateEvent))
	router.HandleFunc("POST /update_event", LoggerMiddleware(h.UpdateEvent))
	router.HandleFunc("POST /delete_event", LoggerMiddleware(h.DeleteEvent))
	router.HandleFunc("GET /events_for_day", LoggerMiddleware(h.GetEventsForDay))
	router.HandleFunc("GET /events_for_week", LoggerMiddleware(h.GetEventsForDay))
	router.HandleFunc("GET /events_for_month", LoggerMiddleware(h.GetEventsForDay))
	return router
}

func (h *EventHandler) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	params, err := parseParams(r, "id", "date")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(BadResponse{Error: err.Error()})
		return
	}
	dateStr := params["date"]
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(BadResponse{Error: err.Error()})
		return
	}
	events := h.ucEvent.GetForDay(date)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Result: events})
}

func (h *EventHandler) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	params, err := parseParams(r, "id", "date")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(BadResponse{Error: err.Error()})
		return
	}
	dateStr := params["date"]
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(BadResponse{Error: err.Error()})
		return
	}
	events := h.ucEvent.GetForWeek(date)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Result: events})
}

func (h *EventHandler) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	params, err := parseParams(r, "id", "date")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(BadResponse{Error: err.Error()})
		return
	}
	dateStr := params["date"]
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(BadResponse{Error: err.Error()})
		return
	}
	events := h.ucEvent.GetForMonth(date)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Result: events})
}

func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event EventRequest
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(BadResponse{Error: err.Error()})
		return
	}
	eventModel, err := event.ToModel()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(BadResponse{Error: err.Error()})
		return
	}
	err = h.ucEvent.Create(eventModel)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(BadResponse{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Result: "ok"})
}

func (h *EventHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var event EventRequest
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(BadResponse{Error: err.Error()})
		return
	}
	eventModel, err := event.ToModel()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(BadResponse{Error: err.Error()})
		return
	}
	err = h.ucEvent.Update(eventModel)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(BadResponse{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Result: "ok"})
}

func (h *EventHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	var event EventRequest
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil || event.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(BadResponse{Error: err.Error()})
		return
	}
	err = h.ucEvent.Delete(event.ID)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(BadResponse{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Result: "ok"})
}

func parseParams(r *http.Request, keys ...string) (map[string]string, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	params := make(map[string]string)
	for _, key := range keys {
		value := r.FormValue(key)
		if value == "" {
			return nil, fmt.Errorf("missing parameter: %s", key) 
		}
		params[key] = value
	}
	return params, nil
}
