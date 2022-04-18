package repository

import "go.uber.org/fx"

// Repositories module
var NewRepository = fx.Options(
	fx.Provide(NewArticleRepository),
)
