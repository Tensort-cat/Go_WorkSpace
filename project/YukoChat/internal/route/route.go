package route

import (
	"yuko_chat/internal/config"
	"yuko_chat/internal/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var GE *gin.Engine

func init() {
	GE = gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Content-Type", "Authorization"}
	GE.Use(cors.New(corsConfig))
	GE.Static("/static/avatars", config.Cfg.StaticSrcConfig.StaticAvatarPath)
	GE.Static("/static/files", config.Cfg.StaticSrcConfig.StaticFilePath)

	GE.POST("/login", controller.Login)
	GE.POST("/register", controller.Register)

	user := GE.Group("/user")
	{
		user.POST("/updateUserInfo", controller.UpdateUserInfo)
		user.POST("/getUserInfo", controller.GetUserInfo)
		user.POST("/sendSmsCode", controller.SendVerificationCode)
	}

	group := GE.Group("/group")
	{
		group.POST("/createGroup", controller.CreateGroup)
		group.POST("/loadMyGroup", controller.GetMyGroups)
		group.POST("/checkGroupAddMode", controller.CheckGroupAddMode)
		group.POST("/enterGroupDirectly", controller.EnterGroupDirectly)
		group.POST("/leaveGroup", controller.LeaveGroup)
		group.POST("/dismissGroup", controller.DismissGroup)
		group.POST("/getGroupInfo", controller.GetGroupInfo)
		group.POST("/updateGroupInfo", controller.UpdateGroupInfo)
		group.POST("/removeGroupMembers", controller.RemoveGroupMembers)
		group.POST("/getGroupInfoList", controller.GetGroupMembers)
	}
}
