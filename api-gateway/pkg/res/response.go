package res

import (
	"api-gateway/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 基础序列化器
type Response struct {
	Status uint        `json:"Status"`
	Data   interface{} `json:"Data"`
	Msg    string      `json:"Msg"`
	Error  string      `json:"Error"`
}

// DataList 带有总数的Data结构
type DataList struct {
	Item  interface{} `json:"Item"`
	Total uint        `json:"Total"`
}

// TokenData 带有token的Data结构
type TokenData struct {
	User  interface{} `json:"User"`
	Token string      `json:"Token"`
}

// Ok 返回200 自定义code data
func Ok(ctx *gin.Context, msgCode int, data interface{}) {
	ctx.JSON(http.StatusOK, ginH(msgCode, data))
}

// Unauthorized 无权限err
func Unauthorized(ctx *gin.Context, msgCode int) {
	ctx.JSON(http.StatusUnauthorized, ginH(msgCode, nil))
}

// InternalError 内部err
func InternalError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, ginH(e.Error, nil))
}

// ForbiddenError 禁止访问err
func ForbiddenError(ctx *gin.Context, msgCode int) {
	ctx.JSON(http.StatusForbidden, ginH(msgCode, nil))
}

// Error 自定义 err
func Error(ctx *gin.Context, httpCode, msgCode int) {
	ctx.JSON(httpCode, ginH(msgCode, nil))
}

func ginH(msgCode int, data interface{}) gin.H {
	return gin.H{
		"code": msgCode,
		"msg":  e.GetMsg(uint(msgCode)),
		"data": data,
	}
}
