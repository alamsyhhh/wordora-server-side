package users

import (
	"errors"
	"log"
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
	log.Println("UpdateUserRole called with userID:", userID, "role:", role)

	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		log.Println("Error fetching user:", err)
		return err
	}
	if user == nil {
		log.Println("User not found")
		return errors.New("user not found")
	}

	user.Role = role
	log.Println("Updating user role:", user)
	return s.repo.UpdateUser(user)
}

