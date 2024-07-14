package user

import (
	"github.com/jackc/pgx/v5"
	"go.uber.org/dig"

	repository "github.com/marcelofabianov/pejota/internal/user/adapter/pgx"
	"github.com/marcelofabianov/pejota/internal/user/application"
	"github.com/marcelofabianov/pejota/internal/user/domain/usecase"
	"github.com/marcelofabianov/pejota/internal/user/port"
)

type UserContainer struct {
	*dig.Container
}

func NewUserContainer(db *pgx.Conn) *UserContainer {
	container := dig.New()

	registerRepositories(container, db)
	registerUseCases(container)
	registerService(container)

	return &UserContainer{Container: container}
}

func registerRepositories(container *dig.Container, db *pgx.Conn) {
	container.Provide(func() port.UserRepository {
		return repository.NewUserRepository(db)
	})
}

func registerUseCases(container *dig.Container) {
	container.Provide(func(repo port.UserRepository) port.GetUserUseCase {
		return usecase.NewGetUserUseCase(repo)
	})
}

func registerService(container *dig.Container) {
	container.Provide(func(getUser port.GetUserUseCase) port.UserServiceApplication {
		return application.NewUserServiceApplication(getUser)
	})
}