package postgresql

import (
	"context"
	"fmt"

	"sinea.xyz/rbac/v1/domain"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	createRoleQueryTempalte = `
		INSERT INTO %s
		VALUES ($1)
		ON CONFLICT DO NOTHING
`
	deleteRoleQueryTemplate = `
		DELETE FROM %s
		WHERE role_name=$1
`
	grantPermissionsQueryTemplate = `
	INSERT INTO %s
	VALUES ($1, $2)
	ON CONFLICT DO NOTHING
`
	revokePermissionsQueryTempalte = `
	DELETE FROM %s
	WHERE role_name=$1 AND permission=$2
`
)

type RolesRepository struct {
	pool *pgxpool.Pool

	createRoleQuery        string
	deleteRoleQuery        string
	grantPermissionsQuery  string
	revokePermissionsQuery string
}

func (r *RolesRepository) CreateRole(ctx context.Context, roleName string) (*domain.Role, error) {
	connection, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating role %s: %w", roleName, err)
	}
	defer connection.Release()

	_, err = connection.Query(ctx, r.createRoleQuery, roleName)
	if err != nil {
		return nil, fmt.Errorf("error creating role %s: %w", roleName, err)
	}
	return &domain.Role{
		Name: roleName,
	}, nil
}

func (r *RolesRepository) DeleteRole(ctx context.Context, roleName string) error {
	connection, err := r.pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("error deleting role %s: %w", roleName, err)
	}
	defer connection.Release()

	_, err = connection.Query(ctx, r.deleteRoleQuery, roleName)
	if err != nil {
		return fmt.Errorf("error deleting role %s: %w", roleName, err)
	}
	return nil
}

func (r *RolesRepository) GrantPermissions(ctx context.Context, roleName string, permission domain.Permission, otherPermissions ...domain.Permission) error {
	connection, err := r.pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("error granting permissions role %s: %w", roleName, err)
	}
	defer connection.Release()
	tx, err := connection.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error granting permissions role %s: %w", roleName, err)
	}
	permissions := append([]domain.Permission{permission}, otherPermissions...)
	for _, permission := range permissions {
		tx.Exec(ctx, r.grantPermissionsQuery, roleName, permission.String())
	}
	if err := tx.Commit(ctx); err != nil {
		_ = tx.Rollback(ctx)
		return fmt.Errorf("error granting permissions to role %s: %w", roleName, err)
	}

	return nil
}

func (r *RolesRepository) RevokePermissions(ctx context.Context, roleName string, permission domain.Permission, otherPermissions ...domain.Permission) error {
	return nil
}

func NewRolesRepository(pool *pgxpool.Pool, rolesTable string, rolePermissionsTable string) *RolesRepository {
	return &RolesRepository{
		pool:                   pool,
		createRoleQuery:        fmt.Sprintf(createRoleQueryTempalte, rolesTable),
		deleteRoleQuery:        fmt.Sprintf(deleteRoleQueryTemplate, rolesTable),
		grantPermissionsQuery:  fmt.Sprintf(grantPermissionsQueryTemplate, rolePermissionsTable),
		revokePermissionsQuery: fmt.Sprintf(revokePermissionsQueryTempalte, rolePermissionsTable),
	}
}
