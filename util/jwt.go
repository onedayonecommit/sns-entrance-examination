package util

import (
	"errors"
	"fmt"
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

func VerifyJwt(tokenValue string) (string,error){
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Env Load Failed")
	}

	if tokenValue == ""{
		return "", errors.New("jwt is not found")
	}

	var secretKey = []byte(os.Getenv("SECRET_KEY"))

	parsedToken,err := jwt.Parse(tokenValue,func(t *jwt.Token) (interface{}, error) {
		if _,ok := t.Method.(*jwt.SigningMethodHMAC); !ok{ // gpt
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"]) // gpt
		}
		return secretKey,nil
	})
	if err != nil {
		return "",err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid { // gpt
		address,ok:=claims["address"].(string)
		if !ok {
			return "",errors.New("jwt parse failed")
		}
		return address,nil
	} 
	return "",nil

}
