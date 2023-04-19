package main

import (
	"fmt"

	"github.com/onedayonecommit/sns/mysql"
)



func main(){
	dbClose:= mysql.ConnectDatabase()
	defer dbClose.Close()
	fmt.Println("db connect success")
}