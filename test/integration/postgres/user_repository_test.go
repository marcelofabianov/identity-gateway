package postgres_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/marcelofabianov/identity-gateway/internal/adapter/outbound/postgres"
	"github.com/marcelofabianov/identity-gateway/internal/domain"
	"github.com/marcelofabianov/identity-gateway/internal/port/outbound"
	"github.com/marcelofabianov/identity-gateway/test/container"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	container *container.PostgresContainer
}

func (s *UserRepositoryTestSuite) SetupSuite() {
	ctx := context.Background()
	s.container = container.NewPostgresContainer(s.T(), ctx)
}

func (s *UserRepositoryTestSuite) TearDownSuite() {
	ctx := context.Background()
	if s.container != nil && s.container.Container != nil {
		if err := s.container.Container.Terminate(ctx); err != nil {
			s.T().Logf("Error terminating PostgreSQL container: %v", err)
		}
	}
}

func NewRealm(db *sql.DB) (string, error) {
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

	ctx := context.Background()
	err := repo.Create(ctx, input)
	if err != nil {
		return "", err
	}

	return realm.ID.String(), nil
}

func (s *UserRepositoryTestSuite) TestCreateUser_Success() {
	//Arrange
	ctx := context.Background()
	db, err := sql.Open("postgres", s.container.ConnStr)
	if err != nil {
		s.T().Fatalf("Error opening connection to PostgreSQL: %v", err)
	}
	defer db.Close()

	realmID, err := NewRealm(db)
	if err != nil {
		s.T().Fatalf("Error creating realm: %v", err)
	}

	user := domain.User{
		ID:               domain.NewID(),
		RealmID:          realmID,
		Name:             "Marcelo",
		Email:            domain.Email("marcelo@email.com"),
		Password:         domain.Password("hash-password"),
		DocumentRegistry: domain.DocumentRegistry("01234567890"),
		Enabled:          domain.Enabled(true),
		CreatedAt:        domain.NewCreatedAt(),
		UpdatedAt:        domain.NewUpdatedAt(),
		Version:          domain.NewVersion(),
	}

	inputRepo := outbound.CreateUserRepositoryInput{User: user}

	repo := postgres.NewUserRepository(db)

	//Act
	err = repo.Create(ctx, inputRepo)

	if err != nil {
		s.T().Fatalf("Error creating user: %v", err)
	}

	// Assert
	s.NoError(err)

	query := `SELECT id FROM users WHERE id = $1`

	var id string
	err = db.QueryRowContext(ctx, query, inputRepo.User.ID).Scan(&id)

	s.NoError(err)
	s.Equal(inputRepo.User.ID.String(), id)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
