package route

import (
	"ex-go/internal/http/middleware"
	"github.com/gin-gonic/gin"
)

func InitEngine(router *Router) *gin.Engine {

	app := gin.New()

	app.NoMethod(middleware.NoMethodMiddleware())
	app.NoRoute(middleware.NoRouteMiddleware())

	app.Use(middleware.CORSMiddleware())
	app.Use(middleware.TraceMiddleware())

	router.RegisterApi(app)

	return app
}
