package handler

import "net/http"

type DonationService interface {
}

func NewDonationHandler(svc DonationService) *DonationHandler {
	mh := new(DonationHandler)

	return mh
}

type DonationHandler struct {
	svc DonationService
}

func (mh *DonationHandler) Index(w http.ResponseWriter, r *http.Request)         {}
func (mh *DonationHandler) New(w http.ResponseWriter, r *http.Request)           {}
func (mh *DonationHandler) Create(w http.ResponseWriter, r *http.Request)        {}
func (mh *DonationHandler) Show(w http.ResponseWriter, r *http.Request)          {}
func (mh *DonationHandler) Edit(w http.ResponseWriter, r *http.Request)          {}
func (mh *DonationHandler) UpdateFull(w http.ResponseWriter, r *http.Request)    {}
func (mh *DonationHandler) UpdatePartial(w http.ResponseWriter, r *http.Request) {}
func (mh *DonationHandler) Delete(w http.ResponseWriter, r *http.Request)        {}
