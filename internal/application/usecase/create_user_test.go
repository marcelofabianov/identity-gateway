package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/marcelofabianov/identity-gateway/internal/application/usecase"
	"github.com/marcelofabianov/identity-gateway/internal/domain"
	"github.com/marcelofabianov/identity-gateway/internal/port/inbound"
	inboundMock "github.com/marcelofabianov/identity-gateway/internal/port/inbound/mock"
	"github.com/marcelofabianov/identity-gateway/internal/port/outbound"
	outboundMock "github.com/marcelofabianov/identity-gateway/internal/port/outbound/mock"
)

type CreateUserUseCaseTestSuite struct {
	suite.Suite
	useCase *usecase.CreateUserUseCase
	repo    *outboundMock.MockUserRepository
	hasher  *inboundMock.MockPasswordHasher
}

func (s *CreateUserUseCaseTestSuite) SetupTest() {
	s.repo = new(outboundMock.MockUserRepository)
	s.hasher = new(inboundMock.MockPasswordHasher)
	s.useCase = usecase.NewCreateUserUseCase(s.repo, s.hasher)
}

func (s *CreateUserUseCaseTestSuite) TearDownTest() {
	s.repo.AssertExpectations(s.T())
	s.hasher.AssertExpectations(s.T())
}

func (s *CreateUserUseCaseTestSuite) TestExecute_Success() {
	//Arrange
	inputUC := inbound.CreateUserUseCaseInput{
		RealmID:          "realm_id",
		Name:             "Marcelo",
		Email:            domain.Email("marcelo@email.com"),
		Password:         domain.Password("plain-password"),
		DocumentRegistry: domain.DocumentRegistry("01234567890"),
	}

	hashedPassword := "hashed-password"

	ctx := context.Background()

	s.hasher.On("Hash", inputUC.Password.String()).Return(hashedPassword, nil)

	s.repo.On("Create", ctx, mock.MatchedBy(func(args outbound.CreateUserRepositoryInput) bool {
		return args.User.RealmID == inputUC.RealmID && args.User.Name == inputUC.Name && args.User.Email == inputUC.Email && args.User.DocumentRegistry == inputUC.DocumentRegistry
	})).Return(nil)

	// Act
	result, err := s.useCase.Execute(ctx, inputUC)

	// Assert
	s.Require().NoError(err)
	s.Equal(inputUC.RealmID, result.User.RealmID)
	s.Equal(inputUC.Name, result.User.Name)
	s.Equal(inputUC.Email, result.User.Email)
	s.Equal(inputUC.DocumentRegistry, result.User.DocumentRegistry)
	s.NotNil(result.User.ID)
	s.NotNil(result.User.CreatedAt)
	s.NotNil(result.User.UpdatedAt)
	s.Equal(result.User.Version, domain.Version(0))
}

func (s *CreateRealmUseCaseTestSuite) TestExecute_Fail_UserPasswordHashFailedError() {
	//...
}

func (s *CreateRealmUseCaseTestSuite) TestExecute_Fail_UserRepositoryCreateFailedError() {
	//...
}

func TestCreateUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CreateUserUseCaseTestSuite))
}
