package tender_handler

import (
	"avito_tech/internal/model"
	"avito_tech/internal/service/tender_service"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

type TenderHandlerImpl struct {
	tenderService tender_service.TenderService
}

type GetTenderParams struct {
	limit, offset int
	serviceTypes  []model.ServiceType
}

var QueryParamsError = errors.New("query params error")

func NewTenderHandler(tenderService tender_service.TenderService) *TenderHandlerImpl {
	return &TenderHandlerImpl{tenderService: tenderService}
}

func (h *TenderHandlerImpl) Get(c *gin.Context) {
	getTenderParams, err := GetHandlerQueryParams(c)
	if err != nil {
		return
	}

	tenders, err := h.tenderService.Get(getTenderParams.limit, getTenderParams.offset, getTenderParams.serviceTypes)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tenders not found"})
		return
	}

	c.JSON(http.StatusOK, tenders)
}

func (h *TenderHandlerImpl) GetStatus(c *gin.Context) {
	var tenderID uuid.UUID
	err := tenderID.Scan(c.Param("tenderId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tenderId bad format"})
	}

	username := c.Param("username")

	status, err := h.tenderService.GetStatus(tenderID, username)
	if err != nil {
		switch err {
		case model.ErrForbidden:
			c.JSON(http.StatusForbidden, gin.H{"error": "tenders not found"})
		default:
			c.JSON(http.StatusNotFound, gin.H{"error": "tenders not found"})
		}
		return
	}

	c.JSON(http.StatusOK, status)
}

func (h *TenderHandlerImpl) UpdateStatus(c *gin.Context) {
	var tenderID uuid.UUID
	err := tenderID.Scan(c.Param("tenderId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tenderId bad format"})
	}

	username := c.Param("username")
	status := c.Param("tenderStatus")

	tender, err := h.tenderService.UpdateStatus(tenderID, status, username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tenders not found"})
		return
	}

	c.JSON(http.StatusOK, tender)
}

func (h *TenderHandlerImpl) Create(c *gin.Context) {
	var tender model.Tender
	if err := c.ShouldBindJSON(&tender); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.tenderService.Create(&tender); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tender)
}

func (h *TenderHandlerImpl) GetByCreatorUsername(c *gin.Context) {
	username := c.Param("username")

	getTenderParams, err := GetHandlerQueryParams(c)
	if err != nil {
		return
	}

	tenders, err := h.tenderService.GetByCreatorUsername(getTenderParams.limit, getTenderParams.offset, username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tenders not found"})
		return
	}

	c.JSON(http.StatusOK, tenders)
}

func (h *TenderHandlerImpl) Update(c *gin.Context) {
	var tenderID uuid.UUID
	err := tenderID.Scan(c.Param("tenderId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tenderId bad format"})
	}

	var tender model.Tender
	if err := c.ShouldBindJSON(&tender); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tender.ID = tenderID

	if err := h.tenderService.Update(&tender); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tender)
}

func (h *TenderHandlerImpl) Rollback(c *gin.Context) {
	var id uuid.UUID

	err := id.Scan(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	version, err := strconv.Atoi(c.Param("version"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid version"})
		return
	}

	tender, err := h.tenderService.Rollback(id, version)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tender)
}

func GetHandlerQueryParams(c *gin.Context) (*GetTenderParams, error) {
	var err error
	var limit, offset int
	var serviceTypes []model.ServiceType

	limitQueryParam := c.Query("limit")
	if limitQueryParam == "" {
		limit = 5
	} else {
		limit, err = strconv.Atoi(limitQueryParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "limit bad format"})

			return nil, QueryParamsError
		}
	}

	offsetQueryParam := c.Query("offset")
	if offsetQueryParam == "" {
		offset = 0
	} else {
		offset, err = strconv.Atoi(offsetQueryParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "offset bad format"})
			return nil, QueryParamsError
		}
	}

	serviceTypeQueryParamArr := c.QueryArray("service_type")
	for _, serviceType := range serviceTypeQueryParamArr {
		st := model.ServiceType(serviceType)
		if !st.IsValid() {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid service type"})
			return nil, QueryParamsError
		}

		serviceTypes = append(serviceTypes, st)
	}

	return &GetTenderParams{limit, offset, serviceTypes}, nil
}
