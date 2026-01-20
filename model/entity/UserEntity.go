package entity

import (
	"bara-playdate-api/utils"
	"time"
)

type TableMstUser struct {
	Id              string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	RoleId          string    `json:"roleId"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	Fullname        string    `json:"fullname"`
	IsGender        string    `json:"isGender"`
	Address         string    `json:"address"`
	HpNumber        string    `json:"hpNumber"`
	Password        string    `json:"password"`
	UrlSignature    string    `json:"urlSignature"`
	DateActivation  time.Time `json:"dateActivation"`
	EmailVerifiedAt time.Time `json:"emailVerifiedAt"`
	BirthDate       time.Time `json:"birthDate"`
	HireDate        time.Time `json:"hireDate"`
	CreatedBy       string    `json:"createdBy"`
	UpdatedBy       string    `json:"updatedBy"`
	IsActive        string    `json:"isActive"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

func (TableMstUser) TableName() string {
	return utils.NewEnv().DbSchema + ".mst_user"
}
