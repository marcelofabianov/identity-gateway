package container

import (
	"database/sql"

	"go.uber.org/dig"

	"github.com/marcelofabianov/identity-gateway/internal/adapter/outbound/postgres"
	"github.com/marcelofabianov/identity-gateway/internal/application/usecase"
	"github.com/marcelofabianov/identity-gateway/internal/port/inbound"
	"github.com/marcelofabianov/identity-gateway/internal/port/outbound"
	"github.com/marcelofabianov/identity-gateway/pkg/hasher"
)

type UserContainer struct {
	*dig.Container
}

func NewUserContainer(db *sql.DB) *UserContainer {
	container := dig.New()

	userRegisterPackages(container)
	userRegisterRepositories(container, db)
	userRegisterUseCase(container)

	return &UserContainer{container}
}

func userRegisterPackages(c *dig.Container) {
	c.Provide(func() inbound.PasswordHasher {
		return hasher.NewHasher()
	})
}

func userRegisterRepositories(c *dig.Container, db *sql.DB) {
	c.Provide(func() outbound.UserRepository {
		return postgres.NewUserRepository(db)
	})
}

func userRegisterUseCase(c *dig.Container) {
	c.Provide(func(repo outbound.UserRepository, hasher inbound.PasswordHasher) inbound.CreateUserUseCase {
		return usecase.NewCreateUserUseCase(repo, hasher)
	})
}
