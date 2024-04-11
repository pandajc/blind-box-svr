package biz

import (
	"context"
	"errors"
	"fengshen-svr/common/errs"
	"fengshen-svr/dto/params"
	"fengshen-svr/global"
	"fengshen-svr/model"
	"github.com/gin-gonic/gin"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IUserBiz interface {
	Login(ctx context.Context, param *params.UserReq) (*params.UserResp, error)
}

type UserBiz struct {
}

var userBiz IUserBiz

func NewUserBiz() {
	userBiz = &UserBiz{}
}

func GetUserBiz() IUserBiz {
	return userBiz
}

func (s *UserBiz) Login(ctx context.Context, param *params.UserReq) (*params.UserResp, error) {
	var openid string
	if gc, ok := ctx.(*gin.Context); ok {
		openid = gc.GetHeader("X-WX-OPENID")
		for key, values := range gc.Request.Header {
			log.Infof("%v %v", key, values)
		}
	}

	if len(openid) == 0 {
		return nil, errs.NewRespErr(errs.ErrCodeInvalidParam, "no openid")
	}

	userTab, err := s.queryUserByOpenid(openid)
	if err != nil {
		return nil, err
	}
	if userTab != nil {
		if param.GameTime > userTab.GameTime {
			log.Infof("id %v openid %v preparing to update db, param %+v", userTab.ID, userTab.Openid, param)
			err := s.updateUser(userTab.ID, param)
			if err != nil {
				log.Errorf("updateUser err=%v", err)
				return nil, err
			}
			resp := &params.UserResp{
				ID:       userTab.ID,
				GameTime: param.GameTime,
				Money:    param.Money,
				Update:   false,
			}
			return resp, nil
		} else {
			log.Infof("id %v openid %v use db data %+v, param %+v", userTab.ID, userTab.Openid, userTab, param)
			resp := &params.UserResp{
				ID:       userTab.ID,
				GameTime: userTab.GameTime,
				Money:    userTab.Money,
				Update:   true,
			}
			return resp, nil
		}
	}
	log.Infof("openid %v create user", openid)
	newUser, err := s.createUser(openid)
	if err != nil {
		return nil, err
	}
	resp := &params.UserResp{
		ID:       newUser.ID,
		GameTime: newUser.GameTime,
		Money:    newUser.Money,
		Update:   true,
	}
	return resp, nil

}

func (s *UserBiz) updateUser(id int64, param *params.UserReq) error {
	return model.UserTabMgr(global.GetDB()).Where("id = ?", id).UpdateColumns(map[string]interface{}{
		"game_time": param.GameTime, "money": param.Money, "update_time": time.Now()}).Error
}

func (s *UserBiz) createUser(openid string) (*model.UserTab, error) {
	newUser := &model.UserTab{
		Openid:     openid,
		Money:      0,
		GameTime:   0,
		State:      1,
		CreateTime: time.Now(),
	}
	err := model.UserTabMgr(global.GetDB()).Create(newUser).Error
	if err != nil {
		log.Errorf("create err=%v", err)
		return nil, err
	}
	return newUser, nil
}

func (s *UserBiz) queryUserByOpenid(openid string) (*model.UserTab, error) {
	user := &model.UserTab{}
	err := model.UserTabMgr(global.GetDB()).Where("openid= ? and state = ?", openid, 1).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		log.Errorf("queryUserByOpenid call first err=%v", err)
		return nil, err
	}
	return user, nil
}
