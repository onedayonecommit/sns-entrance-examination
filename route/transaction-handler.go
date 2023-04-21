package route

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/onedayonecommit/sns/mysql"
)

func TxHandler(res http.ResponseWriter,req *http.Request){
	if req.Method != "GET"{
		http.Error(res,"request method is not allowed",http.StatusMethodNotAllowed)
		return
	}
	
	vars := mux.Vars(req)
	txId := vars["TRANSACTION_ID"]
	
	db:= mysql.ConnectDatabase()
	close,_ := db.DB()
	defer close.Close()
	var tx []struct{
		From string `json:"from"`
		To string	`json:"to"`
		Amount uint `json:"amount"`
		Coin string `json:"coin"`
	} 
	
	err := db.Table("transactions").Select("from, to, amount, coin").Where("id = ?",txId).Find(&tx).Error
	if err !=nil {
		http.Error(res,"this tx id not found info",http.StatusBadRequest)
		return
	}
	res.Header().Set("Content-Type","application/json")
	json.NewEncoder(res).Encode(tx)

	return
}