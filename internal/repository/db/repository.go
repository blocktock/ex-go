package db

import (
	"context"
	"ex-go/internal/model/user"
	"github.com/google/wire"
)

var DBRepoSet = wire.NewSet(
	UserSet,
) // end

type IUserRepository interface {
	Save(ctx context.Context, user *user.User) error
	Update(ctx context.Context, user *user.User) error
	FindByUid(ctx context.Context, user *user.User) error
	GetUserList(ctx context.Context, user *user.User) error
}
