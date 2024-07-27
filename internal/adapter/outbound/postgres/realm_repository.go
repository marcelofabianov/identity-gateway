package postgres

import (
	"context"
	"database/sql"

	"github.com/marcelofabianov/identity-gateway/internal/port/outbound"
)

type RealmRepository struct {
	Db *sql.DB
}

func NewRealmRepository(db *sql.DB) *RealmRepository {
	return &RealmRepository{Db: db}
}

func (r *RealmRepository) Create(ctx context.Context, input outbound.CreateRealmRepositoryInput) error {
	query := `
		INSERT INTO realms (id, identity_provider_id, name, created_at, updated_at, version)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.Db.ExecContext(ctx, query,
		input.ID,
		input.IdentityProviderID,
		input.Name,
		input.CreatedAt,
		input.UpdatedAt,
		input.Version,
	)
	if err != nil {
		return err
	}

	return nil
}
