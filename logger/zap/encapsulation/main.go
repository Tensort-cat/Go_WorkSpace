package main

import (
	"zap_demo/encapsulation/utils"

	"go.uber.org/zap"
)

type User struct {
	Name string
}

func main() {
	user := &User{Name: "Kevin"}
	utils.Info("test log", zap.Any("user", user))
	utils.Debug("test log", zap.Any("user", user))
	utils.Warn("test log", zap.Any("user", user))
	utils.Error("test log", zap.Any("user", user))
}
