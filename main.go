package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/onedayonecommit/sns/mysql"
	"github.com/onedayonecommit/sns/route"
)



func main(){
	db:= mysql.ConnectDatabase()
	fmt.Println("db connect success")
	close,_:= db.DB()
	defer close.Close()
	http.HandleFunc("/v1/test",route.SignupHandler)
	
	err:= http.ListenAndServe(":3000",nil)
	if err != nil{
		log.Fatalln("server open failed")
	}
}