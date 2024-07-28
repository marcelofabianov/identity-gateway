package usecase

import (
	"context"

	"github.com/marcelofabianov/identity-gateway/internal/domain"
	"github.com/marcelofabianov/identity-gateway/internal/domain/errors"
	"github.com/marcelofabianov/identity-gateway/internal/port/inbound"
	"github.com/marcelofabianov/identity-gateway/internal/port/outbound"
)

type CreateUserUseCase struct {
	userRepository outbound.CreateUserRepository
	hasher         inbound.PasswordHasher
}

func NewCreateUserUseCase(userRepository outbound.CreateUserRepository, hasher inbound.PasswordHasher) *CreateUserUseCase {
	return &CreateUserUseCase{userRepository, hasher}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, input inbound.CreateUserUseCaseInput) (inbound.CreateUserUseCaseOutput, error) {
	password, err := uc.hasher.Hash(input.Password.String())
	if err != nil {
		return inbound.CreateUserUseCaseOutput{}, errors.NewUserPasswordHashFailedError(err)
	}

	user := domain.User{
		ID:               domain.NewID(),
		RealmID:          input.RealmID,
		Name:             input.Name,
		Email:            input.Email,
		Password:         domain.Password(password),
		DocumentRegistry: input.DocumentRegistry,
		Enabled:          domain.Enabled(false),
		CreatedAt:        domain.NewCreatedAt(),
		UpdatedAt:        domain.NewUpdatedAt(),
		Version:          domain.NewVersion(),
	}

	inputRepo := outbound.CreateUserRepositoryInput{User: user}

	if err := uc.userRepository.Create(ctx, inputRepo); err != nil {
		return inbound.CreateUserUseCaseOutput{}, errors.NewUserRepositoryCreateFailedError(err)
	}

	return inbound.CreateUserUseCaseOutput{User: user}, nil
}
