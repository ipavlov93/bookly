package model

import (
	"fmt"
	"net/mail"
	"time"
)

type Participant struct {
	ID string // type string to store hash or GUID
	// fields set by participant himself or auth provider
	FirstName    string
	LastName     string
	ContactEmail mail.Address

	User        User
	Description string
	DeletedAt   time.Time // soft delete
}

func NewParticipant(firstName, lastName, email string, userID ...string) *Participant {
	//mail.ParseAddress(email)

	uID := ""
	if len(userID) > 0 {
		uID = userID[0]
	}

	return &Participant{
		FirstName: firstName,
		LastName:  lastName,
		ContactEmail: mail.Address{
			Name:    fmt.Sprintf("%s %s", firstName, lastName),
			Address: email,
		},
		User: User{
			ID: uID,
		},
	}
}

func (p Participant) String() string {
	return fmt.Sprintf("ID:%s, FullName:%s %s", p.ID, p.FirstName, p.LastName)
}

func (p Participant) Equals(participant *Participant) bool {
	if p.ID != participant.ID {
		return false
	}
	return p.ContactEmail.Address == participant.ContactEmail.Address
}
