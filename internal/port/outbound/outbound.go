package outbound

import (
	"context"
)

// Realm / Repository

type CreateRealmRepositoryInput struct {
	ID                 string
	IdentityProviderID string
	Name               string
	CreatedAt          string
	UpdatedAt          string
	Version            int64
}

type CreateRealmRepository interface {
	Create(ctx context.Context, input CreateRealmRepositoryInput) error
}

type RealmRepository interface {
	CreateRealmRepository
}
