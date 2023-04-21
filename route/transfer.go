package route

import (
	"net/http"
	"strconv"

	"github.com/onedayonecommit/sns/mysql"
	"github.com/onedayonecommit/sns/mysql/model"
	"github.com/onedayonecommit/sns/util"
	"gorm.io/gorm"
)

func TransferHandler(res http.ResponseWriter, req *http.Request){
	if req.Method != "POST" {
		http.Error(res,"request method is not allowed",http.StatusMethodNotAllowed)
		return
	}

	err := req.ParseForm()
	if err != nil {
		http.Error(res,"contentType is wrong",http.StatusBadRequest)
		return
	}
	var coin model.Coin
	var coin2 model.Coin

	cookie,err := req.Cookie("koa:sess")
	if err != nil{
		if err == http.ErrNoCookie{
			http.Error(res,"you don't have cookie",http.StatusBadRequest)
			return
		}
	}
	tokenValue := cookie.Value
	from,err := util.VerifyJwt(tokenValue)
	if err.Error() == "jwt is not found"{
		http.Error(res,"jwt already exp, re login",http.StatusBadRequest)
		return
	} else if err != nil{
		http.Error(res,"unknown Error",http.StatusInternalServerError)
		return
	}

	to := req.FormValue("to")
	token := req.FormValue("token")
	amount := req.FormValue("amount")
	db := mysql.ConnectDatabase()
	
	close,_ := db.DB()
	defer close.Close()
	
	value,err := strconv.Atoi(amount)
	if err != nil {
		http.Error(res,"amount is not number",http.StatusBadRequest)
		return
	}

	err = db.Select("BTC,ETH").Where("address = ?",from).First(&coin).Error
	if err != nil {
		http.Error(res,"this wallet address not exists",http.StatusBadRequest)
		return
	}

	err = db.Where("address = ?",to).First(&coin2).Error // 전달받을 수 있는 지갑주소인지 확인
	if err != nil{
		http.Error(res,"this wallet address not exists",http.StatusBadRequest)
		return
	}

	if token == "BTC" && coin.BTC >=uint(value) {
		tx:= db.Begin()
		close,_ := db.DB()
		defer close.Close()
		
		err := tx.Model(&model.Coin{}).Where("address = ?",from).Update("BTC",gorm.Expr("BTC - ?",uint(value))).Error
		if err !=nil {
			http.Error(res,"transfer service is failed",http.StatusInternalServerError)
			tx.Rollback()
			return
		}
		
		err = tx.Model(&model.Coin{}).Where("address = ?",to).Update("BTC",gorm.Expr("BTC + ?",uint(value))).Error
		if err != nil {
			http.Error(res,"transfer service is failed",http.StatusInternalServerError)
			tx.Rollback()
			return
		}
		
		err = tx.Create(&model.Transaction{To: to,From: from,Coin: token,Amount: uint(value) }).Error
		if err != nil {
			http.Error(res,"transaction create failed",http.StatusInternalServerError)
			tx.Rollback()
			return
		}

		err = tx.Commit().Error
		if err != nil {
			http.Error(res,"transaction save failed",http.StatusInternalServerError)
			tx.Rollback()
			return
		}
		res.WriteHeader(http.StatusOK)
		return
	} else if token == "ETH" && coin.ETH >=uint(value) {
		tx:= db.Begin()
		close,_ := db.DB()
		defer close.Close()
		
		err := tx.Model(&model.Coin{}).Where("address = ?",from).Update("ETH",gorm.Expr("ETH - ?",uint(value))).Error
		if err !=nil {
			http.Error(res,"transfer service is failed",http.StatusInternalServerError)
			tx.Rollback()
			return
		}
		
		err = tx.Model(&model.Coin{}).Where("address = ?",to).Update("ETH",gorm.Expr("ETH + ?",uint(value))).Error
		if err != nil {
			http.Error(res,"transfer service is failed",http.StatusInternalServerError)
			tx.Rollback()
			return
		}
		
		err = tx.Create(&model.Transaction{To: to,From: from,Coin: token,Amount: uint(value) }).Error
		if err != nil {
			http.Error(res,"transaction create failed",http.StatusInternalServerError)
			tx.Rollback()
			return
		}

		err = tx.Commit().Error
		if err != nil {
			http.Error(res,"transaction save failed",http.StatusInternalServerError)
			tx.Rollback()
			return
		}
		res.WriteHeader(http.StatusOK)
		return
	}
	return
}