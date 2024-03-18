package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

var UserSet = wire.NewSet(wire.Struct(new(UserCacheRepository), "*"), wire.Bind(new(IUserCacheRepository), new(*UserCacheRepository)))

type UserCacheRepository struct {
	Cache *redis.Client
}

func (r *UserCacheRepository) Save() error {
	// 实现保存用户到 MySQL 数据库的逻辑
	return nil
}
