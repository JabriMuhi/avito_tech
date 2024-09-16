package bid_repository

import (
	"avito_tech/internal/model"
	"github.com/google/uuid"
)

type BidRepository interface {
	Create(bid *model.Bid) error
	GetByID(id uuid.UUID) (*model.Bid, error)
	GetByTenderID(tenderID uuid.UUID) ([]*model.Bid, error)
	GetByOrganizationID(organizationID uuid.UUID) ([]*model.Bid, error)
	GetByCreatorUsername(username string) ([]*model.Bid, error)
	Update(bid *model.Bid) error
	Rollback(bidID uuid.UUID, version int) error
}
