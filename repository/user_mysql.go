package repository

import (
	"errors"
	"gorm.io/gorm"
	"userCreation/domain"
)

type userRepositoryMySQL struct {
	db *gorm.DB
}

func NewUserRepositoryMySQL(db *gorm.DB) domain.UserRepository {
	return &userRepositoryMySQL{
		db: db,
	}
}

func (u *userRepositoryMySQL) Create(user domain.User) error {
	result := u.db.Create(user)
	if result.Error != nil {
		return errors.New("unexpected error while creating user")
	}
	return nil
}

func (u *userRepositoryMySQL) Delete(email string) error {
	result := u.db.Delete(&domain.User{}, email)
	if result.Error != nil {
		return errors.New("unexpected error while deleting user")
	}
	return nil
}

func (u *userRepositoryMySQL) Check(email string) (domain.User, error) {
	var user domain.User
	result := u.db.First(&user, "email = ?", email)
	if result.Error != nil {
		return domain.User{}, errors.New("unexpected error while retrieving user")
	}
	return user, nil
}