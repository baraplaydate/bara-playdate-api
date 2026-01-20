package impl

import (
	"context"
	"errors"

	"bara-playdate-api/model/entity"
	"bara-playdate-api/repository"

	"gorm.io/gorm"
)

func NewAuthRepositoryImpl(DB *gorm.DB) repository.AuthRepository {
	return &authRepositoryImpl{DB: DB}
}

type authRepositoryImpl struct {
	*gorm.DB
}

func (repository *authRepositoryImpl) Authentication(ctx context.Context, usernameOrEmail string) (entity.TableMstUser, error) {
	var userResult entity.TableMstUser
	result := repository.DB.WithContext(ctx).
		Where("username = ? OR email = ?", usernameOrEmail, usernameOrEmail).
		Find(&userResult)
	if result.RowsAffected == 0 {
		return entity.TableMstUser{}, errors.New("user not found")
	}
	return userResult, nil
}
