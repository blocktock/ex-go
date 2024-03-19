// internal/http/controllers/user.handle.go

package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ControllerSet = wire.NewSet(
	UserSet,
) // end

type IUserHandle interface {
	//创建用户
	CreateUser(c *gin.Context)
}
