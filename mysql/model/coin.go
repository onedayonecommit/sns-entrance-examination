package model

type Coin struct{
	ETH uint `gorm:"type:decimal(23,18);not null;default:100"`
	BTC uint `gorm:"type:decimal(23,18);not null;default:100"`
	Address string `gorm:"type:varchar(16);not null"`
	Wallet Wallet `gorm:"foreignKey:Address;references:Wallet_address"`
}