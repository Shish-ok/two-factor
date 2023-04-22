package core

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"two-factor-auth/internal/api"
	"two-factor-auth/internal/config"
	"two-factor-auth/internal/services/auth"
	redis_storage "two-factor-auth/internal/storage/redis"
)

func Core() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(redis_storage.NewRedisStorage, fx.As(new(auth.Storage))),
		),
		fx.Provide(
			context.Background,
			redis_storage.NewRedisStorage,
			config.NewConfig,
			api.NewAPI,
			gin.Default,
			auth.NewService,
		),
		fx.Invoke(restAPIHook),
	)
}

func restAPIHook(lifecycle fx.Lifecycle, api *api.API) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go api.Run()
				return nil
			},
		},
	)
}
