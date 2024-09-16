package tender_service

import (
	"avito_tech/internal/model"
	"github.com/google/uuid"
)

type TenderService interface {
	Get(limit, offset int, serviceTypes []model.ServiceType) ([]*model.Tender, error)
	Create(tender *model.Tender) error
	GetStatus(tenderID uuid.UUID, username string) (string, error)
	UpdateStatus(tenderID uuid.UUID, status, username string) (*model.Tender, error)
	GetByCreatorUsername(limit, offset int, username string) ([]*model.Tender, error)
	Update(tender *model.Tender) error
	Rollback(tenderID uuid.UUID, version int) (*model.Tender, error)
}
