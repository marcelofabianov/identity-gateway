package outbound

import (
	"context"

	"github.com/marcelofabianov/identity-gateway/internal/domain"
)

// Realm / Repository

type CreateRealmRepositoryInput struct {
	Realm domain.Realm
}

type CreateRealmRepository interface {
	Create(ctx context.Context, input CreateRealmRepositoryInput) error
}

type RealmRepository interface {
	CreateRealmRepository
}

// User / Repository

type CreateUserRepositoryInput struct {
	User domain.User
}

type CreateUserRepository interface {
	Create(ctx context.Context, input CreateUserRepositoryInput) error
}

type UserRepository interface {
	CreateUserRepository
}
