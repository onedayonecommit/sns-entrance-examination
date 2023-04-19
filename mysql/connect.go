package mysql

import (
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id        uint   `gorm:"primaryKey;autoIncrement"`
	Email     string `gorm:"not null"`
	Password  string `gorm:"not null"`
	FullName  string `gorm:"not null"`
}
  
type Wallet struct {
	WalletAddress   string `gorm:"type:varchar(16);not null;uniqueIndex"`
	UserID    uint   `gorm:"unique;not null"`
	User      User   `gorm:"foreignKey:UserID;primaryKey;references:Id"`
}
  
type Coin struct{
	ETH uint `gorm:"not null;default:100"`
	BTC uint `gorm:"not null;default:100"`
	Address string `gorm:"type:varchar(16);not null"`
	Wallet Wallet `gorm:"foreignKey:Address;references:WalletAddress"`
}

func ConnectDatabase() *sql.DB {
	dsn := "root:1q2w3e4r5t!@tcp(localhost:3306)/sns?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	err = db.AutoMigrate(&User{},&Wallet{},&Coin{})
	if err != nil {
		log.Fatalf("Error migrating schema: %v", err)
	}
	dbClose, _ := db.DB()
	return dbClose
}
