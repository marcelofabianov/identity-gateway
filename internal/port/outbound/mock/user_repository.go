package mock

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/marcelofabianov/identity-gateway/internal/port/outbound"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, input outbound.CreateUserRepositoryInput) error {
	args := m.Called(ctx, input)
	return args.Error(0)
}
