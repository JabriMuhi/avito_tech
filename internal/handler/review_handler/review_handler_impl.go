package review_handler

import (
	"avito_tech/internal/model"
	"avito_tech/internal/service/review_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type ReviewHandlerImpl struct {
	reviewService review_service.ReviewService
}

func NewReviewHandler(reviewService review_service.ReviewService) *ReviewHandlerImpl {
	return &ReviewHandlerImpl{reviewService: reviewService}
}

func (h *ReviewHandlerImpl) Create(c *gin.Context) {
	var review model.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.reviewService.Create(&review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, review)
}

func (h *ReviewHandlerImpl) GetByBidID(c *gin.Context) {
	var bidID uuid.UUID

	err := bidID.Scan(c.Param("bidID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid bid id"})
		return
	}

	reviews, err := h.reviewService.GetByBidID(bidID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "reviews not found"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

func (h *ReviewHandlerImpl) GetByAuthorUsername(c *gin.Context) {
	username := c.Param("username")

	reviews, err := h.reviewService.GetByAuthorUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "reviews not found"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

func (h *ReviewHandlerImpl) GetByOrganizationID(c *gin.Context) {

	var organizationID uuid.UUID

	err := organizationID.Scan(c.Param("organizationID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization id"})
		return
	}

	reviews, err := h.reviewService.GetByOrganizationID(organizationID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "reviews not found"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

func (h *ReviewHandlerImpl) GetReviewsByBid(c *gin.Context) {
	var bidID uuid.UUID

	err := bidID.Scan(c.Param("bidID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bid not found"})
	}

}
