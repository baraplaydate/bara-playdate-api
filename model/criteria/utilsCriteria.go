package criteria

type IsActiveCriteria struct {
	IsActive  string `json:"isActive" validate:"required"`
	UpdatedBy string `json:"updatedBy" validate:"required"`
}

type SearchCriteria struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type SearchByKeyCriteria struct {
	Key string `json:"key" binding:"required"`
}

type SearchByValueCriteria struct {
	Value string `json:"value" binding:"required"`
}

type GetListOfOptions struct {
	SearchBy string   `json:"search_by"`
	Status   []string `json:"status"`
}

type GetMultipleListOfOptions struct {
	SearchBy  string   `json:"search_by"`
	Status    []string `json:"status"`
	ParamList []string `json:"param_list"`
}

type SearchMultipleListOfOptions struct {
	SearchBy    string   `json:"search_by"`
	Status      []string `json:"status"`
	Attribute1  string   `json:"attribute_1"`
	Attribute2  string   `json:"attribute_2"`
	Attribute3  string   `json:"attribute_3"`
	Attribute4  string   `json:"attribute_4"`
	Attribute5  string   `json:"attribute_5"`
	Attribute6  string   `json:"attribute_6"`
	Attribute7  string   `json:"attribute_7"`
	Attribute8  string   `json:"attribute_8"`
	Attribute9  string   `json:"attribute_9"`
	Attribute10 string   `json:"attribute_10"`
}
