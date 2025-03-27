package dmodel

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID   int64  `db:"id"`
	UUID string `db:"uuid"`
	// fields set by participant himself or auth provider
	FirstName    string    `db:"first_name"`
	LastName     string    `db:"last_name"`
	EmailAddress string    `db:"email_address"`
	Company      string    `db:"company"`
	Description  string    `db:"description"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	// soft delete
	DeletedAt sql.NullTime `db:"deleted_at"`
}

func NewUser(uuid string, firstName, lastName, emailAddress, company, description string) *User {
	return &User{
		UUID:         uuid,
		FirstName:    firstName,
		LastName:     lastName,
		EmailAddress: emailAddress,
		Company:      company,
		Description:  description,
	}
}

func (p User) String() string {
	return fmt.Sprintf("ID:%d, FullName:%s %s", p.ID, p.FirstName, p.LastName)
}
