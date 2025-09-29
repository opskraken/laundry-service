package userRepo

import "github.com/jmoiron/sqlx"

type UserStatus string

type User struct {
	Id          int    `json:"id"`
	UserName    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	Status      string `json:"status" validate:"uppercase"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	// Get(id int) (*User, error)
	// List() ([]*User, error)
	// Update(id int, user User) (*User, error)
	// Delete(id int) error
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}
