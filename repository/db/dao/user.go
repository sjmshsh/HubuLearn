package dao

import (
	"HubuLearn/repository/db/model"
	"context"
	"github.com/jinzhu/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{}
}

// FindUserByUserName 根据用户名找到用户
func (dao *UserDao) FindUserByUserName(userName string) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).
		First(&user).Error

	return
}

// FindUserByUserId 根据用户id找到用户
func (dao *UserDao) FindUserByUserId(id string) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", id).
		First(&user).Error

	return
}

// CreateUser 创建User
func (dao *UserDao) CreateUser(user *model.User) (err error) {
	err = dao.DB.Model(&model.User{}).Create(user).Error

	return
}
