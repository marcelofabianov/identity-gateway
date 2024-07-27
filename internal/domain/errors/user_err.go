package errors

import (
	"errors"
)

var (
	ErrUserPasswordHashFailed     = errors.New("error_user_password_hash_failed")
	ErrUserRepositoryCreateFailed = errors.New("error_user_repository_create_failed")
	ErrUserEmailAlreadyExists     = errors.New("error_user_email_already_exists")
)

func NewUserPasswordHashFailedError(err error) error {
	return ErrUserPasswordHashFailed
}

func NewUserRepositoryCreateFailedError(err error) error {
	return ErrUserRepositoryCreateFailed
}

func NewUserEmailAlreadyExistsError(err error) error {
	return ErrUserEmailAlreadyExists
}

func IsUserPasswordHashFailed(err error) bool {
	return errors.Is(err, ErrUserPasswordHashFailed)
}

func IsUserRepositoryCreateFailed(err error) bool {
	return errors.Is(err, ErrUserRepositoryCreateFailed)
}

func IsUserEmailAlreadyExists(err error) bool {
	return errors.Is(err, ErrUserEmailAlreadyExists)
}
