package service

import (
	"context"
	"ex-go/internal/model/user"
	"ex-go/internal/repository/cache"
	"ex-go/internal/repository/db"
	"ex-go/internal/schema"
	"github.com/google/wire"
)

var UserSet = wire.NewSet(wire.Struct(new(UserService), "*"), wire.Bind(new(IUserService), new(*UserService)))

//var _ IUserService = (*UserService)(nil)

type UserService struct {
	UserRepo      db.IUserRepository
	UserCacheRepo cache.IUserCacheRepository
}

func (s *UserService) CreateUser(ctx context.Context, userParams schema.CreateUserParams) error {

	data := user.User{
		Id:          1,
		EnShortName: userParams.Name,
	}

	return s.UserRepo.Save(ctx, &data)
}
