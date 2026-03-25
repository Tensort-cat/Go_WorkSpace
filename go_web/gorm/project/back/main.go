package main

import (
	"fmt"
	"go_web/gorm/project/back/dao"
	"go_web/gorm/project/back/models"
	"go_web/gorm/project/back/routes"
)

func main() {
	err := dao.InitMysql()
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	defer dao.Close()

	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	// 注册路由
	r := routes.SetRouter()
	r.Run()
}
