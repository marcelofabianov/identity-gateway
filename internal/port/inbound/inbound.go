package inbound

import (
	"context"

	"github.com/marcelofabianov/identity-gateway/internal/domain"
)

// Realm / UseCase

type CreateRealmUseCaseInput struct {
	IdentityProviderID string
	Name               string
}

type CreateRealmUseCaseOutput struct {
	Realm domain.Realm
}

type CreateRealmUseCase interface {
	Execute(ctx context.Context, input CreateRealmUseCaseInput) (CreateRealmUseCaseOutput, error)
}

// Realm / Service

type CreateRealmServiceInput struct {
	IdentityProviderID string
	Name               string
}

type CreateRealmServiceOutput struct {
	ID                 domain.ID
	IdentityProviderID string
	Name               string
	CreatedAt          domain.CreatedAt
	UpdatedAt          domain.UpdatedAt
	DeletedAt          domain.DeletedAt
	Version            domain.Version
}

type RealmService interface {
	Create(ctx context.Context, input CreateRealmServiceInput) (CreateRealmServiceOutput, error)
}
