package repository

import (
	"context"

	"bara-playdate-api/model/criteria"
	"bara-playdate-api/model/entity"
	"bara-playdate-api/utils/paginate"
)

type UserRepository interface {
	Insert(ctx context.Context, user entity.TableMstUser) entity.TableMstUser
	Update(ctx context.Context, user entity.TableMstUser) entity.TableMstUser
	FindByParam(ctx context.Context, key string, value string) (entity.TableMstUser, error)
	FindById(ctx context.Context, id string) (entity.TableMstUser, error)
	FindAll(ctx context.Context, paging paginate.Datapaging, options criteria.GetListOfOptions) (int64, *[]entity.TableMstUser, error)
}
