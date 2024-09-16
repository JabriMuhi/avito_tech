package model

import (
	"github.com/google/uuid"
	"time"
)

type OrganizationType string

const (
	IE  OrganizationType = "IE"
	LLC OrganizationType = "LLC"
	JSC OrganizationType = "JSC"
)

type Organization struct {
	ID          uuid.UUID                 `json:"id"`
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Type        OrganizationType          `json:"type"`
	CreatedAt   time.Time                 `json:"createdAt"`
	UpdatedAt   time.Time                 `json:"updatedAt"`
	Responsible []OrganizationResponsible `gorm:"foreignKey:OrganizationID"`
}
