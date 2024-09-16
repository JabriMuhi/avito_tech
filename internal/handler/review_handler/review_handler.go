package review_handler

import "github.com/gin-gonic/gin"

type ReviewHandler interface {
	Create(c *gin.Context)
	GetByBidID(c *gin.Context)
	GetByAuthorUsername(c *gin.Context)
	GetByOrganizationID(c *gin.Context)
	GetReviewsByBid(c *gin.Context)
}
