package entity

import (
	"bara-playdate-api/utils"
	"time"
)

type TableMstAccRole struct {
	Id          string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	RoleName    string    `json:"roleName"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"createdBy"`
	UpdatedBy   string    `json:"updatedBy"`
	IsActive    string    `json:"isActive"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (TableMstAccRole) TableName() string {
	return utils.NewEnv().DbSchema + ".mst_acc_role"
}
