package postgres_test

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"

	"github.com/marcelofabianov/identity-gateway/internal/adapter/outbound/postgres"
	"github.com/marcelofabianov/identity-gateway/internal/domain"
	"github.com/marcelofabianov/identity-gateway/internal/port/outbound"
	"github.com/marcelofabianov/identity-gateway/test/container"
)

type RealmRepositoryTestSuite struct {
	suite.Suite
	container *container.PostgresContainer
}

func (s *RealmRepositoryTestSuite) SetupSuite() {
	ctx := context.Background()
	s.container = container.NewPostgresContainer(s.T(), ctx)
}

func (s *RealmRepositoryTestSuite) TearDownSuite() {
	ctx := context.Background()
	if s.container != nil && s.container.Container != nil {
		if err := s.container.Container.Terminate(ctx); err != nil {
			s.T().Logf("Error terminating PostgreSQL container: %v", err)
		}
	}
}

func (s *RealmRepositoryTestSuite) TestCreateRealm_Success() {
	// Arrange
	ctx := context.Background()
	db, err := sql.Open("postgres", s.container.ConnStr)
	if err != nil {
		s.T().Fatalf("Error opening connection to PostgreSQL: %v", err)
	}
	defer db.Close()

	realm := domain.Realm{
		ID:                 domain.NewID(),
		IdentityProviderID: domain.NewID().String(),
		Name:               "realm-name",
		CreatedAt:          domain.NewCreatedAt(),
		UpdatedAt:          domain.NewUpdatedAt(),
		Version:            domain.NewVersion(),
	}

	input := outbound.CreateRealmRepositoryInput{Realm: realm}

	repo := postgres.NewRealmRepository(db)

	// Act
	err = repo.Create(ctx, input)

	if err != nil {
		s.T().Fatalf("Error creating realm: %v", err)
	}

	// Assert
	s.NoError(err)

	query := `SELECT id FROM realms WHERE id = $1`

	var id string
	err = db.QueryRowContext(ctx, query, input.Realm.ID).Scan(&id)

	s.NoError(err)
	s.Equal(input.Realm.ID.String(), id)
}

func TestRealmRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RealmRepositoryTestSuite))
}
