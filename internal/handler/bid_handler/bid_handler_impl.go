package bid_handler

import (
	"avito_tech/internal/model"
	"avito_tech/internal/service/bid_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

type BidHandlerImpl struct {
	bidService bid_service.BidService
}

func NewBidHandler(bidService bid_service.BidService) *BidHandlerImpl {
	return &BidHandlerImpl{bidService: bidService}
}

func (h *BidHandlerImpl) Create(c *gin.Context) {
	var bid model.Bid
	if err := c.ShouldBindJSON(&bid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.bidService.Create(&bid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, bid)
}

func (h *BidHandlerImpl) GetByID(c *gin.Context) {
	var id uuid.UUID

	err := id.Scan(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid bid id"})
		return
	}

	bid, err := h.bidService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "bid not found"})
		return
	}

	c.JSON(http.StatusOK, bid)
}

func (h *BidHandlerImpl) GetByTenderID(c *gin.Context) {
	var tenderID uuid.UUID

	err := tenderID.Scan(c.Param("tenderID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tender id"})
		return
	}

	bids, err := h.bidService.GetByTenderID(tenderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "bids not found"})
		return
	}

	c.JSON(http.StatusOK, bids)
}

func (h *BidHandlerImpl) GetByOrganizationID(c *gin.Context) {
	var organizationID uuid.UUID
	err := organizationID.Scan(c.Param("organizationID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization id"})
		return
	}

	bids, err := h.bidService.GetByOrganizationID(organizationID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "bids not found"})
		return
	}

	c.JSON(http.StatusOK, bids)
}

func (h *BidHandlerImpl) GetByCreatorUsername(c *gin.Context) {
	username := c.Param("username")

	bids, err := h.bidService.GetByCreatorUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "bids not found"})
		return
	}

	c.JSON(http.StatusOK, bids)
}

func (h *BidHandlerImpl) Update(c *gin.Context) {
	var id uuid.UUID

	err := id.Scan(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid bid id"})
		return
	}

	var bid model.Bid
	if err := c.ShouldBindJSON(&bid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.bidService.Update(&bid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bid)
}

func (h *BidHandlerImpl) Rollback(c *gin.Context) {
	var id uuid.UUID

	err := id.Scan(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid bid id"})
		return
	}

	version, err := strconv.Atoi(c.Param("version"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid version"})
		return
	}

	if err := h.bidService.Rollback(uuid.UUID{}, version); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bid, err := h.bidService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "bid not found"})
		return
	}

	c.JSON(http.StatusOK, bid)
}
