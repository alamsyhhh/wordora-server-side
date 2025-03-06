package users

import (
	"wordora/app/modules/users/model"
)

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetMe(userID string) (*model.User, error) {
	return s.repo.GetUserByID(userID)
}

func (s *UserService) UpdateUserRole(userID, role string) error {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return err
	}

	user.Role = role
	return s.repo.UpdateUser(user)
}
