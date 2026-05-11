package handler

import "net/http"

type EventService interface {
}

func NewEventHandler(svc EventService) *EventHandler {
	mh := new(EventHandler)

	return mh
}

type EventHandler struct {
	svc EventService
}

func (mh *EventHandler) Index(w http.ResponseWriter, r *http.Request)         {}
func (mh *EventHandler) New(w http.ResponseWriter, r *http.Request)           {}
func (mh *EventHandler) Create(w http.ResponseWriter, r *http.Request)        {}
func (mh *EventHandler) Show(w http.ResponseWriter, r *http.Request)          {}
func (mh *EventHandler) Edit(w http.ResponseWriter, r *http.Request)          {}
func (mh *EventHandler) UpdateFull(w http.ResponseWriter, r *http.Request)    {}
func (mh *EventHandler) UpdatePartial(w http.ResponseWriter, r *http.Request) {}
func (mh *EventHandler) Delete(w http.ResponseWriter, r *http.Request)        {}
