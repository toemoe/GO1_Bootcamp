package di

import (
	"context"
	"net/http"
	"tic-tac-toe/internal/api"
	"tic-tac-toe/internal/domain/repository"
	service "tic-tac-toe/internal/domain/service"
	app "tic-tac-toe/internal/game"
	repo "tic-tac-toe/internal/storage"

	"go.uber.org/fx"
)

func BuildContainer() *fx.App {
	return fx.New(
		fx.Provide(
			fx.Annotate(repo.NewMemoryRepository, fx.As(new(repository.GameRepository))),
			fx.Annotate(service.NewMinimaxService, fx.As(new(service.GameService))),
			app.NewUseCase,
			api.NewHandler,
		),
		fx.Invoke(func(lc fx.Lifecycle, h *api.Handler) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					api.RegisterRoutes(h)
					go http.ListenAndServe(":8080", nil)
					return nil
				},
			})
		}),
	)
}
