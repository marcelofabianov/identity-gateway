package mock

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/marcelofabianov/identity-gateway/internal/port/inbound"
)

type MockCreateUserUseCase struct {
	mock.Mock
}

func (m *MockCreateUserUseCase) Execute(ctx context.Context, input inbound.CreateUserUseCaseInput) (inbound.CreateUserUseCaseOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(inbound.CreateUserUseCaseOutput), args.Error(1)
}
