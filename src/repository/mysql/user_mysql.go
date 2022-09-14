package mysql

import (
	"context"
	"database/sql"
	"gin-api/src/models"
)

type userMysqlRepository struct {
	conn *sql.DB
}

func NewUserMysqlRepository(conn *sql.DB) models.IUserRepository {
	return &userMysqlRepository{conn}
}

// GetUserByEmailAndPassword implements models.UserRepository
func (*userMysqlRepository) GetUserByEmailAndPassword(ctx context.Context, username string, password string) (*models.User, error) {
	panic("unimplemented")
}

// InsertUser implements models.UserRepository
func (*userMysqlRepository) InsertUser(ctx context.Context, email string, name string, password string) error {
	panic("unimplemented")
}
