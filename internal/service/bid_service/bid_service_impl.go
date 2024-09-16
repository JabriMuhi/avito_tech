package bid_service

import (
	"avito_tech/internal/model"
	"avito_tech/internal/repository/bid_repository"
	"avito_tech/internal/repository/tender_repository"
	"errors"
	"github.com/google/uuid"
)

type BidServiceImpl struct {
	bidRepository    bid_repository.BidRepository
	tenderRepository tender_repository.TenderRepository
}

func NewBidService(bidRepository bid_repository.BidRepository) *BidServiceImpl {
	return &BidServiceImpl{bidRepository: bidRepository}
}

func (s *BidServiceImpl) Create(bid *model.Bid) error {
	//if bid.Name == "" || bid.Description == "" || bid.Status == "" || bid.TenderID == 0 || bid.OrganizationID == 0 || bid.CreatorUsername == "" {
	//	return errors.New("invalid bid data")
	//}

	tender, err := s.tenderRepository.GetByID(bid.TenderID)
	if err != nil {
		return err
	}
	if tender.Status != "PUBLISHED" {
		return errors.New("tender is not published")
	}

	return s.bidRepository.Create(bid)
}

func (s *BidServiceImpl) GetByID(id uuid.UUID) (*model.Bid, error) {
	return s.bidRepository.GetByID(id)
}

func (s *BidServiceImpl) GetByTenderID(tenderID uuid.UUID) ([]*model.Bid, error) {
	return s.bidRepository.GetByTenderID(tenderID)
}

func (s *BidServiceImpl) GetByOrganizationID(organizationID uuid.UUID) ([]*model.Bid, error) {
	return s.bidRepository.GetByOrganizationID(organizationID)
}

func (s *BidServiceImpl) GetByCreatorUsername(username string) ([]*model.Bid, error) {
	return s.bidRepository.GetByCreatorUsername(username)
}

func (s *BidServiceImpl) Update(bid *model.Bid) error {
	if bid.Status == "CANCELED" || bid.Status == "APPROVED" {
		return errors.New("invalid bid status")
	}

	return s.bidRepository.Update(bid)
}

func (s *BidServiceImpl) Rollback(bidID uuid.UUID, version int) error {
	return s.bidRepository.Rollback(bidID, version)
}
