package util

import "golang.org/x/crypto/bcrypt"

func GenerateHashPw(pw string) (string,error){
	hashPw,err := bcrypt.GenerateFromPassword([]byte(pw),bcrypt.DefaultCost)
	if err != nil{
		return "",err
	}
	return string(hashPw),nil
} 

func CompareHashPw(hashPw,pw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPw),[]byte(pw))
	if err != nil{
		return false
	}
	return true
}