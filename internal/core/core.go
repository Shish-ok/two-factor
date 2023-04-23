package core

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"two-factor-auth/internal/api"
	"two-factor-auth/internal/config"
	"two-factor-auth/internal/services/auth"
	"two-factor-auth/internal/services/redis_save"
	redis_storage "two-factor-auth/internal/storage/redis"
)

func Core() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(redis_storage.NewRedisStorage, fx.As(new(auth.Storage))),
			fx.Annotate(redis_storage.NewRedisStorage, fx.As(new(redis_save.Storage))),
		),
		fx.Provide(
			context.Background,
			redis_storage.NewRedisStorage,
			config.NewConfig,
			api.NewAPI,
			gin.Default,
			auth.NewService,
			redis_save.NewService,
		),
		fx.Invoke(restAPIHook),
	)
}

func restAPIHook(lifecycle fx.Lifecycle, api *api.API, save *redis_save.Service) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				save.RolloutData(ctx)
				go api.Run()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				save.SaveData(ctx)
				ctx.Done()
				return nil
			},
		},
	)
}
