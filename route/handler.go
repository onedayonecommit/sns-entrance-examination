package route

import (
	"github.com/onedayonecommit/sns/mysql"
	"github.com/onedayonecommit/sns/mysql/model"
)

// 가입된 데이터 없으면 err 있으면 nil,password
func UserCheck(userEmail string) (error,string) {
	db := mysql.ConnectDatabase()
	var user model.User
	err:= db.Where("email = ? ",userEmail).First(&user).Error
	close,_ := db.DB()
	defer close.Close() // 함수가 종료되면 연결하였던 DB 연결을 끊어줘야 메모리릭 방지 가능
	if err != nil {
		return err,"" // 전달받은 이메일로 가입된 데이터가 없을시
	}
	return nil,user.Password // 전달 받은 이메일로 가입된 데이터가 있다면 패스워드 반환
}

func UserWallet(userEmail string) (string){
	db := mysql.ConnectDatabase()
	var wallet model.Wallet
	close,_ := db.DB()
	defer close.Close()

	err:= db.Where("user_id = ? ",userEmail).First(&wallet).Error
	if err !=nil {
		return ""
	}
	return wallet.Wallet_address
}