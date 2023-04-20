package route

import (
	"fmt"
	"net/http"

	"github.com/onedayonecommit/sns/mysql"
	"github.com/onedayonecommit/sns/mysql/model"
	"github.com/onedayonecommit/sns/util"
)

func SignupHandler(res http.ResponseWriter, req *http.Request){
	if req.Method !="POST"{
		http.Error(res,"request method is not allowed",http.StatusMethodNotAllowed) // 허가 되지 않은 요청방식은 반환
		return
	}

	err:= req.ParseForm()
	if err != nil{
		http.Error(res,"contentType is wrong",http.StatusBadRequest)
		return
	}

	email:= req.FormValue("email")
	pw := req.FormValue("password")
	fullname:= req.FormValue("fullname")

	err,_ = UserCheck(*&email)
	if err != nil{
		db:= mysql.ConnectDatabase()
		tx := db.Begin()

		err := tx.Error
		if err != nil {
			fmt.Println("transaction failed")
			http.Error(res,"Transaction Failed",http.StatusInternalServerError)
			return
		}

		hashPw,err := util.GenerateHashPw(*&pw)
		if err != nil{
			fmt.Println("password hashing failed")
			http.Error(res,"password hasing is failed",http.StatusInternalServerError)
			return
		}

		err = tx.Create(&model.User{Email: *&email, Password: hashPw, FullName: *&fullname}).Error
		if err != nil {
			fmt.Println("user create is failed")
			http.Error(res,"user create is failed",http.StatusBadRequest)
			tx.Rollback()
			return
		}
		
		var hex,_ = util.GenerateHex(7)

		err = tx.Create(&model.Wallet{Wallet_address: hex,UserId: *&email}).Error
		if err != nil{
			fmt.Println("wallet address create is failed")
			http.Error(res,"wallet create is failed",http.StatusBadRequest)
			tx.Rollback()
			return
		}

		err = tx.Create(&model.Coin{Address: hex}).Error
		if err != nil{
			fmt.Println("coin wallet create is failed")
			http.Error(res,"coin airdrop is failed",http.StatusBadRequest)
			tx.Rollback()
			return
		}

		err = tx.Commit().Error
		if err != nil {
			fmt.Println("final transaction commit is failed")
			http.Error(res,"final create is failed",http.StatusInternalServerError)
			tx.Rollback()
			return
		}


		res.WriteHeader(http.StatusOK)
		fmt.Fprint(res,"signup successful")

		close,_ := db.DB()
		defer close.Close()

		return 
	}
	
	http.Error(res,"this Email already is exists",http.StatusBadRequest)
	return	
}