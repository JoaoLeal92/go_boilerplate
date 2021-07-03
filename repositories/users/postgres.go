package users

import (
	"errors"

	"github.com/JoaoLeal92/goals_backend/domain/entities"
	"gorm.io/gorm"
)

// UserRepo repository struct
type UserRepo struct {
	db *gorm.DB
}

// NewRepository instantiates a new user repository
func NewRepository(conn *gorm.DB) *UserRepo {
	return &UserRepo{
		db: conn,
	}
}

// CreateUser inserts new user in database
func (r *UserRepo) CreateUser(user *entities.User) error {
	result := r.db.Create(&user)

	if result.Error != nil {
		return errors.New(result.Error.Error())
	}
	return nil
}

// FindUserByEmail searches for user on database by email
func (r *UserRepo) FindUserByEmail(email string) *entities.User {
	var user entities.User

	result := r.db.Where("email = ?", email).First(&user)
	if result.RowsAffected == 0 {
		return &user
	}

	return &user
}

// FindUserByID finds user by given ID
func (r *UserRepo) FindUserByID(ID string) *entities.User {
	var user entities.User

	result := r.db.Preload("YearObjs").Preload("MonthObjs").Preload("DailyObjs").Where("id = ?", ID).First(&user)
	if result.RowsAffected == 0 {
		return &user
	}

	return &user
}
