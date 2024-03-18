package service

import (
	"ex-go/internal/repository/cache"
	"ex-go/internal/repository/db"
	"github.com/google/wire"
)

var LoginSet = wire.NewSet(wire.Struct(new(LoginService), "*"), wire.Bind(new(ILoginService), new(*LoginService)))

type LoginService struct {
	UserRepo      db.IUserRepository
	UserCacheRepo cache.IUserCacheRepository
}

func (s *LoginService) Login() error {
	return nil
}
