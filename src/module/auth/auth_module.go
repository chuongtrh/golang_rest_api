package auth

import (
	"go.uber.org/fx"
)

// Module auth
var Module = fx.Options(

	fx.Provide(NewAuthController),
	fx.Invoke(LoadRoute),
)
