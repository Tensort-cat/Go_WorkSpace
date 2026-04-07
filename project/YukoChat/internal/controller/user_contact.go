package controller

import (
	"net/http"
	"yuko_chat/internal/dto/request"
	"yuko_chat/internal/service"
	"yuko_chat/pkg/constant"
	"yuko_chat/pkg/zlog"

	"github.com/gin-gonic/gin"
)

// 获取联系人列表
func GetContactList(ctx *gin.Context) {
	var req request.GetContactListRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		zlog.Error(err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"code": constant.SYS_ERR_CODE,
			"msg":  constant.SYS_ERR_MSG,
		})
		return
	}
	msg, ret, contactList := service.UserContactService.GetContactList(req)
	JsonBack(ctx, msg, ret, contactList)
}

// LoadMyJoinedGroup 获取我加入的群聊
func LoadMyJoinedGroup(ctx *gin.Context) {
	var req request.GetContactListRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		zlog.Error(err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"code": constant.SYS_ERR_CODE,
			"msg":  constant.SYS_ERR_MSG,
		})
		return
	}
	msg, ret, groupList := service.UserContactService.LoadMyJoinedGroup(req)
	JsonBack(ctx, msg, ret, groupList)
}

// GetContactInfo 获取联系人信息
func GetContactInfo(ctx *gin.Context) {
	var req request.GetContactInfoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		zlog.Error(err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"code": constant.SYS_ERR_CODE,
			"msg":  constant.SYS_ERR_MSG,
		})
		return
	}
	msg, ret, contactInfo := service.UserContactService.GetContactInfo(req)
	JsonBack(ctx, msg, ret, contactInfo)
}
