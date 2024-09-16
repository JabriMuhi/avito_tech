package review_service

import (
	"avito_tech/internal/model"
	"avito_tech/internal/repository/review_repository"
	"errors"
	"github.com/google/uuid"
)

type ReviewServiceImpl struct {
	reviewRepository review_repository.ReviewRepository
}

func NewReviewService(reviewRepository review_repository.ReviewRepository) *ReviewServiceImpl {
	return &ReviewServiceImpl{reviewRepository: reviewRepository}
}

func (s *ReviewServiceImpl) Create(review *model.Review) error {
	if review.Comment == "" || review.AuthorUsername == "" {
		return errors.New("invalid review data")
	}

	return s.reviewRepository.Create(review)
}

func (s *ReviewServiceImpl) GetByBidID(bidID uuid.UUID) ([]*model.Review, error) {
	//валидация
	return s.reviewRepository.GetByBidID(bidID)
}

func (s *ReviewServiceImpl) GetByAuthorUsername(username string) ([]*model.Review, error) {
	//валидация
	return s.reviewRepository.GetByAuthorUsername(username)
}

func (s *ReviewServiceImpl) GetByOrganizationID(organizationID uuid.UUID) ([]*model.Review, error) {
	//валидация
	return s.reviewRepository.GetByOrganizationID(organizationID)
}

func (s *ReviewServiceImpl) GetReviewsByBid(bidID uuid.UUID) ([]*model.Review, error) {
	return s.reviewRepository.GetReviewsByBid(bidID)
}
