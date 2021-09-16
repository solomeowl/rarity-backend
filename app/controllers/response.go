package controllers

import (
	"github.com/gin-gonic/gin"
)

type Reply struct {
	HttpStatus int         `json:"http_status"`
	ErrorCode  int         `json:"error_code"`
	Msg        string      `json:"msg"`
	Total      int         `json:"total"`
	Data       interface{} `json:"data"`
}
type Token struct {
	AccessToken string `json:"access_token"`
}

func Response(g *gin.Context, httpCode, errorCode, total int, msg string, data interface{}) {
	g.JSON(httpCode, Reply{
		HttpStatus: httpCode,
		ErrorCode:  errorCode,
		Msg:        msg,
		Total:      total,
		Data:       data,
	})
}
