package repository

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"user/internal/service"
	"user/pkg/e"
)

type User struct {
	UserID         uint64 `gorm:"primaryKey;not null" json:"user_id"`
	UserName       string `json:"user_name"`
	NickName       string `json:"nick_name"`
	PasswordDigest string `json:"password_digest"`
}

const (
	PasswordCost = 12
)

// CheckUserExist 检查用户是否存在
func (u *User) CheckUserExist(req *service.UserRequest) bool {
	if err := DB.Where("user_name = ?", req.UserName).First(&u).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

// ShowUserInfo 获取用户信息
func (u *User) ShowUserInfo(req *service.UserRequest) error {
	if exist := u.CheckUserExist(req); exist {
		return nil
	}
	return errors.New(e.GetMsg(e.UserNotExist))
}

// UserCreate 创建用户
func (u *User) UserCreate(req *service.UserRequest) error {
	if exist := u.CheckUserExist(req); exist {
		return errors.New(e.GetMsg(e.UserAlreadyExist))
	}

	user := User{
		UserName: req.UserName,
		NickName: req.NickName,
	}

	err := user.SetPasswordDigest(req.Password)
	if err != nil {
		return errors.New(e.GetMsg(e.SetPasswordDigestError) + err.Error())
	}

	err = DB.Create(&user).Error
	if err != nil {
		return errors.New(e.GetMsg(e.UserCreateError) + err.Error())
	}

	return nil
}

// SetPasswordDigest 加密密码
func (u *User) SetPasswordDigest(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(bytes)
	return nil
}

// SerializeUser 序列化用户信息
func SerializeUser(item User) *service.UserModel {
	return &service.UserModel{
		UserID:   item.UserID,
		UserName: item.UserName,
		NickName: item.NickName,
	}
}
