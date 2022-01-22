package usecases

import (
	"domain"
)

// RolePermissions grants and revokes permissions from roles
type RolePermissions struct {
	permissionsRepository domain.PermissionsRepository
}

// GrantPermissions to role
func (p *RolePermissions) GrantPermissions(roleName string, permission domain.Permission, otherPermissions ...domain.Permission) (domain.Permission, error) {
	return p.permissionsRepository.GrantPermissions(roleName, permission, otherPermissions...)
}

// RevokePermissions from role
func (p *RolePermissions) RevokePermissions(roleName string, permission domain.Permission, otherPermissions ...domain.Permission) error {
	return p.permissionsRepository.RevokePermissions(roleName, permission, otherPermissions...)
}
