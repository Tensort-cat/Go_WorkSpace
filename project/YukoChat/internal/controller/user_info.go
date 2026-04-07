package controller

import (
	"net/http"
	"yuko_chat/internal/dto/request"
	"yuko_chat/internal/service"
	"yuko_chat/pkg/constant"
	"yuko_chat/pkg/zlog"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var req request.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		zlog.Error(err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"code": constant.SYS_ERR_CODE,
			"msg":  constant.SYS_ERR_MSG,
		})
		return
	}
	msg, userInfo, ret := service.UserInfoService.Login(req.Telephone, req.Password)
	JsonBack(ctx, msg, ret, userInfo)
}

func Register(ctx *gin.Context) {
	var req request.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		zlog.Error(err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"code": constant.SYS_ERR_CODE,
			"msg":  constant.SYS_ERR_MSG,
		})
		return
	}
	msg, userInfo, ret := service.UserInfoService.Register(req)
	JsonBack(ctx, msg, ret, userInfo)
}

func SendVerificationCode(ctx *gin.Context) {
	var req request.SendVerificationCodeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		zlog.Error(err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"code": constant.SYS_ERR_CODE,
			"msg":  constant.SYS_ERR_MSG,
		})
		return
	}
	msg, ret := service.UserInfoService.SendVerificationCode(req.Email)
	JsonBack(ctx, msg, ret, nil)
}

func GetUserInfo(ctx *gin.Context) {
	var req request.GetUserInfoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		zlog.Error(err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"code": constant.SYS_ERR_CODE,
			"msg":  constant.SYS_ERR_MSG,
		})
		return
	}
	msg, userInfo, ret := service.UserInfoService.GetUserInfo(req.Uuid)
	JsonBack(ctx, msg, ret, userInfo)
}

func UpdateUserInfo(ctx *gin.Context) {
	var req request.UpdateUserInfoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		zlog.Error(err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"code": constant.SYS_ERR_CODE,
			"msg":  constant.SYS_ERR_MSG,
		})
		return
	}
	msg, ret := service.UserInfoService.UpdateUserInfo(req)
	JsonBack(ctx, msg, ret, nil)
}
