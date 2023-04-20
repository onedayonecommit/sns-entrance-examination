package mysql

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/onedayonecommit/sns/mysql/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
  
  

func ConnectDatabase() *gorm.DB {
	err:= godotenv.Load()
	if err != nil {
		log.Fatalln("Env loading failed")
	}
	dbPw:= os.Getenv("DB_PW")
	dsn := fmt.Sprintf("root:%s@tcp(localhost:3306)/sns?charset=utf8mb4&parseTime=True&loc=Local",dbPw) // javascript 에서 `${}`같은 기능
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	err = db.AutoMigrate(&model.User{},&model.User{},&model.Coin{},&model.Transaction{})
	if err != nil {
		log.Fatalf("Error migrating schema: %v", err)
	}
	return db
	
}
