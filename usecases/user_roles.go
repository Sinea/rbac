package usecases

import "domain"

// UserRoles is a use case for managing user roles
type UserRoles struct {
	usersRepository domain.UsersRepository
}

// AssignRoles to this user
func (u *UserRoles) AssignRoles(userID string, role string, otherRoles ...string) error {
	return u.usersRepository.AssignRoles(userID, role, otherRoles...)
}

// UnAssignRoles from this user
func (u *UserRoles) UnAssignRoles(userID string, role string, otherRoles ...string) error {
	return u.usersRepository.UnAssignRoles(userID, role, otherRoles...)
}
