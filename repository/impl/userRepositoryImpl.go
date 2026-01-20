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

func NewUserRepositoryImpl(DB *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{DB: DB}
}

type userRepositoryImpl struct {
	*gorm.DB
}

func (repository *userRepositoryImpl) Insert(ctx context.Context, user entity.TableMstUser) entity.TableMstUser {
	err := repository.DB.WithContext(ctx).Create(&user).Error
	exception.PanicLogging(err)
	return user
}

func (repository *userRepositoryImpl) Update(ctx context.Context, user entity.TableMstUser) entity.TableMstUser {
	err := repository.DB.WithContext(ctx).Where("id = ?", user.Id).Updates(&user).Error
	exception.PanicLogging(err)
	return user
}

func (repository *userRepositoryImpl) FindByParam(ctx context.Context, key string, value string) (entity.TableMstUser, error) {
	var user entity.TableMstUser
	result := repository.DB.WithContext(ctx).Where(key, value).Find(&user)
	if result.RowsAffected == 0 {
		return entity.TableMstUser{}, errors.New("user Not Found")
	}
	return user, nil
}

func (repository *userRepositoryImpl) FindById(ctx context.Context, id string) (entity.TableMstUser, error) {
	var user entity.TableMstUser
	result := repository.DB.WithContext(ctx).Preload("Role").Unscoped().Where("id = ?", id).First(&user)
	if result.RowsAffected == 0 {
		return entity.TableMstUser{}, errors.New("user Not Found")
	}
	return user, nil
}

func (repository *userRepositoryImpl) FindAll(ctx context.Context, paging paginate.Datapaging, options criteria.GetListOfOptions) (int64, *[]entity.TableMstUser, error) {
	var userDatas []entity.TableMstUser

	query := repository.DB.WithContext(ctx).Table(utils.NewEnv().DbSchema + ".mst_user")

	if len(options.Status) > 0 {
		statusValuesStr, err := paginate.PrepareStatusValues(options.Status)
		if err != nil {
			return 0, nil, err
		}

		query.Where("is_active IN ? ", statusValuesStr)
	}

	if paging.FilterValue != "" && options.SearchBy != "" {

		filter := "%" + strings.ToUpper(paging.FilterValue) + "%"

		switch options.SearchBy {
		case "param_name":
			query = query.Where(`
				UPPER(username) LIKE ?
				OR UPPER(email) LIKE ?
			`, filter, filter)
		}

	}

	var totalCount int64
	query.Model(&userDatas).Count(&totalCount)
	if query.Error != nil {
		return 0, nil, query.Error
	}

	paging.BuildQueryGORM(query).Find(&userDatas)

	return totalCount, &userDatas, query.Error
}
