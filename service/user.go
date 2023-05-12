package service

import (
	"HubuLearn/pkg/ctl"
	"HubuLearn/repository/db/dao"
	"HubuLearn/repository/db/model"
	"HubuLearn/types"
	"context"
	"errors"
	"github.com/jinzhu/gorm"
)

type UserSrv struct{}

func GetUserSrv() *UserSrv {
	return &UserSrv{}
}

func (s *UserSrv) Register(ctx context.Context, req *types.UserServiceReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	u, err := userDao.FindUserByUserId(req.UserName)
	switch err {
	case gorm.ErrRecordNotFound:
		// 代表数据库里面没有存放用户的相关信息
		u = &model.User{
			UserName: req.UserName,
		}
		if err = u.SetPassword(req.Password); err != nil {
			return
		}
		if err = userDao.CreateUser(u); err != nil {
			return
		}
		return ctl.RespSuccess(200), nil
	case nil:
		err = errors.New("用户已经存在了")
		return
	default:
		return
	}
}
