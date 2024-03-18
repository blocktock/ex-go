//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package app

import (
	"ex-go/internal/domain/service"
	"ex-go/internal/http/handle"
	"ex-go/internal/http/route"
	"ex-go/internal/repository/cache"
	"ex-go/internal/repository/db"
	"github.com/google/wire"
)

func InitServer() (*App, func(), error) {
	wire.Build(
		AppSet,
		InitDB,
		InitCache,
		db.DBRepoSet,
		cache.CacheRepoSet,
		service.ServiceSet,
		handle.ControllerSet,
		route.InitEngine,
		route.RouterSet,
	)
	return new(App), nil, nil
}
