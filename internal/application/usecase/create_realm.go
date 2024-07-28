package usecase

import (
	"context"

	"github.com/marcelofabianov/identity-gateway/internal/domain"
	"github.com/marcelofabianov/identity-gateway/internal/domain/errors"
	"github.com/marcelofabianov/identity-gateway/internal/port/inbound"
	"github.com/marcelofabianov/identity-gateway/internal/port/outbound"
)

type CreateRealmUseCase struct {
	realmRepository outbound.CreateRealmRepository
}

func NewCreateRealmUseCase(realmRepository outbound.CreateRealmRepository) *CreateRealmUseCase {
	return &CreateRealmUseCase{
		realmRepository: realmRepository,
	}
}

func (uc *CreateRealmUseCase) Execute(ctx context.Context, input inbound.CreateRealmUseCaseInput) (inbound.CreateRealmUseCaseOutput, error) {
	realm := domain.Realm{
		ID:                 domain.NewID(),
		IdentityProviderID: input.IdentityProviderID,
		Name:               input.Name,
		CreatedAt:          domain.NewCreatedAt(),
		UpdatedAt:          domain.NewUpdatedAt(),
		Version:            domain.NewVersion(),
	}

	inputRepo := outbound.CreateRealmRepositoryInput{Realm: realm}

	if err := uc.realmRepository.Create(ctx, inputRepo); err != nil {
		return inbound.CreateRealmUseCaseOutput{}, errors.NewRealmRepositoryCreateFailedError(err)
	}

	return inbound.CreateRealmUseCaseOutput{Realm: realm}, nil
}
