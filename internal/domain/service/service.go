// internal/schema/service/user_service.go

package service

import (
	"context"
	"ex-go/internal/schema"
	"github.com/google/wire"
)

var ServiceSet = wire.NewSet(
	UserSet,
	LoginSet,
) // end

type IUserService interface {
	CreateUser(ctx context.Context, createUserRequestDTO schema.CreateUserParams) error
}

type ILoginService interface {
	Login(ctx context.Context) error
}
