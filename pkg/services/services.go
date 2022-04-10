package services

import "go.uber.org/fx"

// Services module
var NewService = fx.Options(
	fx.Provide(NewArticleService),
)
