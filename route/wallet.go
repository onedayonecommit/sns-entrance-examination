package route

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/onedayonecommit/sns/mysql"
)

func GetAllWalletHandler(res http.ResponseWriter, req *http.Request){
	if req.Method != "GET" {
		http.Error(res,"request method is not allowed",http.StatusMethodNotAllowed)
		return
	}
	var w []struct{
		Wallet_address string `json:"wallet_address"` // 한번에 여러개의 데이터를 받아야하므로 []
	}
	
	db :=mysql.ConnectDatabase()
	
	err:= db.Table("wallets").Select("wallet_address").Find(&w).Error // wallets테이블에서 wallet_address 데이터만 모두 출력
	fmt.Println(w)
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