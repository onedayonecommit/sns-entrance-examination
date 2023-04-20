package route

import (
	"net/http"

	"github.com/onedayonecommit/sns/util"
)

type loginBody struct{
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Login handler
func LoginHandler(res http.ResponseWriter,req *http.Request){
	if req.Method != "POST" {
		http.Error(res,"request method is wrong",http.StatusBadRequest)
		return
	}

	var l loginBody

	err := reqBodyCheck(req,&l)

	if err != nil {
		http.Error(res,"request body is wrong",http.StatusBadRequest)
		return
	}

	checkStatus,hashPw := UserCheck(l.Email)
	if !checkStatus {
		result := util.CompareHashPw(hashPw,l.Password)
		if result {
			// 성공시 Jwt 토큰 발급
			res.WriteHeader(http.StatusOK)
			return
		}
		http.Error(res,"Password do not match",http.StatusBadRequest)
		return
	}
}