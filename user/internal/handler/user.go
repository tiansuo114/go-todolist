package handler

import (
	"context"
	"user/internal/repository"
	"user/internal/service"
	"user/pkg/e"
)

type UserService struct {
	service.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) UserLogin(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user repository.User
	resp = new(service.UserDetailResponse)
	resp.Code = e.Success
	err = user.ShowUserInfo(req)
	if err != nil {
		resp.Code = e.Error
		resp.Message = e.GetMsg(uint(resp.Code))
		return resp, err
	}

	resp.UserDetail = repository.SerializeUser(user)
	return resp, nil
}

func (u *UserService) UserRegister(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user repository.User
	resp = new(service.UserDetailResponse)
	resp.Code = e.Success
	resp.Message = e.GetMsg(uint(resp.Code))

	err = user.UserCreate(req)
	if err != nil {
		resp.Code = e.Error
		resp.Message = e.GetMsg(uint(resp.Code))
		return resp, err
	}

	resp.UserDetail = repository.SerializeUser(user)
	return resp, nil
}
