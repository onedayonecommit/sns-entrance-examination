package route

import (
	"net/http"
	"strconv"

	"github.com/onedayonecommit/sns/mysql"
	"github.com/onedayonecommit/sns/mysql/model"
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

	from := req.FormValue("from")
	to := req.FormValue("to")
	token := req.FormValue("token")
	amount := req.FormValue("amount")
	db := mysql.ConnectDatabase()
	
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

	err = db.Where("address = ?",to).First(&coin2).Error
	if err != nil{
		http.Error(res,"this wallet address not exists",http.StatusBadRequest)
		return
	}

	if token == "BTC" && coin.BTC >=uint(value) {

	} else if token == "ETH" && coin.ETH >=uint(value) {
		
	}

}