package models

import (
	"context"
	"time"
)

type (
	IUserRepository interface {
		InsertUser(ctx context.Context, email, name, password string) error
		GetUserByEmailAndPassword(ctx context.Context, username, password string) (*User, error)
	}

	IUserUsecase interface {
		CreateUser(ctx context.Context, request UserRequest) (int64, error)
		Login(ctx context.Context, request LoginRequest) (*UserReponse, error)
	}
)
type (
	User struct {
		Id         int64      `json:"id"`
		Email      string     `json:"email"`
		Name       string     `json:"name"`
		CreatedAt  time.Time  `json:"createdAt"`
		ModifiedAt *time.Time `json:"modifiedAt"`
		CreatedBy  string     `json:"createdBy"`
		ModifiedBy *string    `json:"modifiedBy"`
	}

	UserReponse struct {
		Id    int64  `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	LoginRequest struct {
		Email    string `binding:"required" json:"email"`
		Password string `binding:"required" json:"password"`
	}
	UserRequest struct {
		Name     string `binding:"required" json:"name"`
		Email    string `binding:"required" json:"email"`
		Password string `binding:"required" json:"password"`
	}
)
