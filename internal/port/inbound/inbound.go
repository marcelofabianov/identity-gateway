package inbound

import (
	"context"

	"github.com/marcelofabianov/identity-gateway/internal/domain"
)

// Realm / UseCase

type CreateRealmUseCaseInput struct {
	IdentityProviderID domain.ID
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
	IdentityProviderID domain.ID
	Name               string
}

type CreateRealmServiceOutput struct {
	ID                 domain.ID
	IdentityProviderID domain.ID
	Name               string
	Enabled            domain.EnabledAt
	CreatedAt          domain.CreatedAt
	UpdatedAt          domain.UpdatedAt
	DeletedAt          domain.DeletedAt
	Version            domain.Version
}

type RealmService interface {
	Create(ctx context.Context, input CreateRealmServiceInput) (CreateRealmServiceOutput, error)
}
