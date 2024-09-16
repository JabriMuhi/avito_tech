package model

import "github.com/google/uuid"

type OrganizationResponsible struct {
	ID             uuid.UUID    `json:"id"`
	OrganizationID uuid.UUID    `json:"organizationId"`
	UserID         uuid.UUID    `json:"userId"`
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
	User           Employee     `gorm:"foreignKey:UserID"`
}
