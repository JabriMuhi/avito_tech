package model

import (
	"github.com/google/uuid"
	"time"
)

type Bid struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Status          string    `json:"status"`
	TenderID        uuid.UUID `json:"tenderId"`
	OrganizationID  int       `json:"organizationId"`
	CreatorUsername string    `json:"creatorUsername"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	Version         int       `json:"version"`
}
