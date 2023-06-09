package model

import "time"

type Transaction struct{
	ID uint `gorm:"autoIncrement;primaryKey;coulmn:transaction_id"`
	To string `gorm:"type:varchar(18);not null"`
	From string `gorm:"type:varchar(18);not null"`
	Coin string `gorm:"type:varchar(3);not null"`
	Amount uint `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at"`
}