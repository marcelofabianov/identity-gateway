package domain

type User struct {
	ID               ID
	RealmID          string
	Name             string
	Email            Email
	Password         Password
	DocumentRegistry DocumentRegistry
	Enabled          Enabled
	CreatedAt        CreatedAt
	UpdatedAt        UpdatedAt
	DeletedAt        DeletedAt
	Version          Version
}
