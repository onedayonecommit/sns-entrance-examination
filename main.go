package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/onedayonecommit/sns/mysql"
)



func main(){
	db:= mysql.ConnectDatabase()
	fmt.Println("db connect success")
	defer db.Close()

	err:= http.ListenAndServe(":3000",nil)
	if err != nil{
		log.Fatalln("server open failed")
	}
}