package criteria

type StoreRoleCriteria struct {
	RoleName    string `json:"roleName" validate:"required"`
	Description string `json:"description" validate:"required"`
	CreatedBy   string `json:"createdBy" validate:"required"`
}

type UpdateRoleCriteria struct {
	RoleName    string `json:"roleName" validate:"required"`
	Description string `json:"description" validate:"required"`
	UpdatedBy   string `json:"updatedBy" validate:"required"`
}
