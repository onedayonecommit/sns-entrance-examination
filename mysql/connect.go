package mysql

import (
	"database/sql"
	"log"

	"github.com/onedayonecommit/sns/mysql/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
  
  

func ConnectDatabase() *sql.DB {
	dsn := "root:1q2w3e4r5t!@tcp(localhost:3306)/sns?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	err = db.AutoMigrate(&model.User{},&model.User{},&model.Coin{})
	if err != nil {
		log.Fatalf("Error migrating schema: %v", err)
	}
	dbClose, _ := db.DB()
	return dbClose
}
