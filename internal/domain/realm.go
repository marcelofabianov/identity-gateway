package domain

type Realm struct {
	ID                 ID
	IdentityProviderID ID
	Name               string
	CreatedAt          CreatedAt
	UpdatedAt          UpdatedAt
	DeletedAt          DeletedAt
	Version            Version
}
