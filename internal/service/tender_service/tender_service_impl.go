package tender_service

import (
	"avito_tech/internal/model"
	"avito_tech/internal/repository/tender_repository"
	"errors"
	"github.com/google/uuid"
)

type TenderServiceImpl struct {
	tenderRepository tender_repository.TenderRepository
}

func NewTenderService(tenderRepository tender_repository.TenderRepository) *TenderServiceImpl {
	return &TenderServiceImpl{tenderRepository: tenderRepository}
}

func (s *TenderServiceImpl) Get(limit, offset int, serviceTypes []model.ServiceType) ([]*model.Tender, error) {
	return s.tenderRepository.Get(limit, offset, serviceTypes)
}

func (s *TenderServiceImpl) Create(tender *model.Tender) error {
	if tender.Name == "" || tender.Description == "" || tender.CreatorUsername == "" {
		return errors.New("invalid tender data")
	}

	return s.tenderRepository.Create(tender)
}

func (s *TenderServiceImpl) GetStatus(tenderID uuid.UUID, username string) (string, error) {
	tender, err := s.tenderRepository.GetByID(tenderID)
	if err != nil {
		return "", err
	}

	if tender.CreatorUsername != username {
		return "", errors.New("user not exist")
	}

	return tender.Status, nil
}

func (s *TenderServiceImpl) UpdateStatus(tenderID uuid.UUID, status, username string) (*model.Tender, error) {
	tender, err := s.tenderRepository.GetByID(tenderID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, model.ErrNotFound
		}

		return nil, err
	}

	if tender.CreatorUsername != username {
		return nil, model.ErrForbidden
	}

	tender.Status = status

	err = s.tenderRepository.Update(&tender)
	if err != nil {
		return nil, err
	}

	return &tender, nil
}

func (s *TenderServiceImpl) GetByCreatorUsername(limit, offset int, username string) ([]*model.Tender, error) {
	return s.tenderRepository.GetByCreatorUsername(limit, offset, username)
}

func (s *TenderServiceImpl) Update(tender *model.Tender) error {
	return s.tenderRepository.Update(tender)
}

func (s *TenderServiceImpl) Rollback(tenderID uuid.UUID, version int) (*model.Tender, error) {
	return s.tenderRepository.Rollback(tenderID, version)
}
