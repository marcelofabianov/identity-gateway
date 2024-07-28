package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/marcelofabianov/identity-gateway/internal/application/service"
	"github.com/marcelofabianov/identity-gateway/internal/domain"
	"github.com/marcelofabianov/identity-gateway/internal/domain/errors"
	"github.com/marcelofabianov/identity-gateway/internal/port/inbound"
	inboundMock "github.com/marcelofabianov/identity-gateway/internal/port/inbound/mock"
)

type RealmServiceTestSuite struct {
	suite.Suite
	createRealmMock *inboundMock.MockCreateRealmUseCase
	service         *service.RealmService
}

func (s *RealmServiceTestSuite) SetupTest() {
	s.createRealmMock = new(inboundMock.MockCreateRealmUseCase)
	s.service = service.NewRealmService(s.createRealmMock)
}

func (s *RealmServiceTestSuite) TearDownTest() {
	s.createRealmMock.AssertExpectations(s.T())
}

func (s *RealmServiceTestSuite) TestCreate_Success() {
	//Arrange
	inputSe := inbound.CreateRealmServiceInput{
		IdentityProviderID: "identity_provider_id",
		Name:               "name",
	}

	ctx := context.Background()

	outboundUC := inbound.CreateRealmUseCaseOutput{
		Realm: domain.Realm{
			ID:                 domain.NewID(),
			IdentityProviderID: inputSe.IdentityProviderID,
			Name:               inputSe.Name,
			CreatedAt:          domain.NewCreatedAt(),
			UpdatedAt:          domain.NewUpdatedAt(),
			Version:            domain.NewVersion(),
		},
	}

	s.createRealmMock.On("Execute", ctx, mock.MatchedBy(func(args inbound.CreateRealmUseCaseInput) bool {
		return args.IdentityProviderID == inputSe.IdentityProviderID && args.Name == inputSe.Name
	})).Return(outboundUC, nil)

	//Act
	result, err := s.service.Create(ctx, inputSe)

	//Assert
	s.Require().NoError(err)
	s.Equal(inputSe.IdentityProviderID, result.IdentityProviderID)
	s.Equal(inputSe.Name, result.Name)
	s.Equal(outboundUC.Realm.ID, result.ID)
	s.Equal(outboundUC.Realm.IdentityProviderID, result.IdentityProviderID)
	s.Equal(outboundUC.Realm.Name, result.Name)
	s.Equal(outboundUC.Realm.CreatedAt, result.CreatedAt)
	s.Equal(outboundUC.Realm.UpdatedAt, result.UpdatedAt)
	s.Nil(result.DeletedAt)
	s.Equal(outboundUC.Realm.Version, result.Version)
}

func (s *RealmServiceTestSuite) TestCreate_Error() {
	//Arrange
	inputSe := inbound.CreateRealmServiceInput{
		IdentityProviderID: "identity_provider_id",
		Name:               "name",
	}

	ctx := context.Background()

	outboundUC := inbound.CreateRealmUseCaseOutput{}

	error := errors.NewRealmRepositoryCreateFailedError(nil)

	s.createRealmMock.On("Execute", ctx, mock.Anything).Return(outboundUC, error)

	//Act
	result, err := s.service.Create(ctx, inputSe)

	//Assert
	s.Error(err)
	s.Equal(error, err)
	s.Equal(inbound.CreateRealmServiceOutput{}, result)
}

func TestRealmServiceTestSuite(t *testing.T) {
	suite.Run(t, new(RealmServiceTestSuite))
}
