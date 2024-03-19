package db

import (
	"context"
	"ex-go/internal/model/user"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var UserSet = wire.NewSet(wire.Struct(new(UserRepository), "*"), wire.Bind(new(IUserRepository), new(*UserRepository)))

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Save(ctx context.Context, user *user.User) error {
	// 实现保存用户到 MySQL 数据库的逻辑
	return nil
}

func (r *UserRepository) Update(ctx context.Context, user *user.User) error {
	// 实现保存用户到 MySQL 数据库的逻辑
	return nil
}

func (r *UserRepository) FindByUid(ctx context.Context, user *user.User) error {
	// 实现保存用户到 MySQL 数据库的逻辑
	return nil
}

func (r *UserRepository) GetUserList(ctx context.Context, user *user.User) error {
	// 实现保存用户到 MySQL 数据库的逻辑
	return nil
}
