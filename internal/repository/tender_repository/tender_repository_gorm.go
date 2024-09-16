package tender_repository

import (
	"avito_tech/internal/model"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type TenderRepositoryGORM struct {
	db *gorm.DB
}

func NewTenderRepositoryGORM(db *gorm.DB) *TenderRepositoryGORM {
	return &TenderRepositoryGORM{db: db}
}

func (r *TenderRepositoryGORM) Get(limit, offset int, serviceTypes []model.ServiceType) ([]*model.Tender, error) {
	var tenders []*model.Tender

	stypes := make([]string, 0, len(serviceTypes))
	for _, serviceType := range serviceTypes {
		stypes = append(stypes, string(serviceType))
	}

	if len(serviceTypes) != 0 {
		err := r.db.Debug().Where("service_type IN (?)", stypes).Offset(offset).Limit(limit).Order("name ASC").Find(&tenders).Error
		return tenders, err
	}

	err := r.db.Offset(offset).Limit(limit).Order("name ASC").Find(&tenders).Error

	return tenders, err
}

func (r *TenderRepositoryGORM) Create(tender *model.Tender) error {
	tender.ID = uuid.New()

	return r.db.Create(tender).Error
}

func (r *TenderRepositoryGORM) GetByID(tenderID uuid.UUID) (model.Tender, error) {
	var tender model.Tender
	err := r.db.Where("id = ?", tenderID).First(&tender).Error

	return tender, err
}

func (r *TenderRepositoryGORM) GetByCreatorUsername(limit, offset int, username string) ([]*model.Tender, error) {
	var tenders []*model.Tender
	err := r.db.Where("creator_username = ?", username).Offset(offset).Limit(limit).Find(&tenders).Error
	return tenders, err
}

func (r *TenderRepositoryGORM) Update(tender *model.Tender) error {
	var version []int
	err := r.db.Select("version").Where("id = ?", tender.ID).Order("version DESC").Limit(1).Pluck("version", &version).Error
	if err != nil {
		return err
	}

	tender.IsCurrent = false

	err = r.db.Save(tender).Error
	if err != nil {
		return err
	}

	tender.Version = version[0] + 1

	return r.db.Create(tender).Error
}

func (r *TenderRepositoryGORM) Rollback(tenderID uuid.UUID, desiredVersion int) (*model.Tender, error) {
	tender, err := r.GetByID(tenderID)
	if err != nil {
		return nil, err
	}

	tender.IsCurrent = false

	err = r.db.Save(tender).Error
	if err != nil {
		return nil, err
	}

	err = r.db.Where("id = ? AND version = ?", tenderID, desiredVersion).First(&tender).Error
	if err != nil {
		return nil, err
	}

	var version int
	err = r.db.Select("version").Where("id = ?", tender.ID).Order("version DESC").First(&tender).Error
	if err != nil {
		return nil, err
	}

	tender.Version = version + 1

	err = r.db.Create(tender).Error
	if err != nil {
		return nil, err
	}

	return &tender, nil
}
