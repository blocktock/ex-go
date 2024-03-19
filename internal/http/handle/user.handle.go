package handle

import (
	"ex-go/internal/domain/service"
	"ex-go/internal/schema"
	"github.com/blocktock/go-pkg/ginx"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var UserSet = wire.NewSet(wire.Struct(new(UserHandle), "*"), wire.Bind(new(IUserHandle), new(*UserHandle)))

type UserHandle struct {
	UserService service.IUserService
}

func (uc *UserHandle) CreateUser(c *gin.Context) {

	var requestData schema.CreateUserParams

	// 解析 JSON 请求体
	if err := c.ShouldBindJSON(&requestData); err != nil {
		ginx.ResError(c, 400, err.Error())
		return
	}

	ctx := c.Request.Context()

	// 调用服务层处理请求
	err := uc.UserService.CreateUser(ctx, requestData)
	if err != nil {
		ginx.ResError(c, 500, err.Error())
		return
	}

	ginx.ResSuccess(c, "User created successfully")
}
