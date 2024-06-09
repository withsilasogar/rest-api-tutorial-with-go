package repository

import (
	"test01/internals/model"
	"test01/x/interfacesx"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUserAccount(userRequest *interfacesx.UserRegistrationRequest) (*model.User, error)
	FetchUserDetails(userEmail string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// Function to save a user
func (r *userRepository) CreateUserAccount(userRequest *interfacesx.UserRegistrationRequest) (*model.User, error) {
	user := &model.User{
		Email:    userRequest.Email,
		Username: userRequest.Username,
		FullName: userRequest.FullName,
		UserRole: model.UserRole,
	}

	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Method to return user credentials
func (r *userRepository) FetchUserDetails(userEmail string) (*model.User, error) {
	user := &model.User{}

	if err := r.db.Where("email = ?", userEmail).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
