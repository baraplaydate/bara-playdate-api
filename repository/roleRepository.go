package repository

import (
	"context"

	"bara-playdate-api/model/criteria"
	"bara-playdate-api/model/entity"
	"bara-playdate-api/utils/paginate"
)

type RoleRepository interface {
	Insert(ctx context.Context, role entity.TableMstAccRole) entity.TableMstAccRole
	Update(ctx context.Context, role entity.TableMstAccRole) entity.TableMstAccRole
	FindById(ctx context.Context, id string) (entity.TableMstAccRole, error)
	FindAll(ctx context.Context, paging paginate.Datapaging, options criteria.GetListOfOptions) (int64, *[]entity.TableMstAccRole, error)
}
