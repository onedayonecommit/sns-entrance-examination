package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/onedayonecommit/sns/mysql"
	"github.com/onedayonecommit/sns/mysql/model"
)

func UserCheck(userEmail string) (bool,string) {
	fmt.Println(userEmail)
	db := mysql.ConnectDatabase()
	var user model.User
	err:= db.Where("email = ? ",userEmail).First(&user).Error
	close,_ := db.DB()
	defer close.Close()
	if err != nil {
		return true,""
	}
	return false,user.Password
}

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

