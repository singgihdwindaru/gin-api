package models

import "time"

type (
	UserRepository interface {
		InsertUser(email, name, password string) error
		GetUserByEmailAndPassword(username, password string) (*User, error)
	}

	UserUsecase interface {
		CreateUser(request UserRequest) (int64, error)
		Login(username, password string) (*UserReponse, error)
	}
)
type (
	User struct {
		Id         string     `json:"id"`
		Email      string     `json:"email"`
		Name       string     `json:"name"`
		CreatedAt  time.Time  `json:"createdAt"`
		ModifiedAt *time.Time `json:"modifiedAt"`
		CreatedBy  string     `json:"createdBy"`
		ModifiedBy *string    `json:"modifiedBy"`
	}

	UserReponse struct {
		Id    string `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	UserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
