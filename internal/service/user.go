package service

import (
	"jk-demo-layout/api/demo"
	"jk-demo-layout/internal/biz"
	"log"
)

type UserService struct {
	bz  *biz.UserBiz
	log *log.Logger
}

func NewUserService(bz *biz.UserBiz, log *log.Logger) *UserService {
	return &UserService{bz: bz, log: log}
}

func (u *UserService) SayHello(id string) *demo.HelloDTO {
	userDO := u.bz.Get(id)
	return &demo.HelloDTO{
		Name:     userDO.Name,
		UserName: userDO.UserName,
	}
}
