package user

import (
	"github.com/jackc/pgx/v5"
	"go.uber.org/dig"
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
	//...
}

func registerUseCases(container *dig.Container) {
	//...
}

func registerService(container *dig.Container) {
	//...
}
