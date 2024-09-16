package model

import (
	"github.com/google/uuid"
	"time"
)

type Tender struct {
	ID              uuid.UUID   `json:"id"`
	Name            string      `json:"name"`
	ServiceType     ServiceType `json:"service_type"`
	Description     string      `json:"description"`
	Status          string      `json:"status"`
	OrganizationID  uuid.UUID   `json:"organizationId"`
	CreatorUsername string      `json:"creatorUsername"`
	CreatedAt       time.Time   `json:"createdAt"`
	UpdatedAt       time.Time   `json:"updatedAt"`
	Version         int         `json:"version"`
	IsCurrent       bool        `json:"isCurrent"`
}

type ServiceType string

var (
	Construction = ServiceType("Construction")
	Delivery     = ServiceType("Delivery")
	Manufacture  = ServiceType("Manufacture")
)

func (s ServiceType) IsValid() bool {
	return s == Construction || s == Delivery || s == Manufacture
}
