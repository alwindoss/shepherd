package handler

import (
	"fmt"
	"net/http"
	"shepherd/templates/pages/members"

	"github.com/a-h/templ"
)

type MemberService interface {
}

func NewMemberHandler(svc MemberService) *MemberHandler {
	mh := new(MemberHandler)

	return mh
}

type MemberHandler struct {
	svc MemberService
}

func (mh *MemberHandler) Index(w http.ResponseWriter, r *http.Request) {
	m1 := members.Member{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Role:      "Treasurer",
		Email:     "john.doe@email.com",
	}
	m2 := members.Member{
		ID:        "2",
		FirstName: "Alice",
		LastName:  "Johnson",
		Role:      "President",
		Email:     "alice.johnson@email.com",
	}
	m3 := members.Member{
		ID:        "3",
		FirstName: "Bob",
		LastName:  "Smith",
		Role:      "Secretary",
		Email:     "bob.smith@email.com",
	}
	var mbrs []members.Member
	mbrs = append(mbrs, m1)
	mbrs = append(mbrs, m2)
	mbrs = append(mbrs, m3)
	templ.Handler(members.Index(mbrs)).ServeHTTP(w, r)
}
func (mh *MemberHandler) New(w http.ResponseWriter, r *http.Request) {
	m1 := members.Member{
		ID:        "1",
		FirstName: "Alwin",
		LastName:  "Doss",
		Role:      "Treasurer",
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
		Role:      "Treasurer",
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
		Role:      "Treasurer",
		Email:     "alwin@email.com",
	}
	templ.Handler(members.Edit(m1)).ServeHTTP(w, r)
}
func (mh *MemberHandler) UpdateFull(w http.ResponseWriter, r *http.Request)    {}
func (mh *MemberHandler) UpdatePartial(w http.ResponseWriter, r *http.Request) {}
func (mh *MemberHandler) Delete(w http.ResponseWriter, r *http.Request)        {}
