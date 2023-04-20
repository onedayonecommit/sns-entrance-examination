package util

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func GenerateJwt(address string) (string ,error){
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Env Load Failed")
	}
	
	secretKey := []byte(os.Getenv("SECRTE_KEY"))
	
	tokenInfo  := jwt.MapClaims{
		"address":address,
		"exp":time.Now().Add(time.Minute * 15).Unix(), // Unix() == 초 단위로 변경 기준은 node.js랑 동일
		"issuer" : "GyeongHwan",
	}
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256,tokenInfo)
	accessToken,err := token.SignedString([]byte(secretKey))
	if err != nil{
		return "",err
	}
	return accessToken,nil
}