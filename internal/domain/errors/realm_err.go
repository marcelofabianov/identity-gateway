package errors

import "errors"

var (
	ErrRealmRepositoryCreateFailed = errors.New("error_realm_repository_create_failed")
)

func NewRealmRepositoryCreateFailedError(err error) error {
	return ErrRealmRepositoryCreateFailed
}

func IsRealmRepositoryCreateFailed(err error) bool {
	return errors.Is(err, ErrRealmRepositoryCreateFailed)
}
