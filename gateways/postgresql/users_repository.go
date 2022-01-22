package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	insertUserQueryTemplate = `
	INSERT INTO %s (user_id)
	VALUES ($1)
	ON CONFLICT DO NOTHING
`
	deleteUserQueryTemplate = `
	DELETE FROM %s
	WHERE user_id = $1
`
	assignRoleQueryTemplate = `
	INSERT INTO %s
	VALUES ($1, $2)
	ON CONFLICT DO NOTHING
`
	unAssignRoleQueryTempalte = `
	DELETE FROM %s
	WHERE user_id=$1 and role_name=$2
`
)

type UsersRepository struct {
	pool *pgxpool.Pool

	insertUserQuery   string
	deleteUserQuery   string
	assignRoleQuery   string
	unAssignRoleQuery string
}

func (u *UsersRepository) Create(ctx context.Context, userID string) (bool, error) {
	connection, err := u.pool.Acquire(ctx)
	if err != nil {
		return false, fmt.Errorf("error acquiring connection: %w", err)
	}
	defer connection.Release()

	_, err = connection.Exec(ctx, u.insertUserQuery, userID)
	if err != nil {
		return false, fmt.Errorf("error creating user: %w", err)
	}

	return true, nil
}

func (u *UsersRepository) Delete(ctx context.Context, userID string) error {
	connection, err := u.pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("error acquiring connection: %w", err)
	}
	defer connection.Release()

	_, err = connection.Exec(ctx, u.insertUserQuery, userID)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	return nil
}

func (u *UsersRepository) AssignRole(ctx context.Context, userID string, role string) error {
	connection, err := u.pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("error acquiring connection: %w", err)
	}
	defer connection.Release()

	_, err = connection.Exec(ctx, u.assignRoleQuery, userID, role)
	if err != nil {
		return fmt.Errorf("error assigning role %s to user %s: %w", role, userID, err)
	}
	return nil
}

func (u *UsersRepository) UnAssignRole(ctx context.Context, userID string, role string) error {
	connection, err := u.pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("error acquiring connection: %w", err)
	}
	defer connection.Release()

	_, err = connection.Exec(ctx, u.unAssignRoleQuery, userID, role)
	if err != nil {
		return fmt.Errorf("error assigning role %s to user %s: %w", role, userID, err)
	}
	return nil
}

func NewUsersRepository(pool *pgxpool.Pool, usersTable string, userRolesTable string) *UsersRepository {
	return &UsersRepository{
		pool:              pool,
		insertUserQuery:   fmt.Sprintf(insertUserQueryTemplate, usersTable),
		deleteUserQuery:   fmt.Sprintf(deleteUserQueryTemplate, usersTable),
		assignRoleQuery:   fmt.Sprintf(assignRoleQueryTemplate, userRolesTable),
		unAssignRoleQuery: fmt.Sprintf(unAssignRoleQueryTempalte, userRolesTable),
	}
}
