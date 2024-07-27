package usecase

import (
	"context"

	"github.com/marcelofabianov/identity-gateway/internal/domain"
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

func (uc *CreateRealmUseCase) Execute(ctx context.Context, input inbound.CreateRealmUseCaseInput) (inbound.CreateRealmUseCaseInput, error) {
	inputRepo := outbound.CreateRealmRepositoryInput{
		ID:                 domain.NewID().String(),
		IdentityProviderID: input.IdentityProviderID.String(),
		Name:               input.Name,
		CreatedAt:          domain.NewCreatedAt().String(),
		UpdatedAt:          domain.NewUpdatedAt().String(),
		Version:            domain.NewVersion().Int(),
	}

	uc.realmRepository.Create(ctx, inputRepo)

	return input, nil
}
