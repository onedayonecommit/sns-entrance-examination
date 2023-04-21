package route

import (
	"fmt"
	"net/http"
	"time"

	"github.com/onedayonecommit/sns/util"
)

func LoginHandler(res http.ResponseWriter,req *http.Request){
	if req.Method != "POST" {
		http.Error(res,"request method is not allowed",http.StatusMethodNotAllowed)
		return
	}

	err := req.ParseForm()
	if err != nil{
		http.Error(res,"content type is wrong",http.StatusBadRequest)
		return
	}
	email:=req.FormValue("email")
	password := req.FormValue("password")
	err,hashPw := UserCheck(*&email)
	if err == nil {
		result := util.CompareHashPw(hashPw,*&password)
		if result {
			// 성공시 Jwt 토큰 발급
			token,err := util.GenerateJwt(UserWallet(email))
			if err !=nil {
				http.Error(res,"Generate token failed",http.StatusInternalServerError)
			}
			cookie := &http.Cookie{
				Name: "koa:sess",
				Value: token,
				Expires: time.Now().Add(time.Minute*15),
				HttpOnly: true,
				Secure: false,
				SameSite: http.SameSiteLaxMode,
				Path: "/",
			}
			http.SetCookie(res, cookie)
			fmt.Fprintln(res,"login successful",token)
			return
		}
		http.Error(res,"Password do not match",http.StatusBadRequest)
		return
	}
}

func LogOutHandler(res http.ResponseWriter, req *http.Request){
	if req.Method != "POST" {
		http.Error(res,"request method is not allowed",http.StatusMethodNotAllowed)
		return
	}
	cookie:= &http.Cookie{
		Name: "koa:sess",
		Value: "",
		Expires: time.Unix(0,0),
		HttpOnly: true,
		Secure: false,
		SameSite: http.SameSiteLaxMode,
		Path: "/",
	}
	http.SetCookie(res, cookie)
	fmt.Fprintln(res,"logout")
	return
}