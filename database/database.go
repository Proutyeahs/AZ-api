package database

import (
	"fmt"
	"log"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/proutyeahs/AZ-api/models"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"),
)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("failed to connect to the database. \n", err)
		os.Exit(2)
	}

	log.Println("Connected to the database")
	db.Logger = db.Logger.LogMode(logger.Info)

	log.Println("running migrations...")
	db.AutoMigrate(&models.Tile{})

	DB = Dbinstance{
		Db: db,
	}
}
