package service

import (
	"errors"
	"fmt"
	"time"
	"yuko_chat/internal/dao"
	"yuko_chat/internal/dto/request"
	"yuko_chat/internal/dto/respond"
	"yuko_chat/internal/model"
	"yuko_chat/pkg/constant"
	"yuko_chat/pkg/util"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type userInfoService struct {
}

var UserInfoService = new(userInfoService)

func (s *userInfoService) Login(telephone, password string) (string, *respond.LoginRespond, int) {
	var userInfo model.UserInfo
	err := dao.DB.Where("telephone = ? AND password = ?", telephone, password).First(&userInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "手机号或密码错误", nil, -2
		}
		return constant.SYS_ERR_MSG, nil, -1
	}

	loginRsp := &respond.LoginRespond{
		Uuid:      userInfo.Uuid,
		Telephone: userInfo.Telephone,
		Nickname:  userInfo.Nickname,
		Email:     userInfo.Email,
		Avatar:    userInfo.Avatar,
		Gender:    userInfo.Gender,
		Birthday:  userInfo.Birthday,
		Signature: userInfo.Signature,
		IsAdmin:   userInfo.IsAdmin,
		Status:    userInfo.Status,
	}
	year, month, day := userInfo.CreatedAt.Date()
	loginRsp.CreatedAt = fmt.Sprintf("%d.%d.%d", year, month, day)
	return "登录成功", loginRsp, 0
}

func (s *userInfoService) Register(req request.RegisterRequest) (string, *respond.RegisterRespond, int) {
	redisKey := fmt.Sprintf("verify:%s", req.Email)
	storedCode, err := dao.GetKeyNilIsErr(redisKey)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "验证码不存在或已过期", nil, -2
		}
		return "系统错误", nil, -1
	}
	if storedCode != req.VerificationCode {
		return "验证码错误", nil, -1
	}

	userInfo := model.UserInfo{
		Uuid:      util.GenUserUUID(),
		Telephone: req.Telephone,
		Password:  req.Password,
		Nickname:  req.Nickname,
		Email:     req.Email,
		CreatedAt: time.Now(),
	}

	err = dao.DB.Create(&userInfo).Error
	if err != nil {
		return "创建用户失败", nil, -2
	}

	registerRsp := &respond.RegisterRespond{
		Uuid:      userInfo.Uuid,
		Telephone: userInfo.Telephone,
		Nickname:  userInfo.Nickname,
		Email:     userInfo.Email,
		Avatar:    userInfo.Avatar,
		Gender:    userInfo.Gender,
		Birthday:  userInfo.Birthday,
		Signature: userInfo.Signature,
		IsAdmin:   userInfo.IsAdmin,
		Status:    userInfo.Status,
	}
	year, month, day := userInfo.CreatedAt.Date()
	registerRsp.CreatedAt = fmt.Sprintf("%d.%d.%d", year, month, day)

	return "注册成功", registerRsp, 0
}

func (s *userInfoService) SendVerificationCode(email string) (string, int) {
	code, err := util.SendEmail(email)
	if err != nil {
		return "发送邮箱验证码失败", -1
	}

	redisKey := fmt.Sprintf("verify:%s", email)
	if err = dao.SetKeyEx(redisKey, code, 5*time.Minute); err != nil {
		return "缓存邮箱验证码失败", -1
	}

	return "发送邮箱验证码成功", 0
}

func (s *userInfoService) GetUserInfo(uuid string) (string, *respond.GetUserInfoRespond, int) {
	var userInfo model.UserInfo
	err := dao.DB.Where("uuid = ?", uuid).First(&userInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "用户不存在", nil, -2
		}
		return constant.SYS_ERR_MSG, nil, -1
	}

	getUserInfoRsp := &respond.GetUserInfoRespond{
		Uuid:      userInfo.Uuid,
		Telephone: userInfo.Telephone,
		Nickname:  userInfo.Nickname,
		Email:     userInfo.Email,
		Avatar:    userInfo.Avatar,
		Gender:    userInfo.Gender,
		Birthday:  userInfo.Birthday,
		Signature: userInfo.Signature,
		IsAdmin:   userInfo.IsAdmin,
		Status:    userInfo.Status,
	}
	year, month, day := userInfo.CreatedAt.Date()
	getUserInfoRsp.CreatedAt = fmt.Sprintf("%d.%d.%d", year, month, day)
	return "获取用户信息成功", getUserInfoRsp, 0
}

func (s *userInfoService) UpdateUserInfo(req request.UpdateUserInfoRequest) (string, int) {
	updateData := map[string]any{}
	fields := map[string]string{
		"email":     req.Email,
		"nickname":  req.Nickname,
		"birthday":  req.Birthday,
		"signature": req.Signature,
		"avatar":    req.Avatar,
	}

	for key, value := range fields {
		if value != "" {
			updateData[key] = value
		}
	}

	if len(updateData) == 0 {
		return "没有需要更新的信息", -2
	}

	result := dao.DB.Model(&model.UserInfo{}).Where("uuid = ?", req.Uuid).Updates(updateData)
	if result.Error != nil {
		return "更新用户信息失败", -1
	}
	if result.RowsAffected == 0 {
		return "用户不存在", -2
	}

	return "更新用户信息成功", 0
}
