package route

import (
	"ex-go/internal/http/handle"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"))

type Router struct {
	User *handle.UserHandle
}

func (r *Router) RegisterApi(engine *gin.Engine) {
	engine.POST("/user", r.User.CreateUser)
}
