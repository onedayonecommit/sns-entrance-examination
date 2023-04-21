package route

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onedayonecommit/sns/mysql"
	"github.com/onedayonecommit/sns/mysql/model"
	"github.com/onedayonecommit/sns/util"
)

type CoinResult struct{
	Coin string
	Balance uint
}

type Response struct{
	Address string
	Coins []CoinResult
}

type Wallets struct{
	wallets []Response
}

func GetAllWalletHandler(res http.ResponseWriter, req *http.Request){
	if req.Method != "GET" {
		http.Error(res,"request method is not allowed",http.StatusMethodNotAllowed)
		return
	}
	// var w []struct{
	// 	Wallet_address []string `json:"wallets"` // 한번에 여러개의 데이터를 받아야하므로 []
	// }
	
	db :=mysql.ConnectDatabase()
	
	var coin []model.Coin
	// err:= db.Table("wallets").Select("wallet_address").Find(&w).Error // wallets테이블에서 wallet_address 데이터만 모두 출력
	err:= db.Select("address,eth,btc").Find(&coin) // wallets테이블에서 wallet_address 데이터만 모두 출력
	if err != nil {
		res.WriteHeader(http.StatusOK)
		fmt.Fprintln(res,"There is no list of verified wallets.") // 지갑목록이 없으면
		return 
	}

	walletMap:= make(map[string]Response)
	for _,value := range coin{
		address,torF := walletMap[value.Address]
		if torF {
			address.Coins = append(address.Coins,CoinResult{"eth",value.ETH})
			address.Coins = append(address.Coins,CoinResult{"btc",value.BTC})
			walletMap[value.Address] = address
		}
	}

	wallet:= make([]Response,0,len(walletMap))
	for _,value := range walletMap{
		wallet = append(wallet, value)
	}

	Wallets := Wallets{wallets: wallet}

	res.Header().Set("Content-Type","application/json")
	json.NewEncoder(res).Encode(Wallets)

	close,_ := db.DB()
	defer close.Close()

	return
}

func GetBalanceHandler(res http.ResponseWriter,req *http.Request){
	if req.Method != "GET" {
		http.Error(res,"request method is not allowed",http.StatusMethodNotAllowed)
		return
	}
	var coin model.Coin

	vars := mux.Vars(req)
	address:= vars["ADDRESS"]

	db:= mysql.ConnectDatabase()
	close,_ := db.DB()
	defer close.Close()

	err := db.Select("BTC, ETH").Where("address = ? ",address).First(&coin).Error
	fmt.Println(err)
	if err != nil {
		http.Error(res,"this wallet info is not found",http.StatusNotFound)
		return
	}

	resResult := Response{
		Address: address,
		Coins: []CoinResult{
			{Coin: "BTC", Balance: coin.BTC},
			{Coin: "ETH", Balance: coin.ETH},
		},
	}
	res.Header().Set("Content-Type","application/json")
	json.NewEncoder(res).Encode(resResult)
	return
}

func MyWalletBalance(res http.ResponseWriter, req *http.Request){
	if req.Method != "GET"{
		http.Error(res,"request method is not allowed",http.StatusMethodNotAllowed)
		return
	}

	var coin model.Coin
	cookie,err := req.Cookie("koa:sess")
	if err!= nil {
		http.Error(res,"cookie is not found",http.StatusBadRequest)
		return
	}
	tokenValue:= cookie.Value
	address,err:=util.VerifyJwt(tokenValue)

	if err != nil && err.Error() == "jwt is not found" {
		http.Error(res,"jwt is already exp",http.StatusBadRequest)
		return
	} else if err !=nil {
		http.Error(res,"unknown err",http.StatusInternalServerError)
		return
	}

	db := mysql.ConnectDatabase()
	close,_ := db.DB()
	defer close.Close()

	err = db.Select("BTC, ETH").Where("address = ? ",address).First(&coin).Error
	if err != nil {
		http.Error(res,"this wallet info is not found",http.StatusNotFound)
		return
	}

	resResult := Response{
		Address: address,
		Coins: []CoinResult{
			{Coin: "BTC", Balance: coin.BTC},
			{Coin: "ETH", Balance: coin.ETH},
		},
	}
	res.Header().Set("Content-Type","application/json")
	json.NewEncoder(res).Encode(resResult)
	return
}