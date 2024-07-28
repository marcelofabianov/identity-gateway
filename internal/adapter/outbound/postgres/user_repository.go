package postgres

import (
	"context"
	"database/sql"

	"github.com/marcelofabianov/identity-gateway/internal/port/outbound"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, input outbound.CreateUserRepositoryInput) error {
	query := `
		INSERT INTO users (id, realm_id, name, email, password, created_at, updated_at, version)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := r.db.ExecContext(ctx, query,
		input.User.ID.String(),
		input.User.RealmID.String(),
		input.User.Name,
		input.User.Email.String(),
		input.User.Password.String(),
		input.User.CreatedAt.String(),
		input.User.UpdatedAt.String(),
		input.User.Version.Int(),
	)

	if err != nil {
		return err
	}

	return nil
}
