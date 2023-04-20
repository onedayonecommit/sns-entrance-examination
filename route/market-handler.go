package route

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/onedayonecommit/sns/mysql"
	"github.com/onedayonecommit/sns/mysql/model"
	"gorm.io/gorm"
)


type ReqAndCon struct {
	Denom string
	Amount int
}

type ExchangeResult struct{
	Request ReqAndCon
	Converted ReqAndCon
};

func ExchangeHandler(res http.ResponseWriter, req *http.Request){
	if req.Method != "POST" {
		http.Error(res,"request method is not allowed",http.StatusMethodNotAllowed)
		return
	}
	err := req.ParseForm()
	if err != nil {
		http.Error(res,"content type is wrong",http.StatusBadRequest)
	}

	var coin model.Coin

	address := req.FormValue("address") // 임시 방편
	fromToken := req.FormValue("fromToken")
	toToken := req.FormValue("toToken")
	amount := req.FormValue("amount")
	

	// 교환 POST 요청시 jwt에 담은 wallet address 토대로 변경
	db := mysql.ConnectDatabase()
	close,_ := db.DB()
	defer close.Close()

	err = db.Where("address = ? ", address).First(&coin).Error
	if err != nil {
		http.Error(res,"this address is not exists",http.StatusBadRequest)
		return
	}

	if fromToken == "BTC" && toToken =="ETH" {
		tx := db.Begin()
		value,err := strconv.Atoi(amount)

		if err != nil{
			http.Error(res,"amount is not number",http.StatusBadRequest)
			return
		}

		if err := tx.Model(&model.Coin{}).
		Where("address = ? ",address).
		Update("BTC",gorm.Expr("BTC - ? ",value)).
		Error; err != nil{
			http.Error(res,"exchange Service is not working",http.StatusInternalServerError)
			tx.Rollback()
			return
		}

		if err := tx.Model(&model.Coin{}).
		Where("address = ? ",address).
		Update("ETH",gorm.Expr("ETH + ?", value*10)).
		Error; err!=nil{
			http.Error(res,"exchange Service is not working",http.StatusInternalServerError)
			tx.Rollback()
			return
		}
		
		tx.Commit()
		exChangeResult := ExchangeResult{
			Request: ReqAndCon{Denom: "BTC",Amount: value},
			Converted: ReqAndCon{Denom: "ETH",Amount: value*10},
		}
	
		res.Header().Set("Content-Type","application/json")
		json.NewEncoder(res).Encode(exChangeResult)
		return
	} else if fromToken == "ETH" && toToken =="BTC" {
		tx := db.Begin()

		value,err := strconv.Atoi(amount)
		if err != nil{
			http.Error(res,"amount is not number",http.StatusBadRequest)
			return
		}

		if err := tx.Model(&model.Coin{}).
		Where("address = ? ",address).
		Update("ETH",gorm.Expr("BTC - ? ",value)).
		Error; err != nil{
			http.Error(res,"exchange Service is not working",http.StatusInternalServerError)
			tx.Rollback()
			return
		}

		if err := tx.Model(&model.Coin{}).
		Where("address = ? ",address).
		Update("BTC",gorm.Expr("ETH + ?", value/10)).
		Error; err!=nil{
			http.Error(res,"exchange Service is not working",http.StatusInternalServerError)
			tx.Rollback()
			return
		}
		
		tx.Commit()
		exChangeResult := ExchangeResult{
			Request: ReqAndCon{Denom: "ETH",Amount: value},
			Converted: ReqAndCon{Denom: "BTC",Amount: value/10},
		}
	
		res.Header().Set("Content-Type","application/json")
		json.NewEncoder(res).Encode(exChangeResult)
		return
	}

	return 
}