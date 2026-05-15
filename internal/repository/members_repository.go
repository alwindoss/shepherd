package repository

import (
	"database/sql"
	"time"
)

func NewMembersRepository(db *sql.DB) *MembersRepository {
	mr := new(MembersRepository)
	mr.db = db
	return mr
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

type MembersRepository struct {
	db *sql.DB
}

func (mr *MembersRepository) FetchAllMembers() ([]*Member, error) {
	query := "SELECT * FROM members"
	rows, err := mr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var members []*Member
	for rows.Next() {
		m := new(Member)

		err := rows.Scan(&m.ID, &m.FirstName, &m.LastName, &m.Email, &m.Phone, &m.AadharNumber, &m.PAN, &m.Address, &m.RequestedAnnonimity, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
		if err != nil {
			return nil, err
		}
		members = append(members, m)
	}
	return members, nil
}
