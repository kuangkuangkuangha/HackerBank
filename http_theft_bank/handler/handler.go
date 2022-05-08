package handler

import (
	"http-theft-bank/log"
	"http-theft-bank/util"
	"net/http"

	"http-theft-bank/pkg/errno"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SetResponseHeader ... 设置响应头
func SetResponseHeader(c *gin.Context, key, value string) {
	c.Header(key, value)
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	log.Info(message,
		zap.String("X-Request-Id", util.GetReqID(c)))

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func SendBadRequest(c *gin.Context, err error, data interface{}, cause string) {
	code, message := errno.DecodeErr(err)
	log.Error(message,
		zap.String("X-Request-Id", util.GetReqID(c)),
		zap.String("cause", cause))

	c.JSON(http.StatusBadRequest, Response{
		Code:    code,
		Message: message + ": " + cause,
		Data:    data,
	})
}

func SendError(c *gin.Context, err error, data interface{}, cause string) {
	code, message := errno.DecodeErr(err)
	log.Error(message,
		zap.String("X-Request-Id", util.GetReqID(c)),
		zap.String("cause", cause))

	c.JSON(http.StatusInternalServerError, Response{
		Code:    code,
		Message: message + ": " + cause,
		Data:    data,
	})
}
