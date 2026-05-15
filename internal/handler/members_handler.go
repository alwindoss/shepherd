package handler

import (
	"fmt"
	"net/http"
	"shepherd/internal/service"
	"shepherd/templates/pages/members"
	"strconv"

	"github.com/a-h/templ"
)

type MemberService interface {
	FetchAllMembers() ([]*service.Member, error)
}

func NewMemberHandler(svc MemberService) *MemberHandler {
	mh := new(MemberHandler)
	mh.svc = svc
	return mh
}

type MemberHandler struct {
	svc MemberService
}

func (mh *MemberHandler) Index(w http.ResponseWriter, r *http.Request) {
	svcMembers, err := mh.svc.FetchAllMembers()
	if err != nil {
		return
	}
	var mbrs []members.Member
	for _, m := range svcMembers {
		mem := members.Member{
			ID:                 strconv.Itoa(m.ID),
			FirstName:          m.FirstName,
			LastName:           m.LastName,
			Email:              m.Email,
			Phone:              m.Phone,
			AadharNumber:       m.AadharNumber,
			PAN:                m.PAN,
			Address:            m.Address,
			RequestedAnonymity: m.RequestedAnnonimity,
		}
		mbrs = append(mbrs, mem)
	}
	templ.Handler(members.Index(mbrs)).ServeHTTP(w, r)
}
func (mh *MemberHandler) New(w http.ResponseWriter, r *http.Request) {
	m1 := members.Member{
		ID:        "1",
		FirstName: "Alwin",
		LastName:  "Doss",
		Email:     "alwin@email.com",
	}
	templ.Handler(members.New(m1)).ServeHTTP(w, r)
}
func (mh *MemberHandler) Create(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	emailAddress := r.FormValue("emailAddress")
	phoneNumber := r.FormValue("phoneNumber")
	aadharNumber := r.FormValue("aadharNumber")
	pan := r.FormValue("pan")
	address := r.FormValue("address")
	requestedAnonymity := r.FormValue("requestedAnonymity")

	fmt.Println("#######################################")
	fmt.Println("#######################################")
	fmt.Printf("Received form data: First Name: %s, Last Name: %s, Email: %s, Phone: %s, Aadhar Number: %s, PAN: %s, Address: %s, Requested Anonymity: %s\n",
		firstName, lastName, emailAddress, phoneNumber, aadharNumber, pan, address, requestedAnonymity)
}

func (mh *MemberHandler) Show(w http.ResponseWriter, r *http.Request) {
	m1 := members.Member{
		ID:        "1",
		FirstName: "Alwin",
		LastName:  "Doss",
		Email:     "alwin@email.com",
	}
	deletePath := fmt.Sprintf("/members/%s", m1.ID)
	templ.Handler(members.Show("Member Details", m1, deletePath)).ServeHTTP(w, r)
}
func (mh *MemberHandler) Edit(w http.ResponseWriter, r *http.Request) {
	m1 := members.Member{
		ID:        "1",
		FirstName: "Alwin",
		LastName:  "Doss",
		Email:     "alwin@email.com",
	}
	templ.Handler(members.Edit(m1)).ServeHTTP(w, r)
}
func (mh *MemberHandler) UpdateFull(w http.ResponseWriter, r *http.Request)    {}
func (mh *MemberHandler) UpdatePartial(w http.ResponseWriter, r *http.Request) {}
func (mh *MemberHandler) Delete(w http.ResponseWriter, r *http.Request)        {}
