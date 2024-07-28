package mock

import (
	"github.com/stretchr/testify/mock"
)

type MockPasswordHasher struct {
	mock.Mock
}

func (m *MockPasswordHasher) Hash(data string) (string, error) {
	args := m.Called(data)
	return args.String(0), args.Error(1)
}

func (m *MockPasswordHasher) Compare(data, encodedHash string) (bool, error) {
	args := m.Called(data, encodedHash)
	return args.Bool(0), args.Error(1)
}
