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
		INSERT INTO realm (id, identity_provider_id, name, enabled_at, created_at, updated_at, deleted_at, version)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := r.Db.ExecContext(ctx, query,
		input.ID,
		input.IdentityProviderID,
		input.Name,
		input.EnabledAt,
		input.CreatedAt,
		input.UpdatedAt,
		input.DeletedAt,
		input.Version,
	)
	if err != nil {
		return err
	}

	return nil
}
