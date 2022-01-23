package domain

import "context"

// PermissionsRepository should not exist maybe ?
type PermissionsRepository interface {
	CreatePermission(resource Resource, action Action) (Permission, error)
	DeletePermission(permission Permission) error
}

// UsersRepository manages users and theyr roles
type UsersRepository interface {
	Create(ctx context.Context, userID string) (bool, error)
	Delete(ctx context.Context, userID string) error

	AssignRole(ctx context.Context, userID string, roles string) error
	UnAssignRole(ctx context.Context, userID string, roles string) error
}

// RolesRepository manages roles and theys permissions
type RolesRepository interface {
	CreateRole(ctx context.Context, roleName string) (*Role, error)
	DeleteRole(ctx context.Context, roleName string) error

	GrantPermissions(ctx context.Context, roleName string, permission Permission, otherPermissions ...Permission) error
	RevokePermissions(ctx context.Context, roleName string, permission Permission, otherPermissions ...Permission) error
}

// UserAccessService is a service that checks if a user has a certain permission
type UserAccessChecker interface {
	HasAccess(ctx context.Context, userID string, permission Permission) (bool, error)
}
