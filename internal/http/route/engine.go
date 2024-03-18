package route

import (
	"ex-go/internal/http/middleware"
	"github.com/gin-gonic/gin"
)

func InitEngine(router *Router) *gin.Engine {

	app := gin.New()

	app.Use(middleware.CORSMiddleware())

	router.RegisterApi(app)

	return app
}
