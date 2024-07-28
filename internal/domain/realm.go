package domain

type Realm struct {
	ID                 ID
	IdentityProviderID string
	Name               string
	CreatedAt          CreatedAt
	UpdatedAt          UpdatedAt
	DeletedAt          DeletedAt
	Version            Version
}
