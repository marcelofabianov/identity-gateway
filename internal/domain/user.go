package domain

type User struct {
	ID        ID
	RealmID   ID
	Name      string
	Email     Email
	Password  Password
	EnabledAt EnabledAt
	CreatedAt CreatedAt
	UpdatedAt UpdatedAt
	DeletedAt DeletedAt
	Version   Version
}
