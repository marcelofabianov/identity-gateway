package service

import (
	"context"

	"github.com/marcelofabianov/identity-gateway/internal/port/inbound"
)

type RealmService struct {
	createRealm inbound.CreateRealmUseCase
}

func NewRealmService(createRealm inbound.CreateRealmUseCase) *RealmService {
	return &RealmService{
		createRealm: createRealm,
	}
}

func (s *RealmService) Create(ctx context.Context, input inbound.CreateRealmServiceInput) (inbound.CreateRealmServiceOutput, error) {
	inputUC := inbound.CreateRealmUseCaseInput{
		IdentityProviderID: input.IdentityProviderID,
		Name:               input.Name,
	}

	outputUC, err := s.createRealm.Execute(ctx, inputUC)
	if err != nil {
		return inbound.CreateRealmServiceOutput{}, err
	}

	// Todo: dispatch realm.created event

	return inbound.CreateRealmServiceOutput{
		ID:                 outputUC.Realm.ID,
		IdentityProviderID: outputUC.Realm.IdentityProviderID,
		Name:               outputUC.Realm.Name,
		CreatedAt:          outputUC.Realm.CreatedAt,
		UpdatedAt:          outputUC.Realm.UpdatedAt,
		DeletedAt:          nil,
		Version:            outputUC.Realm.Version,
	}, nil
}
