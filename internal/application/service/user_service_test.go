package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/marcelofabianov/identity-gateway/internal/application/service"
	"github.com/marcelofabianov/identity-gateway/internal/domain"
	"github.com/marcelofabianov/identity-gateway/internal/port/inbound"
	inboundMock "github.com/marcelofabianov/identity-gateway/internal/port/inbound/mock"
)

type UserServiceTestSuite struct {
	suite.Suite
	createUserMock *inboundMock.MockCreateUserUseCase
	service        *service.UserService
}

func (s *UserServiceTestSuite) SetupTest() {
	s.createUserMock = new(inboundMock.MockCreateUserUseCase)
	s.service = service.NewUserService(s.createUserMock)
}

func (s *UserServiceTestSuite) TearDownTest() {
	s.createUserMock.AssertExpectations(s.T())
}

func (s *UserServiceTestSuite) TestCreate_Success() {
	// Arrange
	inputSe := inbound.CreateUserServiceInput{
		RealmID:          "realm_id",
		Name:             "Marcelo",
		Email:            "marcelo@email.com",
		Password:         "plain-password",
		DocumentRegistry: "01234567890",
	}

	ctx := context.Background()

	outputUC := inbound.CreateUserUseCaseOutput{
		User: domain.User{
			ID:               domain.NewID(),
			RealmID:          inputSe.RealmID,
			Name:             inputSe.Name,
			Email:            domain.Email(inputSe.Email),
			Password:         domain.Password(inputSe.Password),
			DocumentRegistry: domain.DocumentRegistry(inputSe.DocumentRegistry),
		},
	}

	s.createUserMock.On("Execute", ctx, mock.MatchedBy(func(args inbound.CreateUserUseCaseInput) bool {
		return args.RealmID == inputSe.RealmID && args.Name == inputSe.Name && args.Email == domain.Email(inputSe.Email) && args.Password == domain.Password(inputSe.Password) && args.DocumentRegistry == domain.DocumentRegistry(inputSe.DocumentRegistry)
	})).Return(outputUC, nil)

	// Act
	result, err := s.service.Create(ctx, inputSe)

	// Assert
	s.Require().NoError(err)
	s.Equal(inputSe.RealmID, result.RealmID)
	s.Equal(inputSe.Name, result.Name)
	s.Equal(inputSe.Email, result.Email.String())
	s.Equal(inputSe.DocumentRegistry, result.DocumentRegistry.String())
	s.Equal(outputUC.User.ID, result.ID)
	s.Equal(outputUC.User.RealmID, result.RealmID)
	s.Equal(outputUC.User.Name, result.Name)
	s.Equal(outputUC.User.Email, result.Email)
	s.Equal(outputUC.User.DocumentRegistry, result.DocumentRegistry)
	s.Equal(domain.Enabled(false), result.Enabled)
	s.NotNil(result.CreatedAt)
	s.NotNil(result.UpdatedAt)
	s.Nil(result.DeletedAt)
	s.Equal(result.Version, domain.Version(0))
}

func (s *UserServiceTestSuite) TestCreate_Fail_CreateUserUseCaseError() {
	//...
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
