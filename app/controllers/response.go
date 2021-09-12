package controllers

import (
	"github.com/gin-gonic/gin"
)

type Reply struct {
	HttpStatus int         `json:"http_status"`
	ErrorCode  int         `json:"error_code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}
type Token struct {
	AccessToken string `json:"access_token"`
}

func Response(g *gin.Context, httpCode, errorCode int, msg string, data interface{}) {
	g.JSON(httpCode, Reply{
		HttpStatus: httpCode,
		ErrorCode:  errorCode,
		Msg:        msg,
		Data:       data,
	})
}
