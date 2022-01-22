package usecases

import "domain"

// UsersManagementUserUseCase handles adding and removing users to the system
type UsersManagementUserUseCase struct {
	usersRepository domain.UsersRepository
}

// CreateUser adds a new user
func (u *UsersManagementUserUseCase) CreateUser(userID string) (bool, error) {
	return u.usersRepository.Create(userID)
}

// DeleteUser deletes a user from the database
func (u *UsersManagementUserUseCase) DeleteUser(userID string) (bool, error) {
	return u.usersRepository.Delete(userID)
}
