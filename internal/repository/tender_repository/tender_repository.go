package tender_repository

import (
	"avito_tech/internal/model"
	"github.com/google/uuid"
)

type TenderRepository interface {
	Get(limit, offset int, serviceTypes []model.ServiceType) ([]*model.Tender, error)
	Create(tender *model.Tender) error
	GetByID(tenderID uuid.UUID) (model.Tender, error)
	GetByCreatorUsername(limit, offset int, username string) ([]*model.Tender, error)
	Update(tender *model.Tender) error
	Rollback(uuid2 uuid.UUID, version int) (*model.Tender, error)
}
