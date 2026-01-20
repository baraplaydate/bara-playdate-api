package criteria

type UserSearchCriteria struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type StoreUserCriteria struct {
	RoleId    string `json:"roleId" form:"roleId" validate:"required"`
	Username  string `json:"username" form:"username" validate:"required"`
	Email     string `json:"email" form:"email" validate:"required"`
	Fullname  string `json:"fullname" form:"fullname" validate:"required"`
	IsGender  string `json:"isGender" form:"isGender"`
	Address   string `json:"address" form:"address"`
	HpNumber  string `json:"hpNumber" form:"hpNumber"`
	Password  string `json:"password" form:"password" validate:"required"`
	BirthDate string `json:"birthDate" form:"birthDate"`
	HireDate  string `json:"hireDate" form:"hireDate"`
	CreatedBy string `json:"createdBy" form:"createdBy" validate:"required"`
}

type UpdateUserCriteria struct {
	RoleId    string `json:"roleId" form:"roleId" validate:"required"`
	Username  string `json:"username" form:"username" validate:"required"`
	Email     string `json:"email" form:"email" validate:"required"`
	Fullname  string `json:"fullname" form:"fullname" validate:"required"`
	IsGender  string `json:"isGender" form:"isGender"`
	Address   string `json:"address" form:"address"`
	HpNumber  string `json:"hpNumber" form:"hpNumber"`
	Password  string `json:"password" form:"password"`
	BirthDate string `json:"birthDate" form:"birthDate"`
	HireDate  string `json:"hireDate" form:"hireDate"`
	CreatedBy string `json:"createdBy" form:"createdBy" validate:"required"`
}
