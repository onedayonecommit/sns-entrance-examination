package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/onedayonecommit/sns/mysql"
	"github.com/onedayonecommit/sns/mysql/model"
)

func UserCheck(userEmail string) bool {
	fmt.Println(userEmail)
	db := mysql.ConnectDatabase()
	err:= db.Where("email = ? ",userEmail).First(&model.User{}).Error
	close,_ := db.DB()
	defer close.Close()
	if err != nil {
		return true
	}
	return false
}

func reqBodyCheck(req *http.Request,s *signupBody) error {
	data,err := ioutil.ReadAll(req.Body)

	if err!= nil {
		return err
	}
	
	err = json.Unmarshal([]byte(data),&s)
	
	if err != nil{
		return err
	}
	return nil
}

