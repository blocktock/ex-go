package db

import (
	"ex-go/internal/model/user"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var UserSet = wire.NewSet(wire.Struct(new(UserRepository), "*"), wire.Bind(new(IUserRepository), new(*UserRepository)))

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Save(user *user.User) error {
	// 实现保存用户到 MySQL 数据库的逻辑
	return nil
}

func (r *UserRepository) Update(user *user.User) error {
	// 实现保存用户到 MySQL 数据库的逻辑
	return nil
}

func (r *UserRepository) FindByUid(user *user.User) error {
	// 实现保存用户到 MySQL 数据库的逻辑
	return nil
}

func (r *UserRepository) GetUserList(user *user.User) error {
	// 实现保存用户到 MySQL 数据库的逻辑
	return nil
}
