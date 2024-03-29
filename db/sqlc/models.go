// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	RoleRegistration Role = "Registration"
	RoleSupport      Role = "Support"
)

func (e *Role) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Role(s)
	case string:
		*e = Role(s)
	default:
		return fmt.Errorf("unsupported scan type for Role: %T", src)
	}
	return nil
}

type NullRole struct {
	Role  Role
	Valid bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRole) Scan(value interface{}) error {
	if value == nil {
		ns.Role, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Role.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.Role, nil
}

type Admin struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"fullName"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Role     Role      `json:"role"`
}

type Contestant struct {
	ID               uuid.UUID    `json:"id"`
	FullName         string       `json:"fullName"`
	Email            string       `json:"email"`
	Position         string       `json:"position"`
	RegisteredAt     sql.NullTime `json:"registeredAt"`
	Description      string       `json:"description"`
	Region           string       `json:"region"`
	EthereumAddress  string       `json:"ethereumAddress"`
	NationalIDNumber int64        `json:"nationalIDNumber"`
	ImageAddress     string       `json:"imageAddress"`
}

type User struct {
	ID              uuid.UUID      `json:"id"`
	FullName        string         `json:"fullName"`
	Email           string         `json:"email"`
	Password        string         `json:"password"`
	IsAdmin         sql.NullBool   `json:"isAdmin"`
	ImageAddress    sql.NullString `json:"imageAddress"`
	VotedPresident  sql.NullBool   `json:"votedPresident"`
	VotedGovernor   sql.NullBool   `json:"votedGovernor"`
	RegisteredVoter sql.NullBool   `json:"registeredVoter"`
}

type Voter struct {
	ID               uuid.UUID    `json:"id"`
	FullName         string       `json:"fullName"`
	Email            string       `json:"email"`
	RegisteredAt     time.Time    `json:"registeredAt"`
	Voted            sql.NullBool `json:"voted"`
	Verified         sql.NullBool `json:"verified"`
	NationalIDNumber int64        `json:"nationalIDNumber"`
	ImageAddress     string       `json:"imageAddress"`
	EthereumAddress  string       `json:"ethereumAddress"`
	Region           string       `json:"region"`
}
