package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/onedayonecommit/sns/mysql"
	"github.com/onedayonecommit/sns/mysql/model"
)

// 가입된 데이터 없으면 True 있으면 false,password
func UserCheck(userEmail string) (bool,string) {
	fmt.Println(userEmail)
	db := mysql.ConnectDatabase()
	var user model.User
	err:= db.Where("email = ? ",userEmail).First(&user).Error
	close,_ := db.DB()
	defer close.Close() // 함수가 종료되면 연결하였던 DB 연결을 끊어줘야 메모리릭 방지 가능
	if err != nil {
		return true,"" // 전달받은 이메일로 가입된 데이터가 없을시
	}
	return false,user.Password // 전달 받은 이메일로 가입된 데이터가 있다면 패스워드 반환
}

// 전달받은 요청값 검증 실패시 err 성공시 nil
func reqBodyCheck(req *http.Request,body interface{}) error {
	data,err := ioutil.ReadAll(req.Body)

	if err!= nil {
		return err
	}
	
	err = json.Unmarshal([]byte(data),&body)
	
	if err != nil{
		return err
	}
	return nil
}

