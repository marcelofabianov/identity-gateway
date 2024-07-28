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

	realmRegisterRepositories(container, db)
	realmRegisterUseCase(container)
	realmRegisterServices(container)

	return &RealmContainer{container}
}

func realmRegisterRepositories(c *dig.Container, db *sql.DB) {
	c.Provide(func() outbound.RealmRepository {
		return postgres.NewRealmRepository(db)
	})
}

func realmRegisterUseCase(c *dig.Container) {
	c.Provide(func(repo outbound.RealmRepository) inbound.CreateRealmUseCase {
		return usecase.NewCreateRealmUseCase(repo)
	})
}

func realmRegisterServices(c *dig.Container) {
	c.Provide(func(createRealm inbound.CreateRealmUseCase) inbound.RealmService {
		return service.NewRealmService(createRealm)
	})
}
