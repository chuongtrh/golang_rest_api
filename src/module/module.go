package module

import (
	auth "demo_api/src/module/auth"
	"demo_api/src/module/user"

	"go.uber.org/fx"
)

// Module load all module
var Module = fx.Options(
	user.Module,
	auth.Module,
)
