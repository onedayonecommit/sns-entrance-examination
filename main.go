package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/onedayonecommit/sns/mysql"
	"github.com/onedayonecommit/sns/route"
)



func main(){
	db:= mysql.ConnectDatabase()
	fmt.Println("db connect success")
	close,_:= db.DB()
	defer close.Close()
	router := mux.NewRouter()
	router.HandleFunc("/v1/user/register",route.SignupHandler)
	router.HandleFunc("/v1/user/login",route.LoginHandler)
	router.HandleFunc("/v1/wallets",route.GetAllWalletHandler)
	router.HandleFunc("/v1/wallet/balance/{ADDRESS}",route.GetBalanceHandler)
	router.HandleFunc("/v1/market/convert",route.ExchangeHandler)
	router.HandleFunc("/v1/wallet/transfer",route.TransferHandler)
	
	err:= http.ListenAndServe(":3000",router)
	if err != nil{
		log.Fatalln("server open failed")
	}
}