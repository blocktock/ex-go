package handle

import (
	"ex-go/internal/domain/service"
	"ex-go/internal/schema"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var UserSet = wire.NewSet(wire.Struct(new(UserHandle), "*"))

type UserHandle struct {
	UserService service.IUserService
}

func (uc *UserHandle) CreateUser(c *gin.Context) {

	var requestData schema.CreateUserRequestDTO

	// 解析 JSON 请求体
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 调用服务层处理请求
	err := uc.UserService.CreateUser(requestData)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User created successfully"})
}
