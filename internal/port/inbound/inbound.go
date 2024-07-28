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

// User / UseCase

// PKG
type PasswordHasher interface {
	Hash(data string) (string, error)
	Compare(data, encodedHash string) (bool, error)
}

type CreateUserUseCaseInput struct {
	RealmID          string
	Name             string
	Email            domain.Email
	Password         domain.Password
	DocumentRegistry domain.DocumentRegistry
}

type CreateUserUseCaseOutput struct {
	User domain.User
}

type CreateUserUseCase interface {
	Execute(ctx context.Context, input CreateUserUseCaseInput) (CreateUserUseCaseOutput, error)
}

// User / Service

type CreateUserServiceInput struct {
	RealmID          string
	Name             string
	Email            string
	Password         string
	DocumentRegistry string
}

type CreateUserServiceOutput struct {
	ID               domain.ID
	RealmID          string
	Name             string
	Email            domain.Email
	DocumentRegistry domain.DocumentRegistry
	Enabled          domain.Enabled
	CreatedAt        domain.CreatedAt
	UpdatedAt        domain.UpdatedAt
	DeletedAt        domain.DeletedAt
	Version          domain.Version
}

type UserService interface {
	Create(ctx context.Context, input CreateUserServiceInput) (CreateUserServiceOutput, error)
}
