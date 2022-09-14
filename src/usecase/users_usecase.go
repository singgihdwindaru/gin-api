package usecase

import (
	"context"
	"gin-api/src/models"
)

type userUsecase struct {
	userMysqlRepo models.IUserRepository
}

func NewUserUsecase(userMysqlRepo models.IUserRepository) models.IUserUsecase {
	return &userUsecase{
		userMysqlRepo: userMysqlRepo,
	}
}

// CreateUser implements models.IUserUsecase
func (*userUsecase) CreateUser(ctx context.Context, request models.UserRequest) (int64, error) {
	panic("unimplemented")
}

// Login implements models.IUserUsecase
func (*userUsecase) Login(ctx context.Context, request models.LoginRequest) (*models.UserReponse, error) {
	panic("unimplemented")
}
