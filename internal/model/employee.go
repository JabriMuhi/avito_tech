package model

import (
	"github.com/google/uuid"
	"time"
)

type Employee struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `gorm:"unique;not null" json:"username"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
