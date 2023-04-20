package model

import "time"

type Transaction struct{
	ID uint `gorm:"autoIncrement;primaryKey;coulmn:transaction_id"`
	TransactionId string `gorm:"varchar(32);unique;not null"`
	To string `gorm:"varchar(16);not null"`
	From string `gorm:"varchar(16);not null"`
	Token string `gorm:"varchar(3);not null"`
	Amount uint `gorm:"type:decimal(23,18);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at"`
}