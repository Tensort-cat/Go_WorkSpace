package service

import (
	"yuko_chat/internal/dao"
	"yuko_chat/internal/dto/request"
	"yuko_chat/internal/dto/respond"
	"yuko_chat/internal/model"
	"yuko_chat/pkg/constant"
	"yuko_chat/pkg/zlog"
)

type messageService struct {
}

var MessageService = new(messageService)

// GetMessageList 获取聊天记录
func (m *messageService) GetMessageList(req request.GetMessageListRequest) (string, int, []respond.GetMessageListRespond) {
	var messages []model.Message
	err := dao.DB.Where(
		"(send_id = ? and receive_id = ?) or (send_id = ? and receive_id = ?)", req.UserOneId, req.UserTwoId, req.UserTwoId, req.UserOneId).
		Find(&messages).
		Error
	if err != nil {
		zlog.Error(err.Error())
		return constant.SYS_ERR_MSG, -1, nil
	}

	var res []respond.GetMessageListRespond
	for _, msg := range messages {
		res = append(res, respond.GetMessageListRespond{
			SendId:     msg.SendId,
			SendName:   msg.SendName,
			SendAvatar: msg.SendAvatar,
			ReceiveId:  msg.ReceiveId,
			Type:       msg.Type,
			Content:    msg.Content,
			Url:        msg.Url,
			FileType:   msg.FileType,
			FileName:   msg.FileName,
			FileSize:   msg.FileSize,
			CreatedAt:  msg.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return "获取聊天记录成功", 0, res
}

// GetGroupMessageList 获取群聊消息记录
func (m *messageService) GetGroupMessageList(req request.GetGroupMessageListRequest) (string, int, []respond.GetGroupMessageListRespond) {
	var messages []model.Message
	err := dao.DB.Where("receive_id = ?", req.GroupId).Order("created_at asc").Find(&messages).Error
	if err != nil {
		zlog.Error(err.Error())
		return constant.SYS_ERR_MSG, -1, nil
	}

	var res []respond.GetGroupMessageListRespond
	for _, msg := range messages {
		res = append(res, respond.GetGroupMessageListRespond{
			SendId:     msg.SendId,
			SendName:   msg.SendName,
			SendAvatar: msg.SendAvatar,
			ReceiveId:  msg.ReceiveId,
			Type:       msg.Type,
			Content:    msg.Content,
			Url:        msg.Url,
			FileType:   msg.FileType,
			FileSize:   msg.FileSize,
			FileName:   msg.FileName,
			CreatedAt:  msg.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return "获取群聊聊天记录成功", 0, res
}

// todo: 文件上传与下载
