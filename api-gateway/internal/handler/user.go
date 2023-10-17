package handler

import (
	"api-gateway/internal/service"
	"api-gateway/pkg/e"
	"api-gateway/pkg/res"
	"api-gateway/pkg/util"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(ginCtx *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserErr(ginCtx.Bind(userReq))

	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["user"].(service.UserServiceClient)

	UserResp, err := userService.UserRegister(context.Background(), &userReq)
	PanicIfUserErr(err)

	r := res.Response{
		Status: uint(UserResp.Code),
		Data:   UserResp,
		Msg:    e.GetMsg(uint(UserResp.Code)),
		Error:  err.Error(),
	}

	ginCtx.JSON(http.StatusOK, r)
}

func UserLogin(ginCtx *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserErr(ginCtx.Bind(userReq))

	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["user"].(service.UserServiceClient)

	UserResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicIfUserErr(err)

	token, err := util.GenerateToken(UserResp.UserDetail)

	r := res.Response{
		Status: uint(UserResp.Code),
		Data: res.TokenData{
			User:  UserResp.UserDetail,
			Token: token,
		},
		Msg:   e.GetMsg(uint(UserResp.Code)),
		Error: err.Error(),
	}

	ginCtx.JSON(http.StatusOK, r)
}
