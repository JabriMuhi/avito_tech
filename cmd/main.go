package main

import (
	_ "avito_tech/docs"
	"avito_tech/internal/app"
	"avito_tech/internal/model"
	"avito_tech/pkg/postgresql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load("./docker/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := postgresql.InitDB()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)

	db.AutoMigrate(&model.Employee{}, &model.Organization{}, &model.OrganizationResponsible{}, &model.Bid{}, &model.Review{})

	db.Exec(`CREATE TABLE IF NOT EXISTS tenders
(
    id               uuid not null,
    name             text,
    service_type     text,
    description      text,
    status           text,
    organization_id  uuid,
    creator_username text,
    created_at       timestamp with time zone,
    updated_at       timestamp with time zone,
    version          integer,
    is_current       boolean
);`)

	appBuild := app.NewApp(db)

	api := appBuild.Router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "ok")
		})

		api.StaticFile("/openapi.yml", "docs/openapi.yml")

		bids := api.Group("/bids")
		{
			bids.POST("/new", appBuild.BidHandler.Create)
			bids.GET("/:id", appBuild.BidHandler.GetByID)
			bids.GET("/tenders/:tenderID/list", appBuild.BidHandler.GetByTenderID)
			bids.GET("/organizations/:organizationID", appBuild.BidHandler.GetByOrganizationID)
			bids.GET("/users/:username", appBuild.BidHandler.GetByCreatorUsername)
			//bids.PUT("/:id/status", appBuild.BidHandler.UpdateStatus)
			bids.PATCH("/:id/edit", appBuild.BidHandler.Update)
			bids.PUT("/:id/rollback/:version", appBuild.BidHandler.Rollback)
			//bids.PUT("/:id/submit_decision", appBuild.BidHandler.SubmitDecision)
			//bids.PUT("/:id/feedback", appBuild.BidHandler.Feedback)
		}

		reviews := api.Group("/reviews")
		{
			reviews.POST("/", appBuild.ReviewHandler.Create)
			reviews.GET("/bids/:bidID", appBuild.ReviewHandler.GetByBidID)
			reviews.GET("/users/:username", appBuild.ReviewHandler.GetByAuthorUsername)
			reviews.GET("/organizations/:organizationID", appBuild.ReviewHandler.GetByOrganizationID)
		}

		tenders := api.Group("/tenders")
		{
			tenders.GET("", appBuild.TenderHandler.Get)
			tenders.POST("/new", appBuild.TenderHandler.Create)
			tenders.GET("/my", appBuild.TenderHandler.GetByCreatorUsername)
			tenders.GET("/:tenderId/status", appBuild.TenderHandler.GetStatus)
			tenders.PUT("/:tenderId/status", appBuild.TenderHandler.UpdateStatus)
			tenders.PATCH("/:tenderId/edit", appBuild.TenderHandler.Update)
			tenders.PUT("/:tenderId/rollback/:version", appBuild.TenderHandler.Rollback)
		}
		_, err := os.Stat("docs/openapi.yml")
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("File does not exist.")
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println("File exists.")
		}

		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
			ginSwagger.URL("http://localhost:"+port+"/api/openapi.yml"),
		))
	}

	err = appBuild.Router.Run(":" + port)
	if err != nil {
		return
	}
}
