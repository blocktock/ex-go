package db

import (
	"ex-go/internal/model/user"
	"github.com/google/wire"
)

var DBRepoSet = wire.NewSet(
	UserSet,
) // end

type IUserRepository interface {
	Save(user *user.User) error
	Update(user *user.User) error
	FindByUid(user *user.User) error
	GetUserList(user *user.User) error
}
