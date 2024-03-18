// internal/http/controllers/user.handle.go

package handle

import (
	"github.com/google/wire"
)

var ControllerSet = wire.NewSet(
	UserSet,
) // end
