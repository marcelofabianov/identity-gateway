package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/marcelofabianov/identity-gateway/internal/application/usecase"
	domainError "github.com/marcelofabianov/identity-gateway/internal/domain/errors"
	"github.com/marcelofabianov/identity-gateway/internal/port/inbound"
	"github.com/marcelofabianov/identity-gateway/internal/port/outbound"
	outboundMock "github.com/marcelofabianov/identity-gateway/internal/port/outbound/mock"
)

type CreateRealmUseCaseTestSuite struct {
	suite.Suite
	useCase *usecase.CreateRealmUseCase
	repo    *outboundMock.MockCreateRealmRepository
}

func (s *CreateRealmUseCaseTestSuite) SetupTest() {
	s.repo = new(outboundMock.MockCreateRealmRepository)
	s.useCase = usecase.NewCreateRealmUseCase(s.repo)
}

func (s *CreateRealmUseCaseTestSuite) TearDownTest() {
	s.repo.AssertExpectations(s.T())
}

func (s *CreateRealmUseCaseTestSuite) TestExecute_Success() {
	// Arrange
	input := inbound.CreateRealmUseCaseInput{
		IdentityProviderID: "identity_provider_id",
		Name:               "name",
	}

	ctx := context.Background()

	s.repo.On("Create", ctx, mock.MatchedBy(func(args outbound.CreateRealmRepositoryInput) bool {
		return args.Realm.IdentityProviderID == input.IdentityProviderID && args.Realm.Name == input.Name
	})).Return(nil)

	// Act
	result, err := s.useCase.Execute(ctx, input)

	// Assert
	s.Require().NoError(err)
	s.Equal(input.IdentityProviderID, result.Realm.IdentityProviderID)
	s.Equal(input.Name, result.Realm.Name)
	s.NotNil(result.Realm.ID)
	s.NotNil(result.Realm.CreatedAt)
	s.NotNil(result.Realm.UpdatedAt)
	s.NotNil(result.Realm.Version)
}

func (s *CreateRealmUseCaseTestSuite) TestExecute_Fail_RealmRepositoryCreateFailedError() {
	// Arrange
	input := inbound.CreateRealmUseCaseInput{
		IdentityProviderID: "identity_provider_id",
		Name:               "name",
	}

	ctx := context.Background()

	s.repo.On("Create", ctx, mock.Anything).Return(domainError.NewRealmRepositoryCreateFailedError(nil))

	// Act
	result, err := s.useCase.Execute(ctx, input)

	// Assert
	s.Error(err)
	s.Equal(domainError.ErrRealmRepositoryCreateFailed, err)
	s.Equal(inbound.CreateRealmUseCaseOutput{}, result)
}

func TestCreateRealmUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CreateRealmUseCaseTestSuite))
}
