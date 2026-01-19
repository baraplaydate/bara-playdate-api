package result

type ResponseResult struct {
	ResponseCode        int         `json:"responseCode"`
	ResponseDescription string      `json:"responseDescription"`
	ResponseTime        string      `json:"responseTime"`
	ResponseDatas       interface{} `json:"responseDatas"`
}

type DataPagingResult struct {
	PageNumber       int         `json:"pageNumber"`
	PageSize         int         `json:"pageSize"`
	TotalRecordCount int64       `json:"totalRecordCount"`
	Records          interface{} `json:"records"`
}
