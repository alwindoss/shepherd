package service

import (
	"shepherd/internal/repository"
	"time"
)

type MembersRepository interface {
	FetchAllMembers() ([]*repository.Member, error)
}

func NewMemberService(repo MembersRepository) *MemberService {
	ms := new(MemberService)
	ms.repo = repo
	return ms
}

type Member struct {
	ID                  int
	FirstName           string
	LastName            string
	Email               string
	Phone               string
	AadharNumber        string
	PAN                 string
	Address             string
	RequestedAnnonimity bool
	CreatedAt           *time.Time
	UpdatedAt           *time.Time
	DeletedAt           *time.Time
}

type MemberService struct {
	repo MembersRepository
}

func (ms MemberService) FetchAllMembers() ([]*Member, error) {
	repoMembers, err := ms.repo.FetchAllMembers()
	if err != nil {
		return nil, err
	}
	var members []*Member
	for _, m := range repoMembers {
		mem := new(Member)
		mem.ID = m.ID
		mem.FirstName = m.FirstName
		mem.LastName = m.LastName
		mem.Email = m.Email
		mem.Phone = m.Phone
		mem.AadharNumber = m.AadharNumber
		mem.PAN = m.PAN
		mem.Address = m.Address
		mem.RequestedAnnonimity = m.RequestedAnnonimity
		mem.CreatedAt = m.CreatedAt
		mem.UpdatedAt = m.UpdatedAt
		mem.DeletedAt = m.DeletedAt
		members = append(members, mem)
	}
	return members, nil
}
