package model

type User struct {
	Id        uint   `gorm:"primaryKey;autoIncrement"` // 고유키 유저의 인덱스 번호
	Email     string `gorm:"type:varchar(50);not null"` // 이메일
	Password  string `gorm:"type:varchar(30);not null"` // 비밀번호
	FullName  string `gorm:"type:varchar(16);not null"` // 닉네임
}