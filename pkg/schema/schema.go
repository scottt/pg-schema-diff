package schema

import (
	"context"
	"database/sql"
	"fmt"

	internalschema "github.com/stripe/pg-schema-diff/internal/schema"
)

// GetPublicSchemaHash hash gets the hash of the "public" schema. It can be used to compare against the hash in the migration
// plan to determine if it's still valid
// We do not expose the Schema struct yet because it is subject to change, and we do not want folks depending on its API
func GetPublicSchemaHash(ctx context.Context, conn *sql.Conn) (string, error) {
	schema, err := internalschema.GetPublicSchema(ctx, conn)
	if err != nil {
		return "", fmt.Errorf("getting public schema: %w", err)
	}
	hash, err := schema.Hash()
	if err != nil {
		return "", fmt.Errorf("hashing schema: %w", err)
	}

	return hash, nil
}