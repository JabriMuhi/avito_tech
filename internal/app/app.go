package app

import (
	"avito_tech/internal/handler/bid_handler"
	"avito_tech/internal/handler/review_handler"
	"avito_tech/internal/handler/tender_handler"
	"avito_tech/internal/repository/bid_repository"
	"avito_tech/internal/repository/review_repository"
	"avito_tech/internal/repository/tender_repository"
	"avito_tech/internal/service/bid_service"
	"avito_tech/internal/service/review_service"
	"avito_tech/internal/service/tender_service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router        *gin.Engine
	TenderHandler tender_handler.TenderHandler
	BidHandler    bid_handler.BidHandler
	ReviewHandler review_handler.ReviewHandler
}

func NewApp(db *gorm.DB) *App {
	tenderRepository := tender_repository.NewTenderRepositoryGORM(db)
	tenderService := tender_service.NewTenderService(tenderRepository)
	tenderHandler := tender_handler.NewTenderHandler(tenderService)

	bidRepository := bid_repository.NewBidRepositoryGORM(db)
	bidService := bid_service.NewBidService(bidRepository)
	bidHandler := bid_handler.NewBidHandler(bidService)

	reviewRepository := review_repository.NewReviewRepositoryGORM(db)
	reviewService := review_service.NewReviewService(reviewRepository)
	reviewHandler := review_handler.NewReviewHandler(reviewService)

	router := gin.Default()

	return &App{
		Router:        router,
		TenderHandler: tenderHandler,
		BidHandler:    bidHandler,
		ReviewHandler: reviewHandler,
	}
}
