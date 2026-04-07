package service

import (
	"errors"
	"yuko_chat/internal/dao"
	"yuko_chat/internal/dto/request"
	"yuko_chat/internal/dto/respond"
	"yuko_chat/internal/model"
	"yuko_chat/pkg/constant"
	group_enum "yuko_chat/pkg/enum/group"
	user_enum "yuko_chat/pkg/enum/user"
	"yuko_chat/pkg/zlog"

	"gorm.io/gorm"
)

var UserContactService = new(userContactService)

type userContactService struct {
}

func (s *userContactService) GetContactList(req request.GetContactListRequest) (string, int, []respond.UserListResponse) {
	var contactList []model.UserContact
	err := dao.DB.Where("user_id = ? and status not in (?, ?)", // 去掉删除过的好友
		req.UserId, model.StatusDeletedFriend, model.StatusBeDeletedFriend).
		Find(&contactList).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg := "目前没有联系人"
			zlog.Info(msg)
			return msg, 0, nil
		}
		zlog.Error(err.Error())
		return constant.SYS_ERR_MSG, -1, nil
	}

	var res []respond.UserListResponse
	for _, contact := range contactList {
		// 首先联系人是用户不是群聊
		if contact.ContactType == model.ContactTypeUser {
			// 获取用户信息
			var user model.UserInfo
			err := dao.DB.Where("uuid = ?", req.UserId).First(&user).Error
			if err != nil {
				zlog.Error(err.Error())
				return constant.SYS_ERR_MSG, -1, nil
			}
			res = append(res, respond.UserListResponse{
				UserId:   user.Uuid,
				UserName: user.Nickname,
				Avatar:   user.Avatar,
			})
		}
	}

	return "获取联系人列表成功", 0, res
}

// LoadMyJoinedGroup 获取我加入的群聊
func (s *userContactService) LoadMyJoinedGroup(req request.GetContactListRequest) (string, int, []respond.LoadMyJoinedGroupRespond) {
	// 先找出用户的 contact 记录
	var contactList []model.UserContact
	err := dao.DB.Where("user_id = ?", req).Find(&contactList).Error
	if err != nil {
		zlog.Error(err.Error())
		return constant.SYS_ERR_MSG, -1, nil
	}

	// 找出所有关联的群聊
	var res []respond.LoadMyJoinedGroupRespond
	for _, contact := range contactList {
		if contact.ContactType == model.ContactTypeGroup &&
			contact.Status == model.StatusNormal { // 只处理未退出和未被踢出的群聊
			var group model.GroupInfo
			err = dao.DB.Model(&model.GroupInfo{}).
				Where("uuid = ?", contact.ContactId).
				First(&group).
				Error
			if err != nil {
				zlog.Error(err.Error())
				return constant.SYS_ERR_MSG, -1, nil
			}
			res = append(res, respond.LoadMyJoinedGroupRespond{
				GroupId:   group.Uuid,
				GroupName: group.Name,
				Avatar:    group.Avatar,
			})
		}
	}

	return "获取群聊成功", 0, res

}

// GetContactInfo 获取联系人信息
func (s *userContactService) GetContactInfo(req request.GetContactInfoRequest) (string, int, *respond.GetContactInfoRespond) {
	if req.ContactId[0] == 'G' { // 是群聊
		var group model.GroupInfo
		err := dao.DB.First(&group, "uuid = ?", req.ContactId).Error
		if err != nil {
			zlog.Error(err.Error())
			return constant.SYS_ERR_MSG, -1, nil
		}
		if group.Status != group_enum.DISABLE {
			return "获取联系人成功", 0, &respond.GetContactInfoRespond{
				ContactId:        group.Uuid,
				ContactName:      group.Name,
				ContactAvatar:    group.Avatar,
				ContactNotice:    group.Notice,
				ContactAddMode:   group.AddMode,
				ContactMembers:   group.Members,
				ContactMemberCnt: group.MemberCnt,
				ContactOwnerId:   group.OwnerId,
			}
		}
		msg := "群聊处于禁用状态"
		zlog.Error(msg)
		return msg, -2, nil
	} else { // 是用户
		var user model.UserInfo
		err := dao.DB.First(&user, "uuid = ?", req.ContactId).Error
		if err != nil {
			zlog.Error(err.Error())
			return constant.SYS_ERR_MSG, -1, nil
		}
		if user.Status != user_enum.DISABLE {
			return "获取联系人成功", 0, &respond.GetContactInfoRespond{
				ContactId:        user.Uuid,
				ContactName:      user.Nickname,
				ContactAvatar:    user.Avatar,
				ContactBirthday:  user.Birthday,
				ContactEmail:     user.Email,
				ContactPhone:     user.Telephone,
				ContactGender:    user.Gender,
				ContactSignature: user.Signature,
			}
		}
		msg := "用户处于禁用状态"
		zlog.Info(msg)
		return msg, -2, &respond.GetContactInfoRespond{}
	}
}
