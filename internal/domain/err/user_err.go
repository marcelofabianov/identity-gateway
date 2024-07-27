package err

import (
	"errors"
	"fmt"
)

var (
	ErrUserPasswordHashFailed     = errors.New("error_user_password_hash_failed")
	ErrUserRepositoryCreateFailed = errors.New("error_user_repository_create_failed")
	ErrUserEmailAlreadyExists     = errors.New("error_user_email_already_exists")
)

func NewUserPasswordHashFailedError(err error) error {
	fmt.Println(err)

	return ErrUserPasswordHashFailed
}

func NewUserRepositoryCreateFailedError(err error) error {
	fmt.Println(err)

	return ErrUserRepositoryCreateFailed
}

func NewUserEmailAlreadyExistsError(err error) error {
	fmt.Println(err)

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
