package domain

type Realm struct {
	ID                 string
	IdentityProviderID string
	Name               string
	EnabledAt          EnabledAt
	CreatedAt          CreatedAt
	UpdatedAt          UpdatedAt
	DeletedAt          DeletedAt
	Version            Version
}
