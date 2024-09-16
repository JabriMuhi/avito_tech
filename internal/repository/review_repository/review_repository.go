package review_repository

import (
	"avito_tech/internal/model"
	"github.com/google/uuid"
)

type ReviewRepository interface {
	Create(review *model.Review) error
	GetByBidID(bidID uuid.UUID) ([]*model.Review, error)
	GetByAuthorUsername(username string) ([]*model.Review, error)
	GetByOrganizationID(organizationID uuid.UUID) ([]*model.Review, error)
	GetReviewsByBid(bidID uuid.UUID) ([]*model.Review, error)
}
