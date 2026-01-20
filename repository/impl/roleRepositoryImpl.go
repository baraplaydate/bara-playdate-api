package impl

import (
	"context"
	"errors"
	"strings"

	"bara-playdate-api/exception"
	"bara-playdate-api/model/criteria"
	"bara-playdate-api/model/entity"
	"bara-playdate-api/repository"
	"bara-playdate-api/utils"
	"bara-playdate-api/utils/paginate"

	"gorm.io/gorm"
)

func NewRoleRepositoryImpl(DB *gorm.DB) repository.RoleRepository {
	return &roleRepositoryImpl{DB: DB}
}

type roleRepositoryImpl struct {
	*gorm.DB
}

func (repository *roleRepositoryImpl) Insert(ctx context.Context, role entity.TableMstAccRole) entity.TableMstAccRole {
	err := repository.DB.WithContext(ctx).Create(&role).Error
	exception.PanicLogging(err)
	return role
}

func (repository *roleRepositoryImpl) Update(ctx context.Context, role entity.TableMstAccRole) entity.TableMstAccRole {
	err := repository.DB.WithContext(ctx).Where("id = ?", role.Id).Updates(&role).Error
	exception.PanicLogging(err)
	return role
}

func (repository *roleRepositoryImpl) FindById(ctx context.Context, id string) (entity.TableMstAccRole, error) {
	var role entity.TableMstAccRole
	result := repository.DB.WithContext(ctx).Where("id = ?", id).First(&role)
	if result.RowsAffected == 0 {
		return entity.TableMstAccRole{}, errors.New("role Not Found")
	}
	return role, nil
}

func (repository *roleRepositoryImpl) FindAll(ctx context.Context, paging paginate.Datapaging, options criteria.GetListOfOptions) (int64, *[]entity.TableMstAccRole, error) {
	// var roles []entity.RoleEntity
	// repository.DB.WithContext(ctx).Find(&roles)
	// return roles

	var roleDatas []entity.TableMstAccRole
	query := repository.DB.WithContext(ctx).Table(utils.NewEnv().DbSchema + ".mst_acc_role")

	if len(options.Status) > 0 {
		statusValuesStr, err := paginate.PrepareStatusValues(options.Status)
		if err != nil {
			return 0, nil, err
		}

		query.Where("is_active IN ? ", statusValuesStr)
	}

	if len(paging.FilterValue) > 0 && len(options.SearchBy) > 0 {
		if options.SearchBy == "param_name" {
			query.Where("upper(role_name) like  ? ", "%"+strings.ToUpper(paging.FilterValue)+"%")
		}
	}

	var totalCount int64
	query.Model(&roleDatas).Count(&totalCount)
	if query.Error != nil {
		return 0, nil, query.Error
	}

	paging.BuildQueryGORM(query).Find(&roleDatas)

	return totalCount, &roleDatas, query.Error
}
