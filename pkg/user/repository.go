package user

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(user User) error
	GetUserByEmail(email string) (User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) CreateUser(user User) error {
	return ur.db.Create(&user).Error
}

func (ur *userRepository) GetUserByEmail(email string) (User, error) {
	var user User
	err := ur.db.Where("email = ?", email).First(&user).Error
	return user, err
}
