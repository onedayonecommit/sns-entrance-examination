package route

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/onedayonecommit/sns/mysql"
	"github.com/onedayonecommit/sns/mysql/model"
)

func GetAllWalletHandler(res http.ResponseWriter, req *http.Request){
	if req.Method != "GET" {
		http.Error(res,"request method is not allowed",http.StatusMethodNotAllowed)
		return
	}
	var w model.Wallet
	
	db :=mysql.ConnectDatabase()
	
	err:= db.Where(&w)
	if err != nil {
		res.WriteHeader(http.StatusOK)
		fmt.Fprintln(res,"There is no list of verified wallets.") // 지갑목록이 없으면
		return 
	}
	res.Header().Set("Content-Type","application/json")
	json.NewEncoder(res).Encode(w)

	close,_ := db.DB()
	defer close.Close()

	return
}