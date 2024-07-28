package container

import (
	"database/sql"

	"go.uber.org/dig"

	"github.com/marcelofabianov/identity-gateway/internal/adapter/outbound/postgres"
	"github.com/marcelofabianov/identity-gateway/internal/application/service"
	"github.com/marcelofabianov/identity-gateway/internal/application/usecase"
	"github.com/marcelofabianov/identity-gateway/internal/port/inbound"
	"github.com/marcelofabianov/identity-gateway/internal/port/outbound"
)

type RealmContainer struct {
	*dig.Container
}

func NewRealmContainer(db *sql.DB) *RealmContainer {
	container := dig.New()

	registerRepositories(container, db)
	registerUseCase(container)
	registerServices(container)

	return &RealmContainer{container}
}

func registerRepositories(c *dig.Container, db *sql.DB) {
	c.Provide(func() outbound.RealmRepository {
		return postgres.NewRealmRepository(db)
	})
}

func registerUseCase(c *dig.Container) {
	c.Provide(func(repo outbound.RealmRepository) inbound.CreateRealmUseCase {
		return usecase.NewCreateRealmUseCase(repo)
	})
}

func registerServices(c *dig.Container) {
	c.Provide(func(createRealm inbound.CreateRealmUseCase) inbound.RealmService {
		return service.NewRealmService(createRealm)
	})
}
