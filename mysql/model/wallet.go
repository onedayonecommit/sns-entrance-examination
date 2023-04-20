package model

type Wallet struct {
    Wallet_address string `gorm:"type:varchar(16);not null;uniqueIndex"`
    UserId        string `gorm:"type:varchar(50);not null;primaryKey;foreignKey:Email"`
    User           User   `gorm:"references:Email"`
}