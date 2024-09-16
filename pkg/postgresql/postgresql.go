package postgresql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func InitDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(true)

	if err := db.DB().Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to PostgreSQL")

	return db
}
