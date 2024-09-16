package tender_handler

import "github.com/gin-gonic/gin"

type TenderHandler interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	GetStatus(c *gin.Context)
	UpdateStatus(c *gin.Context)
	GetByCreatorUsername(c *gin.Context)
	Update(c *gin.Context)
	Rollback(c *gin.Context)
}
