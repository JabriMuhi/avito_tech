package bid_repository

import (
	"avito_tech/internal/model"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type BidRepositoryGORM struct {
	db *gorm.DB
}

func NewBidRepositoryGORM(db *gorm.DB) *BidRepositoryGORM {
	return &BidRepositoryGORM{db: db}
}

func (r *BidRepositoryGORM) Create(bid *model.Bid) error {
	return r.db.Create(bid).Error
}

func (r *BidRepositoryGORM) GetByID(id uuid.UUID) (*model.Bid, error) {
	var bid model.Bid
	err := r.db.Where("id = ?", id).First(&bid).Error
	return &bid, err
}

func (r *BidRepositoryGORM) GetByTenderID(tenderID uuid.UUID) ([]*model.Bid, error) {
	var bids []*model.Bid
	err := r.db.Where("tender_id = ?", tenderID).Find(&bids).Error
	return bids, err
}

func (r *BidRepositoryGORM) GetByOrganizationID(organizationID uuid.UUID) ([]*model.Bid, error) {
	var bids []*model.Bid
	err := r.db.Where("organization_id = ?", organizationID).Find(&bids).Error
	return bids, err
}

func (r *BidRepositoryGORM) GetByCreatorUsername(username string) ([]*model.Bid, error) {
	var bids []*model.Bid
	err := r.db.Where("creator_username = ?", username).Find(&bids).Error
	return bids, err
}

func (r *BidRepositoryGORM) Update(bid *model.Bid) error {
	return r.db.Save(bid).Error
}

func (r *BidRepositoryGORM) Rollback(bidID uuid.UUID, version int) error {
	var bid model.Bid
	err := r.db.Where("id = ? AND version = ?", bidID, version).First(&bid).Error
	if err != nil {
		return err
	}
	bid.Version++
	return r.db.Save(&bid).Error
}
