package bid_handler

import "github.com/gin-gonic/gin"

type BidHandler interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
	GetByTenderID(c *gin.Context)
	GetByOrganizationID(c *gin.Context)
	GetByCreatorUsername(c *gin.Context)
	Update(c *gin.Context)
	Rollback(c *gin.Context)
}
