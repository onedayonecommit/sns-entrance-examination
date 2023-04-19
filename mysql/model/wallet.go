package model

type Wallet struct {
	WalletAddress   string `gorm:"type:varchar(16);not null;uniqueIndex"`
	UserID    uint   `gorm:"unique;not null"`
	User      User   `gorm:"foreignKey:UserID;primaryKey;references:Id"`
}