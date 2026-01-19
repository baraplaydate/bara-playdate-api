package resClient

import (
	"bara-playdate-api/exception"
	"bara-playdate-api/model"
	"bara-playdate-api/utils"
	"context"
)

func NewHttpBinRestClient() *HttpBinRestClient {
	return &HttpBinRestClient{}
}

type HttpBinRestClient struct {
}

func (h HttpBinRestClient) PostMethod(ctx context.Context, requestBody *model.HttpBin, response *map[string]interface{}) {
	var headers []utils.HttpHeader
	headers = append(headers, utils.HttpHeader{Key: "X-Key", Value: "123456"})

	httpClient := utils.ClientComponent[model.HttpBin, map[string]interface{}]{
		HttpMethod:     "POST",
		UrlApi:         "https://httpbin.org/post",
		RequestBody:    requestBody,
		ResponseBody:   response,
		Headers:        headers,
		ConnectTimeout: 30000,
		ActiveTimeout:  30000,
	}
	err := httpClient.Execute(ctx)
	exception.PanicLogging(err)
}
