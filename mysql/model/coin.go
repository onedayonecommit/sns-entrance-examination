package model

type Coin struct{
	ETH uint `gorm:"not null;default:100"`
	BTC uint `gorm:"not null;default:100"`
	Address string `gorm:"type:varchar(18);not null"`
	Wallet Wallet `gorm:"foreignKey:Address;references:Wallet_address"`
}