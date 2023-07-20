package user

import (
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUp(name, email, password string) error
	SignIn(email, password string) (User, error)
}

type userService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (us *userService) SignUp(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	return us.userRepo.CreateUser(user)
}

func (us *userService) SignIn(email, password string) (User, error) {
	user, err := us.userRepo.GetUserByEmail(email)
	if err != nil {
		return User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return User{}, err
	}

	return user, nil
}
