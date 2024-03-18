// internal/schema/service/user_service.go

package service

import (
	"ex-go/internal/schema"
	"github.com/google/wire"
)

var ServiceSet = wire.NewSet(
	UserSet,
	LoginSet,
) // end

type IUserService interface {
	CreateUser(createUserRequestDTO schema.CreateUserRequestDTO) error
}

type ILoginService interface {
	Login() error
}
