package service

import (
	"context"

	"github.com/marcelofabianov/identity-gateway/internal/domain"
	"github.com/marcelofabianov/identity-gateway/internal/port/inbound"
)

type UserService struct {
	createUser inbound.CreateUserUseCase
}

func NewUserService(createUser inbound.CreateUserUseCase) *UserService {
	return &UserService{
		createUser: createUser,
	}
}

func (s *UserService) Create(ctx context.Context, input inbound.CreateUserServiceInput) (inbound.CreateUserServiceOutput, error) {
	inputUC := inbound.CreateUserUseCaseInput{
		RealmID:          input.RealmID,
		Name:             input.Name,
		Email:            domain.Email(input.Email),
		Password:         domain.Password(input.Password),
		DocumentRegistry: domain.DocumentRegistry(input.DocumentRegistry),
	}

	outputUC, err := s.createUser.Execute(ctx, inputUC)
	if err != nil {
		return inbound.CreateUserServiceOutput{}, err
	}

	// Todo: dispatch user.created event
	// Todo: Logger

	return inbound.CreateUserServiceOutput{
		ID:               outputUC.User.ID,
		RealmID:          outputUC.User.RealmID,
		Name:             outputUC.User.Name,
		Email:            outputUC.User.Email,
		DocumentRegistry: outputUC.User.DocumentRegistry,
		Enabled:          outputUC.User.Enabled,
		CreatedAt:        outputUC.User.CreatedAt,
		UpdatedAt:        outputUC.User.UpdatedAt,
		Version:          outputUC.User.Version,
		DeletedAt:        nil,
	}, nil
}
