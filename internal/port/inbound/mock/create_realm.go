package mock

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/marcelofabianov/identity-gateway/internal/port/inbound"
)

type MockCreateRealmUseCase struct {
	mock.Mock
}

func (m *MockCreateRealmUseCase) Execute(ctx context.Context, input inbound.CreateRealmUseCaseInput) (inbound.CreateRealmUseCaseOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(inbound.CreateRealmUseCaseOutput), args.Error(1)
}
