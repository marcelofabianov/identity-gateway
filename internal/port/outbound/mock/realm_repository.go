package mock

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/marcelofabianov/identity-gateway/internal/port/outbound"
)

type MockCreateRealmRepository struct {
	mock.Mock
}

func (m *MockCreateRealmRepository) Create(ctx context.Context, input outbound.CreateRealmRepositoryInput) error {
	args := m.Called(ctx, input)
	return args.Error(0)
}
