package repository

import (
	"context"

	"bara-playdate-api/model/entity"
)

type AuthRepository interface {
	Authentication(ctx context.Context, username string) (entity.TableMstUser, error)
}
