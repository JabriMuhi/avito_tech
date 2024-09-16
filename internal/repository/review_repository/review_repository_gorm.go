package review_repository

import (
	"avito_tech/internal/model"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type ReviewRepositoryGORM struct {
	db *gorm.DB
}

func NewReviewRepositoryGORM(db *gorm.DB) *ReviewRepositoryGORM {
	return &ReviewRepositoryGORM{db: db}
}

func (r *ReviewRepositoryGORM) Create(review *model.Review) error {
	return r.db.Create(review).Error
}

func (r *ReviewRepositoryGORM) GetByBidID(bidID uuid.UUID) ([]*model.Review, error) {
	var reviews []*model.Review
	err := r.db.Where("bid_id = ?", bidID).Find(&reviews).Error
	return reviews, err
}

func (r *ReviewRepositoryGORM) GetByAuthorUsername(username string) ([]*model.Review, error) {
	var reviews []*model.Review
	err := r.db.Where("author_username = ?", username).Find(&reviews).Error
	return reviews, err
}

func (r *ReviewRepositoryGORM) GetByOrganizationID(organizationID uuid.UUID) ([]*model.Review, error) {
	var reviews []*model.Review
	err := r.db.Where("organization_id = ?", organizationID).Find(&reviews).Error
	return reviews, err
}

func (r *ReviewRepositoryGORM) GetReviewsByBid(bidID uuid.UUID) ([]*model.Review, error) {
	var reviews []*model.Review
	if err := r.db.Where("bid_id = ?", bidID).Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}
