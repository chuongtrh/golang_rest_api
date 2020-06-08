package user

import (
	"go.uber.org/fx"
)

// Module user
var Module = fx.Options(
	fx.Provide(NewUserRepository),
	fx.Provide(NewUserService),
	fx.Provide(NewUserController),
	fx.Invoke(LoadRoute),
)
