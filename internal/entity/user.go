package entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FullName  string    `json:"fullName"`
	Birthdate time.Time `json:"birthdate"`
}
