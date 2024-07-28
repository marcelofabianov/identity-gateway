package postgres

import (
	"context"
	"database/sql"

	"github.com/marcelofabianov/identity-gateway/internal/port/outbound"
)

type RealmRepository struct {
	db *sql.DB
}

func NewRealmRepository(db *sql.DB) *RealmRepository {
	return &RealmRepository{db: db}
}

func (r *RealmRepository) Create(ctx context.Context, input outbound.CreateRealmRepositoryInput) error {
	query := `
		INSERT INTO realms (id, identity_provider_id, name, created_at, updated_at, version)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(ctx, query,
		input.Realm.ID.String(),
		input.Realm.IdentityProviderID,
		input.Realm.Name,
		input.Realm.CreatedAt.String(),
		input.Realm.UpdatedAt.String(),
		input.Realm.Version.Int(),
	)
	if err != nil {
		return err
	}

	return nil
}
