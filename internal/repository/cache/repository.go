package cache

import (
	"context"
	"github.com/google/wire"
)

var CacheRepoSet = wire.NewSet(
	UserSet,
) // end

type IUserCacheRepository interface {
	Save(ctx context.Context) error
}
