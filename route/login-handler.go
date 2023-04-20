package route

import "net/http"

type loginBody struct{
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func LoginHandler(res http.ResponseWriter,req *http.Request){
	if req.Method != "POST" {
		http.Error(res,"request method is wrong",http.StatusBadRequest)
		return
	}

	var l loginBody
}