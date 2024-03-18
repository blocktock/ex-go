package cache

import (
	"github.com/google/wire"
)

var CacheRepoSet = wire.NewSet(
	UserSet,
) // end

type IUserCacheRepository interface {
	Save() error
}
